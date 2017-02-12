package vue

import (
	"github.com/gopherjs/gopherjs/js"
)

// Filter return interface{} to utilize GopherJS type convertion automatically
type Filter func(oldValue *js.Object) (newValue interface{})

// using interface{} type here to utilize GopherJS type convertion automatically
func NewFilter(fn func(oldValue *js.Object) (newValue interface{})) Filter {
	return Filter(fn)
}

func (f Filter) Register(name string) {
	vue.Call("filter", name, f)
}
