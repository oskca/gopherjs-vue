package vue

import (
	"github.com/gopherjs/gopherjs/js"
)

func Directive(name string, opt js.M) {
	vue.Call("directive", name, opt)
}

func Filter(name string, fn func(*js.Object) *js.Object) {
	vue.Call("filter", name, fn)
}

// Vue.partial( id, [definition] )
// id String
// definition String | Node optional
// Register or retrieve a global partial.
// The definition can be a template string, a querySelector that starts with #,
// a DOM element (whose innerHTML will be used as the template string), or a DocumentFragment.
func Partial(name, definition string) {
	vue.Call("partial", name, definition)
}
