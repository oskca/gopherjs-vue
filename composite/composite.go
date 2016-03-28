// Package composite is an higher level wrapper of gopherjs-vue,
// by providing a more gopher friendly API, this package tries to hide
// the JavaScript details for VueJS easy usage in GopherJS world.
package composite

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-vue"
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

func (c *Composite) On(evt LifeCycleEvent, fn interface{}) *Composite {
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
	c.mixins = append(c.mixins, js.M{
		name: val,
	})
	return c
}

// SetData set model data of the genereated VueJS instance (optional)
func (c *Composite) SetData(structPtr interface{}) *Composite {
	c.conf["data"] = structPtr
	c.conf["methods"] = js.MakeWrapper(structPtr)
	return c
}

// SetDataFunc set model data func of the genereated VueJS instance (optional)
func (c *Composite) SetDataFunc(vmCreator func() interface{}) *Composite {
	c.conf["data"] = vmCreator
	return c
}

// SetTemplate set template of the genereated VueJS instance (optional)
func (c *Composite) SetTemplate(elIdOrTemplate string) *Composite {
	c.conf["template"] = elIdOrTemplate
	return c
}

// AddComponent add sub component to the genereated VueJS instance (optional)
func (c *Composite) AddComponent(name string, sub *vue.Vue) *Composite {
	c.mcom[name] = sub
	return c
}

// AddComponent add EventHandler to the genereated VueJS instance (optional)
func (c *Composite) AddEventHandler(name string, fn interface{}) *Composite {
	c.mevt[name] = fn
	return c
}

// Vuerify create the VueJS instance for finally use
// the VueJS instance becomes usable only after this call
func (c *Composite) Vuerify() *vue.Vue {
	return nil
}
