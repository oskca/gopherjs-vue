package directive

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-vue"
)

type Context struct {
	*js.Object
	// el: the element the directive is bound to.
	El *js.Object `js:"el"`
	// vm: the context ViewModel that owns this directive.
	Vm *vue.Vue `js:"vm"`
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

func makeUpdater(fn func(ctx *Context, newValue, oldValue *js.Object)) *js.Object {
	return js.MakeFunc(func(this *js.Object, args []*js.Object) interface{} {
		ctx := &Context{
			Object: this,
		}
		fn(ctx, args[0], args[1])
		return nil
	})
}

func makeBinder(fn func(*Context)) *js.Object {
	return js.MakeFunc(func(this *js.Object, args []*js.Object) interface{} {
		ctx := &Context{
			Object: this,
		}
		fn(ctx)
		return nil
	})
}

// func Directive(
// 	name string,
// 	update func(ctx *Context, newValue, oldValue *js.Object),
// ) {
// 	fn := js.MakeFunc(func(this *js.Object, args []*js.Object) interface{} {
// 		ctx := &Context{
// 			Object: this,
// 		}
// 		update(ctx, args[0], args[1])
// 		return nil
// 	})
// 	vue.Call("directive", name, fn)
// }

// func DirectiveEx(
// 	name string,
// 	bind func(ctx *Context),
// 	update func(ctx *Context, newValue, oldValue *js.Object),
// 	unbind func(ctx *Context),
// ) {
// 	fnInit := js.MakeFunc(func(this *js.Object, args []*js.Object) interface{} {
// 		ctx := &Context{
// 			Object: this,
// 		}
// 		bind(ctx)
// 		return nil
// 	})
// 	fnUpdate := js.MakeFunc(func(this *js.Object, args []*js.Object) interface{} {
// 		ctx := &Context{
// 			Object: this,
// 		}
// 		update(ctx, args[0], args[1])
// 		return nil
// 	})
// 	fnUnbind := js.MakeFunc(func(this *js.Object, args []*js.Object) interface{} {
// 		ctx := &Context{
// 			Object: this,
// 		}
// 		unbind(ctx)
// 		return nil
// 	})
// 	vue.Call("directive", name, js.M{
// 		"bind":   fnInit,
// 		"update": fnUpdate,
// 		"unbind": fnUnbind,
// 	})
// }

// func Directive(
// 	name string,
// 	update func(ctx *Context, newValue, oldValue *js.Object),
// ) {
// 	vue.Call("directive", name, makeUpdater(update))
// }

// func DirectiveEx(
// 	name string,
// 	bind func(ctx *Context),
// 	update func(ctx *Context, newValue, oldValue *js.Object),
// 	unbind func(ctx *Context),
// ) {
// 	vue.Call("directive", name, js.M{
// 		"bind":   makeBinder(bind),
// 		"update": makeUpdater(update),
// 		"unbind": makeBinder(unbind),
// 	})
// }

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
// func ElementDirective(
// 	name string,
// 	update func(ctx *Context, newValue, oldValue *js.Object),
// ) {
// 	vue.Call("elementDirective", name, makeUpdater(update))
// }

// func ElementDirectiveEx(
// 	name string,
// 	bind func(ctx *Context),
// 	update func(ctx *Context, newValue, oldValue *js.Object),
// 	unbind func(ctx *Context),
// ) {
// 	vue.Call("elementDirective", name, js.M{
// 		"bind":   makeBinder(bind),
// 		"update": makeUpdater(update),
// 		"unbind": makeBinder(unbind),
// 	})
// }

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

func New(name string, updater ...func(ctx *Context, newValue, oldValue *js.Object)) *Directive {
	d := &Directive{
		Name:   name,
		Object: js.Global.Get("Object").New(),
	}
	if len(updater) > 0 {
		d.SetUpdater(updater[0])
	}
	return d
}

func (d *Directive) SetBinder(fn func(ctx *Context)) *Directive {
	d.Set("bind", makeBinder(fn))
	return d
}

func (d *Directive) SetUnBinder(fn func(ctx *Context)) *Directive {
	d.Set("unbind", makeBinder(fn))
	return d
}

func (d *Directive) SetUpdater(fn func(ctx *Context, newValue, oldValue *js.Object)) *Directive {
	d.Set("update", makeUpdater(fn))
	return d
}

func (d *Directive) Register() {
	js.Global.Get("Vue").Call("directive", d.Name, d.Object)
}
