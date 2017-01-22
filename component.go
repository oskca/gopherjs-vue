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
// thus you can use Component.New to create a
// preConfigured VueJS instance(*ViewModel).
type Component struct {
	*ViewModel
}

// New create the component instance
func (c *Component) New() *ViewModel {
	return newViewModel(c.Object.New())
}

func newComponent(o *js.Object) *Component {
	return &Component{
		ViewModel: newViewModel(o),
	}
}

// Register register Component:c in the global namespace
func (c *Component) Register(name string) *Component {
	vue.Call("component", name, c)
	return c
}

func GetComponent(name string) *Component {
	return newComponent(vue.Call("component", name))
}

// NewComponent creates and registers a named global Component
//
//  vmCreator should return a gopherjs struct pointer. see New for more details
func NewComponent(
	vmCreator func() (structPtr interface{}),
	templateStr string,
	replaceMountPoint ...bool,
) *Component {
	// args
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
	// opts
	opt := NewOption()
	opt.Data = vmfn
	opt.Template = templateStr
	opt.OnLifeCycleEvent(EvtBeforeCreate, func(vm *ViewModel) {
		vm.Options.Set("methods", js.MakeWrapper(vmfn()))
		vMap[vmfn()] = vm
	})
	return opt.NewComponent()
}
