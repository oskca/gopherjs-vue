package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-vue"
	"time"
)

type Embed struct {
	*js.Object
	Time time.Time `js:"time"`
	Text string    `js:"text"`
}

func (t *Embed) PtrHello() string {
	return t.Text + "World"
}

func (t Embed) ValueHello() string {
	return t.Text + "World"
}

type Test struct {
	Time time.Time //`js:"time"`
	Text string    //`js:"text"`
}

func (t *Test) PtrHello() string {
	return t.Text + "World"
}

func (t Test) ValueHello() string {
	return t.Text + "World"
}

func main() {
	vt := Test{
		// Object: js.Global.Get("Object").New(),
		Text: "a string",
	}
	vt.Text = "Hello World"
	t := &Test{
	// Object: js.Global.Get("Object").New(),
	}
	t.Text = "Hello World"
	// t.Time = time.Now()
	println("t", t)
	println("vt", vt)
	vm := vue.New("#app", vt)
	println("vm", vm.Object)
	embed := Embed{
		Object: js.Global.Get("Object").New(),
		Text:   "a string",
	}
	js.Global.Set("vm", vm)
	js.Global.Set("vt", vt)
	js.Global.Set("v", t)
	js.Global.Set("MakeWrapper", js.MakeWrapper(vt))
	embed.Text = "embed test"
	println("embed", embed)
	js.Global.Set("embed", embed)
}
