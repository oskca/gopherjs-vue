// Package composite is an higher level wrapper of gopherjs-vue,
// by providing a more gopher friendly API, this package tries to hide
// the JavaScript details for VueJS easy usage in GopherJS world.
package composite

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-vue"
)

type LifeCycleEvent string

const (
	EvtInit          LifeCycleEvent = "init"
	EvtCreated       LifeCycleEvent = "created"
	EvtBeforeCompile LifeCycleEvent = "beforeCompile"
	EvtCompiled      LifeCycleEvent = "compiled"
	EvtReady         LifeCycleEvent = "ready"
	EvtAttached      LifeCycleEvent = "attached"
	EvtDetached      LifeCycleEvent = "detached"
	EvtBeforeDestroy LifeCycleEvent = "beforeDestroy"
	EvtDestroyed     LifeCycleEvent = "destroyed"
)

// Composite are used to organize mutiple sub component together to
// construct VueJS apps or (higher level) components.
type Composite struct {
	v    *vue.Vue
	conf js.M
	// map to sub component
	mcom map[string]*vue.Vue
	// map to event handler
	mevt map[string]interface{}
	// properties
	props []string
	// mixins
	mixins []js.M
}

func NewComposite() *Composite {
	c := &Composite{}
	c.conf = make(js.M)
	c.mcom = make(map[string]*vue.Vue, 0)
	c.mevt = make(map[string]interface{}, 0)
	c.props = []string{}
	c.mixins = []js.M{}
	return c
}

// Vuerify create the VueJS instance for finally use
// the VueJS instance becomes usable only after this call
func (c *Composite) Vuerify() *vue.Vue {
	if len(c.mcom) > 0 {
		c.conf["components"] = c.mcom
	}
	if len(c.mevt) > 0 {
		c.conf["events"] = c.mevt
	}
	if len(c.props) > 0 {
		c.conf["props"] = c.props
	}
	if len(c.mixins) > 0 {
		c.conf["mixins"] = c.mixins
	}
	return vue.Extend(c.conf)
}

// The DOM element that the Vue instance is managing.
// Note that for Fragment Instances, vm.$el will return an
// anchor node that indicates the starting position of the fragment.
//  el can be element id or an HTMLElement directly
func (c *Composite) SetEl(el interface{}) *Composite {
	c.conf["el"] = el
	return c
}

func (c *Composite) SetName(name string) *Composite {
	c.conf["name"] = name
	return c
}

// SetData set model data of the genereated VueJS instance (optional)
func (c *Composite) SetData(structPtr interface{}) *Composite {
	c.conf["data"] = structPtr
	c.conf["methods"] = js.MakeWrapper(structPtr)
	return c
}

// SetDataFunc set model data func of the genereated VueJS instance (optional)
// this is only used when you want to create a component
func (c *Composite) SetDataFunc(vmCreator func() interface{}) *Composite {
	c.conf["data"] = vmCreator
	return c
}

// SetTemplate set template of the genereated VueJS instance (optional)
func (c *Composite) SetTemplate(elIdOrTemplate string) *Composite {
	c.conf["template"] = elIdOrTemplate
	return c
}

// Specify the parent instance for the instance to be created.
// Establishes a parent-child relationship between the two.
// The parent will be accessible as this.$parent for the child,
// and the child will be pushed into the parent’s $children array.
func (c *Composite) SetParent(p *vue.Vue) *Composite {
	c.conf["parent"] = p.Object
	return c
}

// Determines whether to replace the element being mounted on
// with the template. If set to false, the template will
// overwrite the element’s inner content without
// replacing the element itself.
func (c *Composite) SetReplace(replace bool) *Composite {
	c.conf["replace"] = replace
	return c
}

func (c *Composite) OnLifeCycleEvent(evt LifeCycleEvent, fn interface{}) *Composite {
	c.addMixin(string(evt), fn)
	return c
}

// The mixins option accepts an array of mixin objects.
// These mixin objects can contain instance options just
// like normal instance objects, and they will be
// merged against the eventual options using the same
// option merging logic in Vue.extend(). e.g.
// If your mixin contains a created hook and
// the component itself also has one, both functions will be called.
//
// Mixin hooks are called in the order they are provided,
// and called before the component’s own hooks.
func (c *Composite) Mixin(val js.M) *Composite {
	c.mixins = append(c.mixins, val)
	return c
}

func (c *Composite) addMixin(name string, val interface{}) *Composite {
	return c.Mixin(js.M{
		name: val,
	})
}

// AddComponent add sub component to the genereated VueJS instance (optional)
func (c *Composite) AddSubComponent(name string, sub *vue.Vue) *Composite {
	c.mcom[name] = sub
	return c
}

// OnComponentEvent add EventHandler to the genereated VueJS instance (optional)
func (c *Composite) OnComponentEvent(name string, fn interface{}) *Composite {
	c.mevt[name] = fn
	return c
}

// AddProperty add props to the genereated VueJS instance (optional)
// 	props is a list/hash of attributes that are exposed to accept data from
// 	the parent component. It has a simple Array-based syntax and
// 	an alternative Object-based syntax that allows advanced configurations
// 	such as type checking, custom validation and default values.
func (c *Composite) AddProperty(name ...string) *Composite {
	c.props = append(c.props, name...)
	return c
}
