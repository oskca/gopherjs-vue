package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-vue"
)

type Model struct {
	*js.Object        // this is needed for bidirectional data bindings
	IntValue   int    `js:"integer"`
	Str        string `js:"str"`
}

// this would be recognized as Inc in html
func (m *Model) Inc() {
	m.IntValue += 1
	println("inc called")
}

// this would be recognized as Repeat in html
func (m *Model) Repeat() {
	m.Str = m.Str + m.Str
}

// this would be recognized as Reset in html
func (m *Model) Reset() {
	m.Str = "a string "
}

func main() {
	// model
	m := &Model{
		Object: js.Global.Get("Object").New(),
	}
	// field assignment is required in this way to make data passing works
	m.IntValue = 100
	m.Str = "a string"
	// options
	o := vue.NewOption()
	o.SetDataWithMethods(m)
	o.AddComputed("double", func(vm *vue.ViewModel) interface{} {
		println("reading computed double")
		i := vm.Data.Get("integer").Int()
		return i * 2
	})
	v := o.NewViewModel()
	v.Mount("#app")
}
