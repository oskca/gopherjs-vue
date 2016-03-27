// Package vue provides gopherjs bindings for VueJS. see func `New` for
//  futher information
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

	// vm.$refs
	// 	Type: Object
	// 	Read only
	// Details:
	// 	An object that holds child components that have v-ref registered.
	// See also:
	// 	Child Component Refs
	// 	v-ref.
	Refs *js.Object `js:"$refs"`

	// vm.$els
	// 	Type: Object
	// 	Read only
	// Details:
	// 	An object that holds DOM elements that have v-el registered.
	// See also: v-el.
	Els *js.Object `js:"$els"`

	// vm.$watch( expression, callback, [deep, immediate] )
	//  expression String
	//  callback( newValue, oldValue ) Function
	//  deep Boolean optional
	//  immdediate Boolean optional
	// Watch an expression on the Vue instance for changes.
	// The expression can be a single keypath or actual expressions:
	Watch func(
		expression string,
		callback func(newVal, oldVal *js.Object),
		deepWatch bool,
	) (unwatcher func()) `js:"$watch"`

	// vm.$eval( expression )
	// 	expression String
	// Evaluate an expression that can also contain filters.
	// assuming vm.msg = 'hello'
	// vm.$eval('msg | uppercase') // -> 'HELLO'
	Eval func(expression string) *js.Object `js:"$eval"`

	// vm.$get( expression )
	// 	expression String
	// Retrieve a value from the Vue instance given an expression.
	// Expressions that throw errors will be suppressed
	// and return undefined.
	Get func(expression string) *js.Object `js:"$get"`

	// vm.$set( keypath, value )
	// 	keypath String
	// 	value *
	// Set a data value on the Vue instance given a valid keypath.
	// If the path doesn’t exist it will be created.
	Set func(keypath string, val interface{}) `js:"$set"`

	// vm.$add( keypath, value )
	//
	// 	keypath String
	// 	value *
	// Add a root level property to the Vue instance (and also its $data).
	// Due to the limitations of ES5, Vue cannot detect properties directly
	// added to or deleted from an Object,
	// so use this method and vm.$delete when you need to do so. Additionally,
	// all observed objects are augmented with these two methods too.
	Add func(keypath string, val interface{}) `js:"$add"`

	// vm.$delete( keypath )
	// 	keypath String
	// Delete a root level property on the Vue instance (and also its $data).
	Delete func(keypath string) `js:"$delete"`

	// vm.$interpolate( templateString )
	// 	templateString String
	// Evaluate a piece of template string containing
	// mustache interpolations.
	// Note that this method simply performs string interpolation;
	// attribute directives are not compiled.
	//
	// // assuming vm.msg = 'hello'
	// vm.$interpolate('{{msg}} world!') // -> 'hello world!'
	Interpolate func(templateString string) `js:"$interpolate"`

	////////////////////////////////// Events

	// Each vm is also an event emitter.
	// When you have multiple nested ViewModels,
	// you can use the event system to communicate between them.
	//
	// vm.$dispatch( event, [args…] )
	//	event String
	// 	args… optional
	//
	// Dispatch an event from the current vm that propagates
	// all the way up to its $root. If a callback returns false,
	// it will stop the propagation at its owner instance.
	Dispatch func(event string, args ...interface{}) `js:"$dispatch"`

	// vm.$broadcast( event, [args…] )
	// 	event String
	// 	args… optional
	//
	// Emit an event to all children vms of the current vm,
	// which gets further broadcasted to their children all the way down.
	// If a callback returns false, its owner instance will not broadcast
	// the event any further.
	Broadcast func(event string, args ...interface{}) `js:"$broadcast"`

	// vm.$emit( event, [args…] )
	// 	event String
	// args… optional
	//
	// Trigger an event on this vm only.
	Emit func(event string, args ...interface{}) `js:"$emit"`

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

	// DOM
	// All vm DOM manipulation methods work like their jQuery counterparts -
	// except they also trigger Vue.js transitions if there are any declared
	// on vm’s $el. For more details on transitions
	// see Adding Transition Effects.

	// vm.$appendTo( element|selector, [callback] )
	// 	element HTMLElement | selector String
	// callback Function optional
	// Append the vm’s $el to target element. The argument can be either
	// an element or a querySelector string.
	AppendTo func(elementOrselector string) `js:"$appendTo"`

	// vm.$before( element|selector, [callback] )
	// 	element HTMLElement | selector String
	// callback Function optional
	// Insert the vm’s $el before target element.
	Before func(elementOrselector string) `js:"$before"`

	// vm.$after( element|selector, [callback] )
	// 	element HTMLElement | selector String
	// callback Function optional
	// Insert the vm’s $el after target element.
	After func(elementOrselector string) `js:"$after"`

	// vm.$remove( [callback] )
	// 	callback Function optional
	// Remove the vm’s $el from the DOM.
	Remove func() `js:"$remove"`

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
	Mount func(elementOrselector string) *Vue `js:"$mount"`

	// vm.$destroy( [remove] )
	//  remove Boolean optional
	// Completely destroy a vm.
	// Clean up its connections with other existing vms,
	// unbind all its directives and remove its $el from the DOM.
	// Also, all $on and $watch listeners will be automatically removed.
	Destroy func(remove bool) `js:"$destroy"`

	// vm.$compile( element )
	// 	element HTMLElement
	// Partially compile a piece of DOM (Element or DocumentFragment).
	// The method returns a decompile function that tearsdown the directives
	// created during the process.
	// Note the decompile function does not remove the DOM.
	// This method is exposed primarily for
	// writing advanced custom directives.
	Compile func(element string) `js:"$compile"`

	// vm.$addChild( [options, constructor] )
	//  options Object optional
	//  constructor Function optional
	//
	// Adds a child instance to the current instance.
	// The options object is the same in manually instantiating an instance.
	// Optionally you can pass in a constructor created from Vue.extend().
	//
	// There are three implications of
	// a parent-child relationship between instances:
	//  The parent and child can communicate via the event system.
	//  The child has access to all parent assets (e.g. custom directives).
	//  The child, if inheriting parent scope,
	//   has access to parent scope data properties.
	AddChild func(options js.M) `js:"$addChild"`

	// vm.$log( [keypath] )
	//
	// keypath String optional
	// Log the current instance data as a plain object, which is more
	// console-inspectable than a bunch of getter/setters.
	// Also accepts an optional key.
	//
	// vm.$log() // logs entire ViewModel data
	// vm.$log('item') // logs vm.item
	Log func(keypath ...interface{}) `js:"$log"`
}

