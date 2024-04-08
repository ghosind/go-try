package try

import "errors"

var (
	ErrNilFunction = errors.New("try function is required")
	ErrNotFunction = errors.New("try must be a function")
)
