package mapping

import (
	"github.com/Archs/js/JSON"
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-vue"
	"strings"
)

// FromJS set the corresponding VueJS data model field from obj
func FromJS(v *vue.Vue, obj *js.Object) *vue.Vue {
	for _, key := range js.Keys(v.Object) {
		// skip internal/unexported field
		if strings.HasPrefix(key, "$") || strings.HasPrefix(key, "_") {
			continue
		}
		val := obj.Get(key)
		if val == js.Undefined {
			continue
		}
		v.Set(key, val)
	}
	return v
}

func FromJSON(v *vue.Vue, jsonStr string) *vue.Vue {
	return FromJS(v, JSON.Parse(jsonStr))
}

func ToJS(v *vue.Vue) *js.Object {
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

func ToJSON(v *vue.Vue) string {
	return JSON.Stringify(ToJSON(v))
}
