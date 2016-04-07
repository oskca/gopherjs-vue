package vue

import (
	"github.com/gopherjs/gopherjs/js"
)

type Filter struct {
	*js.Object
	Read  interface{} `js:"read"`
	Write interface{} `js:"write"`
}

// using interface{} type here to utilize GopherJS type convertion automatically
func NewFilter(readerfn interface{}) *Filter {
	f := &Filter{
		Object: js.Global.Get("Object").New(),
	}
	f.Read = readerfn
	return f
}

func (f *Filter) Register(name string) {
	vue.Call("filter", name, f.Object)
}
