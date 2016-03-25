package vue

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	vue  = js.Global.Get("Vue")
	vMap = make(map[interface{}]*Vue, 0)
)

type Vue struct {
	*js.Object
}

// New register a new Vue from a struct instance `structPtr`
//
//  all `exported field` of the `struct` would become VueJS Instance's data
//  which can be used in the html to do data binding: v-bind, etc
//
//  all `exported funcs` of the `struct` would become VueJS Instance's methods
//  which can be called as html event handler: v-on, etc
//
//  You can get this *Vue instance through `vue.Get(structPtr)`
//  which is actually `this` of the VueJS(javascript) side of world
func New(cssSelector string, structPtr interface{}) *Vue {
	o := vue.New(js.M{
		"el":      cssSelector,
		"data":    structPtr,
		"methods": js.MakeWrapper(structPtr),
	})
	vm := &Vue{
		Object: o,
	}
	vMap[structPtr] = vm
	return vm
}

// Get returns the related *Vue instance (as this for VueJS)
func Get(structPtr interface{}) *Vue {
	vm, ok := vMap[structPtr]
	if !ok {
		println("Unregistered Vue:", structPtr)
		panic("Vue not registerd yet")
	}
	return vm
}

// vm.$watch( expression, callback, [deep, immediate] )
//
// expression String
// callback( newValue, oldValue ) Function
// deep Boolean optional
// immdediate Boolean optional
//
// Watch an expression on the Vue instance for changes.
// The expression can be a single keypath or actual expressions:
func (v *Vue) WatchRaw(expression string, callback func(newVal, oldVal *js.Object), deepWatch bool) (unwatcher func()) {
	obj := v.Call("$watch", expression, callback, deepWatch)
	return func() {
		obj.Invoke()
	}
}

func (v *Vue) Watch(expression string, callback func(newVal *js.Object)) (unwatcher func()) {
	obj := v.Call("$watch", expression, callback)
	return func() {
		obj.Invoke()
	}
}

// vm.$eval( expression )
//
// expression String
// Evaluate an expression that can also contain filters.
//
// // assuming vm.msg = 'hello'
// vm.$eval('msg | uppercase') // -> 'HELLO'
func (v *Vue) Eval(expression string) *js.Object {
	return v.Call("$eval", expression)
}

// vm.$get( expression )
//
// expression String
//
// Retrieve a value from the Vue instance given an expression.
// Expressions that throw errors will be suppressed and return undefined.
func (v *Vue) Get(expression string) *js.Object {
	return v.Call("$get", expression)
}

// vm.$set( keypath, value )

// keypath String
// value *
//
// Set a data value on the Vue instance given a valid keypath.
// If the path doesn’t exist it will be created.
func (v *Vue) Set(keypath string, val interface{}) {
	v.Call("$set", keypath, val)
}

// vm.$add( keypath, value )
//
// keypath String
// value *
// Add a root level property to the Vue instance (and also its $data). Due to the limitations of ES5, Vue cannot detect properties directly added to or deleted from an Object, so use this method and vm.$delete when you need to do so. Additionally, all observed objects are augmented with these two methods too.
func (v *Vue) Add(keypath string, val interface{}) {
	v.Call("$add", keypath, val)
}

// vm.$delete( keypath )
//
// keypath String
// Delete a root level property on the Vue instance (and also its $data).
func (v *Vue) Delete(keypath string) {
	v.Call("$delete", keypath)
}

// vm.$interpolate( templateString )
// templateString String
// Evaluate a piece of template string containing mustache interpolations.
// Note that this method simply performs string interpolation;
// attribute directives are not compiled.
//
// // assuming vm.msg = 'hello'
// vm.$interpolate('{{msg}} world!') // -> 'hello world!'
func (v *Vue) Interpolate(templateString string) {
	v.Call("$interpolate", templateString)
}

// Events
// Each vm is also an event emitter.
// When you have multiple nested ViewModels,
// you can use the event system to communicate between them.
//
// vm.$dispatch( event, [args…] )
// event String
// args… optional
//
// Dispatch an event from the current vm that propagates all the way up to its $root.
// If a callback returns false, it will stop the propagation at its owner instance.
func (v *Vue) Dispatch(event string, args ...interface{}) {
	args = append([]interface{}{event}, args...)
	v.Call("$dispatch", args...)
}

