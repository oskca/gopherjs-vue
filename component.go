package vue

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	creatorPool = make(map[string]*pool)
)

type pool struct {
	creator   func() (structPtr interface{})
	structPtr interface{}
	counter   int
}

// NewComponent registers a global component. (automatically call Vue.extend)
//
// vmCreator should return a gopherjs struct pointer. see New for more details
func NewComponent(name string, vmCreator func() (structPtr interface{}), templateOrSharpId string, replaceMountPoint ...bool) *Component {
	creatorPool[name] = new(pool)
	creatorPool[name].creator = vmCreator
	creator := func() interface{} {
		p := creatorPool[name]
		if p.counter%3 == 0 {
			p.structPtr = p.creator()
		}
		p.counter += 1
		return p.structPtr
	}
	// args
	replace := true
	if len(replaceMountPoint) > 0 {
		replace = replaceMountPoint[0]
	}
	// create the VueInstance
	com := Extend(js.M{
		"data": creator,
		"init": js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
			// set methods dynamicly before VueInstance doing all the other init
			this.Get("$options").Set("methods", js.MakeWrapper(creator()))
			// register this component instance to vMap
			vMap[creator()] = &Vue{Object: this}
			return nil
		}),
		"template": templateOrSharpId,
		"replace":  replace,
	})
	// register component
	RegisterComponent(name, com)
	return com
}

// Component is actually an Extended Vue SubClass,
// which acts as a Component constructor in VueJS world
type Component struct {
	*Vue
}

// Extend Create a “subclass” of the base Vue constructor. Which is a `Component`
//  The argument should be an object containing component options.
//  The special cases to note here are el and data options
//  - they must be functions when used with Vue.extend().
func Extend(opt js.M) *Component {
	return &Component{
		Vue: &Vue{
			Object: vue.Call("extend", opt),
		},
	}
}

func RegisterComponent(name string, component *Component) {
	vue.Call("component", name, component.Object)
}

func GetComponent(name string) *Component {
	return &Component{
		Vue: &Vue{
			Object: vue.Call("component", name),
		},
	}
}
