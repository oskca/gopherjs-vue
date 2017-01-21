package vue

import (
	"strings"

	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-json"
)

// FromJS set the corresponding VueJS data model field from obj
// new data model field will be created when not exist
func (v *ViewModel) FromJS(obj *js.Object) *ViewModel {
	for _, key := range js.Keys(obj) {
		// skip internal or unexported field
		if strings.HasPrefix(key, "$") || strings.HasPrefix(key, "_") {
			continue
		}
		v.Object.Set(key, obj.Get(key))
	}
	return v
}

func (v *ViewModel) FromJSON(jsonStr string) *ViewModel {
	return v.FromJS(json.Parse(jsonStr))
}

func (v *ViewModel) ToJS() *js.Object {
	obj := js.Global.Get("Object").New()
	for _, key := range js.Keys(v.Object) {
		// skip internal/unexported field
		if strings.HasPrefix(key, "$") || strings.HasPrefix(key, "_") {
			continue
		}
		obj.Set(key, v.Get(key))
	}
	return obj
}

func (v *ViewModel) ToJSON() string {
	return json.Stringify(v.ToJSON())
}