// New creates a VueJS Instance to apply bidings between `structPtr` and
// `selectorOrElementOrFunction` target, it also connects the generated VueJS
// instance with the `structPtr`, you can use `vue.GetVM` from you code to
// access the generated VueJS Instance (can be used as `this` for JavaScirpt side).
//
//  * all `exported fields` of the `struct` would become VueJS Instance's data
//  which can be used in the html to do data binding: v-bind, etc
//
//  * all `exported funcs` of the `struct` would become VueJS Instance's methods
//  which can be called as html event handler: v-on, etc
//
//  * the `struct` talked above should have an embeded anonymous `*js.Object` field
//  and `exported fields` should have proper `js struct tag` for
//  bidirectionaly data bindings
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
// You can get this *Vue instance through `vue.GetVM(structPtr)`
// which acts as `this` of the VueJS(javascript) side of world
func New(selectorOrElementOrFunction interface{}, structPtr interface{}) *Vue {
	o := vue.New(js.M{
		"el":      selectorOrElementOrFunction,
		"data":    structPtr,
		"methods": js.MakeWrapper(structPtr),
	})
	vm := &Vue{
		Object: o,
	}
	vMap[structPtr] = vm
	return vm
}

// GetVM returns coresponding VueJS instance from a gopherjs struct pointer
// (the underlying ViewModel data), this function is mainly in
// gopherjs struct method functions to reference the `VueJS instance`
func GetVM(structPtr interface{}) *Vue {
	vm, ok := vMap[structPtr]
	if !ok {
		println("Vue not registerd yet:", structPtr)
		panic("Vue not registerd yet")
	}
	return vm
}

// WatchEx using a simpler form to do Vue.$watch
func (v *Vue) WatchEx(expression string, callback func(newVal *js.Object)) (unwatcher func()) {
	obj := v.Call("$watch", expression, callback)
	return func() {
		obj.Invoke()
	}
}
