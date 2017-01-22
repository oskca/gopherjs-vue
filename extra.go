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

// Vue.partial( id, [definition] )
// id String
// definition String | Node optional
// Register or retrieve a global partial.
// The definition can be a template string, a querySelector that starts with #,
// a DOM element (whose innerHTML will be used as the template string),
// or a DocumentFragment.
func Partial(name, definition string) {
	vue.Call("partial", name, definition)
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

var Config = &TConfig{}

func init() {
	Config.Object = vue.Get("config")
}
