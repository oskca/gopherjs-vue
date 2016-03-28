package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-vue"
	"time"
)

type Test struct {
	*js.Object
	Time time.Time `js:"Time"`
	Text string    `js:"Text"`
}

func main() {
	t := &Test{
		Object: js.Global.Get("Object").New(),
	}
	t.Text = "Hello World"
	t.Time = time.Now()
	vm := vue.New("#app", t)
	js.Global.Set("vm", vm)
}
