package try

// CatchPanic indicates whether the functions will catch the error that the try function panics or
// not.
var CatchPanic bool = true

// Try executes the try function, and returns the error that the try function returned. It will
// also catch the panic error if it runs under CatchPanic mode.
func Try(try any) ([]any, error) {
	return tryCatch(try, nil, nil)
}

// TryFinally executes the try function, and it executes the finally function after the try
// function is finished.
func TryFinally(try any, finally func()) ([]any, error) {
	return tryCatch(try, nil, finally)
}

// TryCatch executes the try function, and if it returns an error or panics under CatchPanic mode,
// the catch function will be executed.
func TryCatch(try any, catch func(error)) ([]any, error) {
	return tryCatch(try, catch, nil)
}

// TryCatchFinally executes the try function, and if it returns an error or panics under
// CatchPanic mode, the catch function will be executed. The finally function will always be
// executed except the catch function panics.
func TryCatchFinally(try any, catch func(error), finally func()) ([]any, error) {
	return tryCatch(try, catch, finally)
}

func tryCatch(try any, catch func(error), finally func()) ([]any, error) {
	out, err := checkAndExecute(try)
	if err != nil && catch != nil {
		catch(err)
	}

	if finally != nil {
		finally()
	}

	return out, err
}
