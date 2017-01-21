package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-vue"
)

type Com struct {
	*js.Object
	Text string `js:"text"`
}

func (c *Com) Hello() {
	println("hello" + c.Text)
	vm := vue.GetVM(c)
	println("vm from Hello:", vm)
	println("vm.get:", vm.Get("text").String())
}

func New() interface{} {
	cc := &Com{
		Object: js.Global.Get("Object").New(),
	}
	cc.Text = "init value"
	return cc
}

type controller struct {
	*js.Object
}

func main() {
	vue.NewComponent(New, template).Register("my-el")
	vm := vue.New("#app", new(controller))
	js.Global.Set("vm", vm)
	println("vm:", vm)
}

const (
	template = `
    <div>
        <div>Text:{{text}}</div>
        <input v-model="text"  />
        <button @click="Hello">hello</button>
    </div>
`
)