// vm.$broadcast( event, [args…] )
// event String
// args… optional
//
// Emit an event to all children vms of the current vm,
// which gets further broadcasted to their children all the way down.
// If a callback returns false, its owner instance will not broadcast the event any further.
func (v *Vue) Broadcast(event string, args ...interface{}) {
	args = append([]interface{}{event}, args...)
	v.Call("$broadcast", args...)
}

// vm.$emit( event, [args…] )
// event String
// args… optional
//
// Trigger an event on this vm only.
func (v *Vue) Emit(event string, args ...interface{}) {
	args = append([]interface{}{event}, args...)
	v.Call("$emit", args...)
}

type EventCallback func(args ...interface{})

// vm.$on( event, callback )
// event String
// callback Function
//
// Listen for an event on the current vm
func (v *Vue) On(event string, cb EventCallback) {
	v.Call("$on", event, cb)
}

// vm.$once( event, callback )
// event String
// callback Function
//
// Attach a one-time only listener for an event.
func (v *Vue) Once(event string, cb EventCallback) {
	v.Call("$once", event, cb)
}

// vm.$off( [event, callback] )
// event String optional
// callback Function optional
//
// If no arguments are given, stop listening for all events;
// if only the event is given, remove all callbacks for that event;
// if both event and callback are given, remove that specific callback only.
func (v *Vue) Off(event ...string) {
	v.Call("$off", event)
}

// DOM
// All vm DOM manipulation methods work like their jQuery counterparts - except they also trigger Vue.js transitions if there are any declared on vm’s $el. For more details on transitions see Adding Transition Effects.

// vm.$appendTo( element|selector, [callback] )
// element HTMLElement | selector String
// callback Function optional
// Append the vm’s $el to target element. The argument can be either an element or a querySelector string.
func (v *Vue) AppendTo(elementOrselector string) {
	v.Call("$appendTo", elementOrselector)
}

// vm.$before( element|selector, [callback] )
// element HTMLElement | selector String
// callback Function optional
// Insert the vm’s $el before target element.
func (v *Vue) Before(elementOrselector string) {
	v.Call("$before", elementOrselector)
}

// vm.$after( element|selector, [callback] )
// element HTMLElement | selector String
// callback Function optional
// Insert the vm’s $el after target element.
func (v *Vue) After(elementOrselector string) {
	v.Call("$after", elementOrselector)
}

// vm.$remove( [callback] )
// callback Function optional
// Remove the vm’s $el from the DOM.
func (v *Vue) Remove() {
	v.Call("$remove")
}

// Lifecycle

// vm.$mount( [element|selector] )
// element HTMLElement | selector String optional
// If the Vue instance didn’t get an el option at instantiation, you can manually call $mount() to assign an element to it and start the compilation. If no argument is provided, an empty <div> will be automatically created. Calling $mount() on an already mounted instance will have no effect. The method returns the instance itself so you can chain other instance methods after it.
func (v *Vue) Mount(elementOrselector string) {
	v.Call("$mount", elementOrselector)
}

// vm.$destroy( [remove] )
// remove Boolean optional
// Completely destroy a vm. Clean up its connections with other existing vms, unbind all its directives and remove its $el from the DOM. Also, all $on and $watch listeners will be automatically removed.
func (v *Vue) Destroy(remove bool) {
	v.Call("$destroy", remove)
}

// vm.$compile( element )
// element HTMLElement
// Partially compile a piece of DOM (Element or DocumentFragment). The method returns a decompile function that tearsdown the directives created during the process. Note the decompile function does not remove the DOM. This method is exposed primarily for writing advanced custom directives.
func (v *Vue) Compile(element string) {
	v.Call("$compile", element)
}

// vm.$addChild( [options, constructor] )
// options Object optional
// constructor Function optional
//
// Adds a child instance to the current instance.
// The options object is the same in manually instantiating an instance.
// Optionally you can pass in a constructor created from Vue.extend().
//
// There are three implications of a parent-child relationship between instances:
// The parent and child can communicate via the event system.
// The child has access to all parent assets (e.g. custom directives).
// The child, if inheriting parent scope, has access to parent scope data properties.
func (v *Vue) AddChild(options js.M) {
	v.Call("$addChild", options)
}

// vm.$log( [keypath] )
//
// keypath String optional
// Log the current instance data as a plain object, which is more console-inspectable than a bunch of getter/setters. Also accepts an optional key.
//
// vm.$log() // logs entire ViewModel data
// vm.$log('item') // logs vm.item
func (v *Vue) Log(keypath ...interface{}) {
	v.Call("$log", keypath...)
}
