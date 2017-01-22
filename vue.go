// Package vue provides gopherjs bindings for VueJS. see func `New` and
// the examples for detailed usage instructions.
package vue

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	vue  = js.Global.Get("Vue")
	vMap = make(map[interface{}]*ViewModel, 0)
)

// Add in the bottom of the array
func Push(obj *js.Object, any interface{}) (idx int) {
	return obj.Call("push", any).Int()
}

// Remove in the bottom of the array
func Pop(obj *js.Object) (idx int) {
	return obj.Call("pop").Int()
}

//Add in the front of the array
func Unshift(obj *js.Object, any interface{}) (idx int) {
	return obj.Call("unshift", any).Int()
}

//Remove in the front of the array
func Shift(obj *js.Object) (idx int) {
	return obj.Call("shift").Int()
}

//array slice operation
// index	required,position to add to(remove from),negative means reverse
// howmany	required,number of items to remove, 0 means no remove
// items... optional,add new items to the array
func Splice(obj *js.Object, index, howmany int, items ...interface{}) *js.Object {
	args := []interface{}{index, howmany}
	args = append(args, items...)
	return obj.Call("splice", args...)
}

func Sort(obj *js.Object, sorter func(a, b *js.Object) int) *js.Object {
	return obj.Call("sort", sorter)
}

func Reverse(obj *js.Object) *js.Object {
	return obj.Call("reverse")
}

// type Vue represents the JavaScript side VueJS instance or VueJS component
type ViewModel struct {
	*js.Object
	///////////////////////////// Instance Properties
	// vm.$data
	// 	Type: Object
	// Details:
	//  The data object that the Vue instance is observing.
	//  You can swap it with a new object.
	//  The Vue instance proxies access to the properties on its data object.
	Data *js.Object `js:"$data"`

	// vm.$el
	//  Type: HTMLElement
	//  Read only
	// Details:
	//  The DOM element that the Vue instance is managing.
	//	Note that for Fragment Instances, vm.$el will
	//  return an anchor node that
	//	indicates the starting position of the fragment.
	El *js.Object `js:"$el"`

	// vm.$options
	//  Type: Object
	//  Read only
	// Details:
	//  The instantiation options used for the current Vue instance.
	//  This is useful when you want to include custom
	//  properties in the options:
	//	 new Vue({
	//	   customOption: 'foo',
	//	   created: function () {
	//	     console.log(this.$options.customOption) // -> 'foo'
	//	   }
	//	 })
	Options *js.Object `js:"$options"`

	// vm.$parent
	//  Type: Vue instance
	//  Read only
	// Details:
	// 		The parent instance, if the current instance has one.
	Parent *js.Object `js:"$parent"`

	// vm.$root
	//  Type: Vue instance
	//  Read only
	// Details:
	// 		The root Vue instance of the current component tree.
	// 		If the current instance has no parents this value will be itself.
	Root *js.Object `js:"$root"`

	// vm.$children
	//  Type: Array<Vue instance>
	//  Read only
	// Details:
	// 	The direct child components of the current instance.
	Children *js.Object `js:"$children"`

	// vm.$slots
	// 	Type: { [name: string]: ?Array<VNode> }
	// Read only
	// Details:
	// 	Used to programmatically access content distributed by slots.
	//  Each named slot has its own corresponding property
	// 	Accessing vm.$slots is most useful when writing a component with a render function.
	Slots *js.Object `js:"$slots"`

	// vm.$refs
	// 	Type: Object
	// 	Read only
	// Details:
	// 	An object that holds child components that have v-ref registered.
	// See also:
	// 	Child Component Refs
	// 	v-ref.
	Refs *js.Object `js:"$refs"`

	// vm.$isServer
	// Type: boolean
	// Read only
	// Details:
	// Whether the current Vue instance is running on the server.
	IsServer bool `js:"$isServer"`

	// vm.$watch( expression, callback, [deep, immediate] )
	//  expression String
	//  callback( newValue, oldValue ) Function
	//  deep Boolean optional
	//  immdediate Boolean optional
	// Watch an expression on the Vue instance for changes.
	// The expression can be a single keypath or actual expressions:
	WatchEx func(
		expression string,
		callback func(newVal, oldVal *js.Object),
		deepWatch bool,
	) (unwatcher func()) `js:"$watch"`

	// vm.$set( keypath, value )
	// 	keypath String
	// 	value *
	// Set a data value on the Vue instance given a valid keypath.
	// If the path doesn’t exist it will be created.
	Set func(keypath string, val interface{}) `js:"$set"`

	// vm.$delete( keypath )
	// 	keypath String
	// Delete a root level property on the Vue instance (and also its $data).
	Delete func(keypath string) `js:"$delete"`

	////////////////////////////////// Events

	// vm.$on( event, callback )
	// 	event String
	// callback Function
	//
	// Listen for an event on the current vm
	On func(event string, callback interface{}) `js:"$on"`

	// vm.$once( event, callback )
	// 	event String
	// callback Function
	//
	// Attach a one-time only listener for an event.
	Once func(event string, callback interface{}) `js:"$once"`

	// vm.$off( [event, callback] )
	// 	event String optional
	// callback Function optional
	//
	// If no arguments are given, stop listening for all events;
	// if only the event is given, remove all callbacks for that event;
	// if both event and callback are given,
	// remove that specific callback only.
	Off func(event ...string) `js:"$off"`

	// vm.$emit( event, [args…] )
	// 	event String
	// args… optional
	//
	// Trigger an event on this vm only.
	Emit func(event string, args ...interface{}) `js:"$emit"`

	/////////////////////   Lifecycle

	// vm.$mount( [element|selector] )
	// 	element HTMLElement | selector String optional
	// If the Vue instance didn’t get an el option at instantiation,
	// you can manually call $mount() to assign an element to it and
	// start the compilation. If no argument is provided,
	// an empty <div> will be automatically created. Calling $mount()
	// on an already mounted instance will have no effect.
	// The method returns the instance itself so you can chain other
	// instance methods after it.
	Mount func(elementOrselector ...interface{}) *js.Object `js:"$mount"`

	// vm.$forceUpdate()
	// Usage:
	// Force the Vue instance to re-render.
	// Note it does not affect all child components,
	// only the instance itself and child components with inserted slot content.
	ForceUpdate func() `js:"$forceUpdate"`

	// vm.$nextTick( [callback] )
	// Arguments:
	// {Function} [callback]
	// Usage:
	// Defer the callback to be executed after the next DOM update cycle. Use it immediately after you’ve changed some data to wait for the DOM update. This is the same as the global Vue.nextTick, except that the callback’s this context is automatically bound to the instance calling this method.
	NextTick func(cb func()) `js:"$nextTick"`

	// vm.$destroy( [remove] )
	//  remove Boolean optional
	// Completely destroy a vm.
	// Clean up its connections with other existing vms,
	// unbind all its directives and remove its $el from the DOM.
	// Also, all $on and $watch listeners will be automatically removed.
	Destroy func(remove bool) `js:"$destroy"`
}

