package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-vue"
	"github.com/oskca/gopherjs-vue/directive"
)

type Model struct {
	*js.Object
	Text string `js:"Text"`
}

func main() {
	d := directive.New("myd")
	d.SetUpdater(func(ctx *directive.Context, val *js.Object) {
		println("directive name:", ctx.Name)
		println("directive exp:", ctx.Expression)
		println("directive values:", val)
	}).Register()

	m := &Model{
		Object: js.Global.Get("Object").New(),
	}
	m.Text = "a string"
	vue.New("body", m)
	m2 := &Model{
		Object: js.Global.Get("Object").New(),
	}
}
