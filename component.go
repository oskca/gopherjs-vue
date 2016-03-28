package vue

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	creatorPool = make([]*pool, 0)
)

type pool struct {
	creator   func() (structPtr interface{})
	structPtr interface{}
	counter   int
}

// Component is actually an Extended Vue SubClass,
// which acts as a Component constructor in VueJS world
//
// CreateComponent creats a new VueJS component
func CreateComponent(
	vmCreator func() (structPtr interface{}),
	templateOrSharpId string,
	replaceMountPoint ...bool,
) *Vue {
	idx := len(creatorPool)
	creatorPool = append(creatorPool, new(pool))
	creatorPool[idx].creator = vmCreator
	vmfn := func() interface{} {
		p := creatorPool[idx]
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
		"data": vmfn,
		"init": js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
			// set methods dynamicly before VueInstance doing all the other init
			this.Get("$options").Set("methods", js.MakeWrapper(vmfn()))
			// register this component instance to vMap
			vMap[vmfn()] = &Vue{Object: this}
			return nil
		}),
		"template": templateOrSharpId,
		"replace":  replace,
	})
	return com
}

// Extend Create a “subclass” of the base Vue constructor which is a `Component`
//  The argument should be an object containing component options.
//  The special cases to note here are el and data options
//  - they must be functions when used with Vue.extend().
func Extend(opt js.M) *Vue {
	return &Vue{
		Object: vue.Call("extend", opt),
	}
}

func RegisterComponent(name string, component *Vue) *Vue {
	vue.Call("component", name, component.Object)
	return component
}

func GetComponent(name string) *Vue {
	return &Vue{
		Object: vue.Call("component", name),
	}
}

// Component creates and registers a named global component. (automatically call Vue.extend)
//
//  vmCreator should return a gopherjs struct pointer. see New for more details
func Component(
	name string,
	vmCreator func() (structPtr interface{}),
	templateOrSharpId string,
	replaceMountPoint ...bool,
) *Vue {
	com := CreateComponent(vmCreator, templateOrSharpId, replaceMountPoint...)
	return RegisterComponent(name, com)
}
