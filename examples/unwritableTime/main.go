package main

import (
	"github.com/gopherjs/gopherjs/js"
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
	println(time.Now())
	t.Time = time.Now()
}
