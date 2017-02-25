# gopherjs-vue
VueJS bindings for gopherjs

# Usage

Combined the power of [Gopherjs][gopherjs] and [VueJS][vuejs], you can use
`golang struct` to provide the two-way data-binding context for [VueJS][vuejs],
and easily implements the popular browser `MVVM` models in Go.

Currently `ViewModel/Component/Directive/Filter` are supported and wrapped in
a gopherjs friendly way.

These are the basic rules to use this package:

* all `exported fields` of the `golang struct` would become VueJS Instance's
  data which can be used in the html to do data binding: v-bind, etc

* all `exported funcs` of the `golang struct` would become VueJS Instance's
  methods which can be called as html event handler: v-on, etc

* the `golang struct` talked above is actually of pointer type and
  should have an anonymous embeded `*js.Object` field and the `exported fields`
  should have proper `js struct tag` for bidirectionaly data bindings

# Using the debug|dev version of VueJS

This package includes the `minified|product version` of VueJS code by default, 
if you want to include the `debug|dev version` of of VueJS code, please specify
buidling tags `debug` to the `gopherjs build` cmd as this:

    gopherjs build --tags debug main.go

for more details please see the examples.

# Basic example

gopherjs code:

```go
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
    m := &Model{
        Object: js.Global.Get("Object").New(),
    }
    // field assignment is required in this way to make data passing works
    m.IntValue = 100
    m.Str = "a string"
    // create the VueJS viewModel using a struct pointer
    vue.New("#app", m)
}
```


html markup:

```html
<!DOCTYPE html>
<html>

<body>
    <div id="app" v-cloak>
        <div>integer: {{ integer }}
            <input v-model="integer"></input>
        </div>
        <div>str: {{ str }} </div>
        <button v-on:click="Inc">Increase</button>
        <button v-on:click="Repeat">Repeat</button>
        <button v-on:click="Reset">Reset</button>
    </div>
    <script type="text/javascript" src="basic.js"></script>
</body>

</html>
```

compile and run, then there you are :)

[gopherjs]: https://github.com/gopherjs/gopherjs
[vuejs]: http://vuejs.org/
