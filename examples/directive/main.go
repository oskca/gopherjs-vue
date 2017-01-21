package main

import (
	"time"

	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-vue"
)

type Model struct {
	*js.Object
	Text string `js:"Text"`
	Time string `js:"Time"`
}

func main() {
	d := vue.NewDirective()
	d.SetUpdater(func(el *js.Object, ctx *vue.DirectiveBinding, val *js.Object) {
		println("directive name:", ctx.Name)
		println("directive exp:", ctx.Expression)
		println("directive values:", val)
	}).Register("myd")

	m := &Model{
		Object: js.Global.Get("Object").New(),
	}
	m.Text = "a string"
	m.Time = time.Now().String()
	vue.New("#app", m)
}
