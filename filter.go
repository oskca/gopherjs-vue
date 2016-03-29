package vue

import (
	"github.com/gopherjs/gopherjs/js"
)

type Filter struct {
	*js.Object
	Read  interface{} `js:"read"`
	Write interface{} `js:"write"`
	// Read  func(vm *ViewModel, args []*js.Object) *js.Object `js:"read"`
	// Write func(vm *ViewModel, args []*js.Object) *js.Object `js:"write"`
}

func NewFilter(read func(val *js.Object) *js.Object) *Filter {
	f := &Filter{
		Object: js.Global.Get("Object").New(),
	}
	f.Read = read
	return f
}

func (f *Filter) Register(name string) {
	vue.Call("filter", name, f.Object)
}
