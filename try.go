package try

var CatchPanic bool = true

func Try(try any) ([]any, error) {
	return tryCatch(try, nil, nil)
}

func TryFinally(try any, finally func()) ([]any, error) {
	return tryCatch(try, nil, finally)
}

func TryCatch(try any, catch func(error)) ([]any, error) {
	return tryCatch(try, catch, nil)
}

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
