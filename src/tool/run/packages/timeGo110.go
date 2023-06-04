// +build go1.10

package packages

import (
	"reflect"
	"time"

	"github.com/surdeus/goblin/src/tool/run/env"
)

func timeGo110() {
	env.Packages["time"]["LoadLocationFromTZData"] = reflect.ValueOf(time.LoadLocationFromTZData)
}
