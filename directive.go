package vue

import "github.com/gopherjs/gopherjs/js"
import "github.com/oskca/gopherjs-dom"

type DirectiveBinding struct {
	*js.Object
	// name: the name of the directive, without the prefix.
	Name string `js:"name"`
	// value: The value passed to the directive. For example in v-my-directive="1 + 1", the value would be 2.
	Value string `js:"value"`
	// oldValue: The previous value, only available in update and componentUpdated. It is available whether or not the value has changed.
	OldValue string `js:"oldValue"`
	// expression: the expression of the binding, excluding arguments and filters.
	Expression string `js:"expression"`
	// arg: the argument, if present.
	Arg string `js:"arg"`
	// modifiers: an object containing modifiers, if any.
	Modifiers *js.Object `js:"modifiers"`
}

// DirectiveCallback can be used in every directive callback functions
type DirectiveCallback func(el *dom.Element, b *DirectiveBinding, vNode, oldVnode *js.Object)

type Directive struct {
	*js.Object
	// // Name string
	// // advanced options
	// // Custom directive can provide a params array,
	// // and the Vue compiler will automatically extract
	// // these attributes on the element that the directive is bound to.
	// Params []string `js:"params"`
	// // If your custom directive is expected to be used on an Object,
	// // and it needs to trigger update when a nested property inside
	// // the object changes, you need to pass in deep: true in your directive definition.
	// Deep bool `js:"deep"`
	// // If your directive expects to write data back to
	// // the Vue instance, you need to pass in twoWay: true.
	// // This option allows the use of this.set(value) inside
	// // the directive:If your directive expects to write data back to
	// // the Vue instance, you need to pass in twoWay: true.
	// // This option allows the use of this.set(value) inside the directive
	// TwoWay bool `js:"twoWay"`
	// // Passing in acceptStatement:true enables
	// // your custom directive to accept inline statements like v-on does
	// AcceptStatement bool `js:"acceptStatement"`
	// // Vue compiles templates by recursively walking the DOM tree.
	// // However when it encounters a terminal directive,
	// // it will stop walking that element’s children.
	// // The terminal directive takes over the job of compiling the element and
	// // its children. For example, v-if and v-for are both terminal directives.
	// Terminal bool `js:"terminal"`
	// // You can optionally provide a priority number for your directive.
	// // If no priority is specified, a default priority will be used
	// //  - 1000 for normal directives and 2000 for terminal directives.
	// // A directive with a higher priority will be processed earlier than
	// // other directives on the same element. Directives with
	// // the same priority will be processed in the order they appear in
	// // the element’s attribute list, although that order is not
	// // guaranteed to be consistent in different browsers.
	// Priority int `js:"priority"`
}

func NewDirective(updaterCallBack ...interface{}) *Directive {
	d := &Directive{
		Object: js.Global.Get("Object").New(),
	}
	if len(updaterCallBack) > 0 {
		d.SetUpdater(updaterCallBack[0])
	}
	return d
}

func (d *Directive) SetBinder(fnCallback interface{}) *Directive {
	d.Set("bind", fnCallback)
	return d
}

func (d *Directive) SetUnBinder(fnCallback interface{}) *Directive {
	d.Set("unbind", fnCallback)
	return d
}

func (d *Directive) SetUpdater(fnCallback interface{}) *Directive {
	d.Set("update", fnCallback)
	return d
}

func (d *Directive) Register(name string) {
	js.Global.Get("Vue").Call("directive", name, d.Object)
}
