package vue

import (
	"github.com/gopherjs/gopherjs/js"
)

type DirectiveContext struct {
	*js.Object
	// el: the element the directive is bound to.
	El *js.Object `js:"el"`
	// vm: the context ViewModel that owns this directive.
	Vm *ViewModel `js:"vm"`
	// expression: the expression of the binding, excluding arguments and filters.
	Expression string `js:"expression"`
	// arg: the argument, if present.
	Arg string `js:"arg"`
	// name: the name of the directive, without the prefix.
	Name string `js:"name"`
	// modifiers: an object containing modifiers, if any.
	Modifiers *js.Object `js:"modifiers"`
	// descriptor: an object that contains the parsing result of the entire directive.
	Descriptor *js.Object `js:"descriptor"`
	// params: an object containing param attributes. Explained below.
	Params *js.Object `js:"params"`
}

func makeUpdater(fn func(ctx *DirectiveContext, val *js.Object)) *js.Object {
	return js.MakeFunc(func(this *js.Object, args []*js.Object) interface{} {
		ctx := &DirectiveContext{
			Object: this,
		}
		fn(ctx, args[0])
		return nil
	})
}

func makeBinder(fn func(*DirectiveContext)) *js.Object {
	return js.MakeFunc(func(this *js.Object, args []*js.Object) interface{} {
		ctx := &DirectiveContext{
			Object: this,
		}
		fn(ctx)
		return nil
	})
}

type Directive struct {
	*js.Object
	Name string
	// advanced options
	// Custom directive can provide a params array,
	// and the Vue compiler will automatically extract
	// these attributes on the element that the directive is bound to.
	Params []string `js:"params"`
	// If your custom directive is expected to be used on an Object,
	// and it needs to trigger update when a nested property inside
	// the object changes, you need to pass in deep: true in your directive definition.
	Deep bool `js:"deep"`
	// If your directive expects to write data back to
	// the Vue instance, you need to pass in twoWay: true.
	// This option allows the use of this.set(value) inside
	// the directive:If your directive expects to write data back to
	// the Vue instance, you need to pass in twoWay: true.
	// This option allows the use of this.set(value) inside the directive
	TwoWay bool `js:"twoWay"`
	// Passing in acceptStatement:true enables
	// your custom directive to accept inline statements like v-on does
	AcceptStatement bool `js:"acceptStatement"`
	// Vue compiles templates by recursively walking the DOM tree.
	// However when it encounters a terminal directive,
	// it will stop walking that element’s children.
	// The terminal directive takes over the job of compiling the element and
	// its children. For example, v-if and v-for are both terminal directives.
	Terminal bool `js:"terminal"`
	// You can optionally provide a priority number for your directive.
	// If no priority is specified, a default priority will be used
	//  - 1000 for normal directives and 2000 for terminal directives.
	// A directive with a higher priority will be processed earlier than
	// other directives on the same element. Directives with
	// the same priority will be processed in the order they appear in
	// the element’s attribute list, although that order is not
	// guaranteed to be consistent in different browsers.
	Priority int `js:"priority"`
}

func NewDirective(updater ...func(ctx *DirectiveContext, val *js.Object)) *Directive {
	d := &Directive{
		Object: js.Global.Get("Object").New(),
	}
	if len(updater) > 0 {
		d.SetUpdater(updater[0])
	}
	return d
}

func (d *Directive) SetBinder(fn func(ctx *DirectiveContext)) *Directive {
	d.Set("bind", makeBinder(fn))
	return d
}

func (d *Directive) SetUnBinder(fn func(ctx *DirectiveContext)) *Directive {
	d.Set("unbind", makeBinder(fn))
	return d
}

func (d *Directive) SetUpdater(fn func(ctx *DirectiveContext, val *js.Object)) *Directive {
	d.Set("update", makeUpdater(fn))
	return d
}

func (d *Directive) Register(name string) {
	js.Global.Get("Vue").Call("directive", name, d.Object)
}

// In some cases, we may want our directive to be used in the form of
// a custom element rather than as an attribute.
// This is very similar to Angular’s notion of “E” mode directives.
// Element directives provide a lighter-weight alternative to
// full-blown components (which are explained later in the guide).
//
// Element directives cannot accept arguments or
// expressions, but it can read the element’s attributes to
// determine its behavior.
//
// A big difference from normal directives is that
// element directives are terminal, which means once Vue encounters
// an element directive, it will completely skip that element
// - only the element directive itself will be
// able to manipulate that element and its children.
type ElementDirective struct {
	*Directive
}

func NewElementDirective(updater ...func(ctx *DirectiveContext, val *js.Object)) *ElementDirective {
	d := &ElementDirective{
		Directive: NewDirective(updater...),
	}
	if len(updater) > 0 {
		d.SetUpdater(updater[0])
	}
	return d
}

func (d *ElementDirective) Register(name string) {
	js.Global.Get("Vue").Call("elementDirective", name, d.Object)
}
