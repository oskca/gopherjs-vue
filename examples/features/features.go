package main

import (
	"fmt"
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-vue"
	"math/rand"
	"strings"
	"time"
)

// no *js.Object struct can only be manipulated by ViewModel.methods
type Todo struct {
	Time    time.Time
	Content string
}

func NewTodo(content string) *Todo {
	t := &Todo{
		Time:    time.Now(),
		Content: content,
	}
	return t
}

type Model struct {
	*js.Object                 // this is needed for bidirectional data bindings
	IntValue     int           `js:"integer"`
	Str          string        `js:"str"`
	List         []int         `js:"list"`
	Todos        []*Todo       `js:"todos"`
	CheckedItems []string      `js:"CheckedItems"`
	AllItems     []string      `js:"AllItems"`
	Now          func() string `js:"Now"`
}

func (m *Model) Inc() {
	m.IntValue += 1
	println("inc called")
}

func (m *Model) Repeat() {
	m.Str = strings.Repeat(m.Str, 3)
	vm := vue.GetVM(m)
	// println("Get(m):", vm)
	// println("m keys:", js.Keys(m.Object))
	// for i, key := range js.Keys(m.Object) {
	// 	println(i, key)
	// }
	// println("vm keys:", js.Keys(vm.Object))
	// for i, key := range js.Keys(vm.Object) {
	// 	println(i, key)
	// }
	println("integer from vm:", vm.Get("integer").Int())
}

func (m *Model) PopulateTodo() {
	// using append would cause GopherJS internalization problems
	// so it's better to use VueJS ops to manipulates the array
	// m.Todos = append(m.Todos, NewTodo(m.Str))
	vm := vue.GetVM(m)
	todos := vm.Get("todos")
	todos.Unshift(NewTodo(m.Str))
}

func (m *Model) MapTodos() {
	data := []*Todo{}
	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("%05d", rand.Int63n(100000))
		data = append(data, NewTodo(str))
	}
	obj := js.Global.Get("Object").New()
	obj.Set("todos", data)
	obj.Set("wtf", time.Now())
	vm := vue.GetVM(m)
	// wtf would be created in `vm`, this way works but not suggested
	vm.FromJS(obj)
}

func (m *Model) ShiftTodo() {
	// m.Todos = m.Todos[1:]
	vm := vue.GetVM(m)
	todos := vm.Get("todos")
	todos.Shift()
}

func (m *Model) WhatTF() string {
	println("then called", m.IntValue)
	// m.List = append(m.List, m.IntValue)
	return time.Now().String()
}

func (m *Model) DoubleInt() int {
	return 2 * m.IntValue
}

func main() {
	// register a time formating filter
	vue.NewFilter(func(t time.Time) string {
		return t.Format("2006-01-02 15:04:05")
	}).Register("timeFormat")
	// begin vm
	m := &Model{
		Object: js.Global.Get("Object").New(),
	}
	// this is the correct way to initialize the gopherjs struct
	// which would be used in the JavaScript side.
	m.IntValue = 100
	m.Str = "a string"
	m.List = []int{1, 2, 3, 4}
	m.Todos = []*Todo{}
	m.AllItems = []string{"A", "B", "C", "D", "John", "Bill"}
	m.CheckedItems = []string{"A", "B"}
	m.Now = func() string {
		println("now:", m.IntValue)
		return time.Now().String() + fmt.Sprintf(" ==> i:%d", m.IntValue)
	}
	v := vue.New("#app", m)
	v.Watch("integer", func(val *js.Object) {
		println("IntValue changed to:", val, "m.IntValue:", m.IntValue)
		m.Str = fmt.Sprintf("after watch:%d", m.IntValue)
	})
	js.Global.Set("vm", v)
}
