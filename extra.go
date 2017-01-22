package vue

import (
	"github.com/gopherjs/gopherjs/js"
)

// TConfig is used for declaration only
type TConfig struct {
	*js.Object
	// Suppress all Vue logs and warnings.
	Silent bool `js:"silent"`
	// The merge strategy receives the value of that option defined on the parent and child instances as the first and second arguments, respectively.
	OptionMergeStrategies interface{} `js:"optionMergeStrategies"`
	// Configure whether to allow vue-devtools inspection.
	Devtools bool `js:"devtools"`
	// Assign a handler for uncaught errors during component render and watchers.
	ErrorHandler func(err, vm *js.Object) `js:"errorHandler"`
	// Make Vue ignore custom elements defined outside of Vue (e.g., using the Web Components APIs).
	IgnoredElements []string `js:"ignoredElements"`
	// Define custom key alias(es) for v-on.
	KeyCodes map[string]int `js:"keyCodes"`
}

// Vue.extend( options )
// Arguments:
// 	{Object} options
// 	Usage:
// Create a “subclass” of the base Vue constructor. The argument should be an object containing component options.
// The special case to note here is the data option - it must be a function when used with Vue.extend().
func Extend(o *Option) *Component {
	vm := vue.Call("extend", o.prepare())
	return &Component{
		&ViewModel{
			Object: vm,
		},
	}
}

// Defer the callback to be executed after the next DOM update cycle.
// Use it immediately after you’ve changed some data to wait for the DOM update.
func NextTick(cb func()) {
	vue.Call("nextTick", cb)
}

// Vue.set( object, key, value )
//
// Arguments:
//
//  {Object} object
//  {String} key
//  {*} value
//  Returns: the set value.
//
// Set a property on an object. If the object is reactive,
// ensure the property is created as a reactive property and
// trigger view updates. This is primarily used to get
// around the limitation that Vue cannot detect property additions.
func Set(obj, key, value interface{}) {
	vue.Call("set", obj, key, value)
}

// Vue.delete( object, key )
//
// Arguments:
//
//  {Object} object
//  {String} key
//  Usage:
//
// Delete a property on an object.
// If the object is reactive, ensure the deletion triggers view updates.
// This is primarily used to get around the limitation that
// Vue cannot detect property deletions, but you should rarely need to use it.
func Delete(obj, key interface{}) {
	vue.Call("delete", obj, key)
}

// Vue.use( mixin )
//
// Arguments:
//
// 	{Object | Function} plugin
// 	Usage:
//
// Install a Vue.js plugin. If the plugin is an Object, it must expose an install method. If it is a function itself, it will be treated as the install method. The install method will be called with Vue as the argument.
// When this method is called on the same plugin multiple times, the plugin will be installed only once.
func Use(plugin interface{}) {
	vue.Call("use", plugin)
}

// Vue.mixin( mixin )
//
// Arguments:
//
// 	{Object} mixin
// 	Usage:
//
// Apply a mixin globally, which affects every Vue instance created afterwards. This can be used by plugin authors to inject custom behavior into components. Not recommended in application code.
func Mixin(mixin interface{}) {
	vue.Call("mixin", mixin)
}

// Vue.compile( template )
//
// Arguments:
//
// 	{string} template
// 	Usage:
//
// Compiles a template string into a render function. Only available in the standalone build.
func Compile(template string) (renderFn *js.Object) {
	return vue.Call("compile", template).Get("render")
}

var Config = &TConfig{
	Object: js.Global.Get("Object").New(),
}
