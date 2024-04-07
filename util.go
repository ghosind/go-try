package try

import (
	"errors"
	"fmt"
	"reflect"
)

var errorType reflect.Type = reflect.TypeOf((*error)(nil)).Elem()

func checkAndExecute(fn any) (out []any, err error) {
	fv := checkFn(fn)

	if CatchPanic {
		defer func() {
			e := recover()
			if e == nil {
				return
			}
			switch t := e.(type) {
			case error:
				err = t
			case string:
				err = errors.New(t)
			default:
				err = fmt.Errorf("%v", t)
			}
		}()
	}

	out, err = execute(fv)

	return
}

func checkFn(fn any) reflect.Value {
	if fn == nil {
		panic("try function is required")
	}

	fv := reflect.ValueOf(fn)
	if fv.Kind() != reflect.Func {
		panic("try must be a function")
	}

	return fv
}

func execute(fv reflect.Value) ([]any, error) {
	var err error
	ret := make([]any, 0, fv.Type().NumOut())

	out := fv.Call(nil)
	for _, v := range out {
		ret = append(ret, v.Interface())
	}

	if len(out) > 0 {
		last := out[len(out)-1]
		lastType := last.Type()
		if last.Kind() == reflect.Interface && lastType.Implements(errorType) && errorType.Implements(lastType) {
			err = last.Interface().(error)
		}
	}

	return ret, err
}
