package try

import (
	"errors"
	"fmt"
	"reflect"
)

var errorType reflect.Type = reflect.TypeOf((*error)(nil)).Elem()

// checkAndExecute runs the try function, and returns the outputs of the function. It'll catch the
// panic error if the CatchPanic mode is enabled.
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

// checkFn checks the try function, and it panics if the try parameter is nil or is not a function.
func checkFn(fn any) reflect.Value {
	if fn == nil {
		panic(ErrNilFunction)
	}

	fv := reflect.ValueOf(fn)
	if fv.Kind() != reflect.Func {
		panic(ErrNotFunction)
	}

	return fv
}

// execute runs the try function and returns the output of the function, it'll also set the return
// value error if the last return value of the function is an error.
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
		if last.Kind() == reflect.Interface &&
			lastType.Implements(errorType) &&
			errorType.Implements(lastType) {
			err = last.Interface().(error)
		}
	}

	return ret, err
}
