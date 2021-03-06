package nul // import "gopkg.in/webnice/lin.v1/nl"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"fmt"
	"reflect"
	"strconv"
)

func asString(src interface{}) string {
	const floatFormat = 'g'
	var rv reflect.Value

	switch v := src.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	}
	rv = reflect.ValueOf(src)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(rv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(rv.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(rv.Float(), floatFormat, -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(rv.Float(), floatFormat, -1, 32)
	case reflect.Bool:
		return strconv.FormatBool(rv.Bool())
	}

	return fmt.Sprintf("%v", src)
}