// New creates a VueJS Instance to apply bidings between `structPtr` and
// `selectorOrElementOrFunction` target, it also connects the generated
// VueJS instance with the `structPtr`, you can use `vue.GetVM` from
// you code to get the generated VueJS Instance which can be used as `this`
// for JavaScirpt side.
//
//  * all `exported fields` of the `struct` would become VueJS Instance's
//  data which can be used in the html to do data binding: v-bind, etc
//
//  * all `exported funcs` of the `struct` would become VueJS Instance's
//  methods which can be called as html event handler: v-on, etc
//
//  * the `struct` talked above should have an embeded anonymous
//  `*js.Object` field and `exported fields` should have proper
//  `js struct tag` for bidirectionaly data bindings
//
//  * if the `struct` has no embeded anonymous `*js.Object`, it can
//  only be used for information displaying purpose.
//
// Rules for exported functions usage IMPORTANT!:
//
//  * If your func uses any of the `exported fields`, then DONOT modify any.
//  These can be viewed roughly as `computed attribute` in `VueJS` wourld,
//  with the form of function invocation (invoke).
//
//  * If your func modifies any of the `exported fields`, then DONOT use it
//  in any data `DISPLAYing` expression or directive. They can be used as
//  event handlers (their main use cases).
//
// These rules are required for VueJS dependency system to work correctly.
//
// You can get this *ViewModel instance through `vue.GetVM(structPtr)`
// which acts as `this` of the VueJS(javascript) side of world
func New(selectorOrHTMLElement interface{}, structPtr interface{}) *ViewModel {
	opt := NewOption()
	opt.El = selectorOrHTMLElement
	opt.SetDataWithMethods(structPtr)
	vm := opt.NewViewModel()
	vMap[structPtr] = vm
	return vm
}

func newViewModel(o *js.Object) *ViewModel {
	return &ViewModel{
		Object: o,
	}
}

// GetVM returns coresponding VueJS instance from a gopherjs struct pointer
// (the underlying ViewModel data), this function is mainly in
// gopherjs struct method functions to reference the `VueJS instance`
func GetVM(structPtr interface{}) *ViewModel {
	vm, ok := vMap[structPtr]
	if !ok {
		panic("GetVM: Vue not registerd yet")
	}
	return vm
}

// Watch using a simpler form to do Vue.$watch
func (v *ViewModel) Watch(expression string, callback func(newVal *js.Object)) (unwatcher func()) {
	obj := v.Call("$watch", expression, callback)
	return func() {
		obj.Invoke()
	}
}
