package try_test

import (
	"errors"
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-try"
)

func TestTry(t *testing.T) {
	a := assert.New(t)
	expectedErr := errors.New("expected error")

	out, err := try.Try(func() string {
		return "Hello World"
	})
	a.NilNow(err)
	a.EqualNow(out, []any{"Hello World"})

	out, err = try.Try(func() (string, error) {
		return "Hello World", expectedErr
	})
	a.EqualNow(err, expectedErr)
	a.EqualNow(out, []any{"Hello World", expectedErr})

	out, err = try.Try(func() {
		panic(expectedErr)
	})
	a.EqualNow(err, expectedErr)
	a.EqualNow(out, []any{})
}

func TestTryFinally(t *testing.T) {
	a := assert.New(t)
	expectedErr := errors.New("expected error")
	finally := false

	out, err := try.TryFinally(func() string {
		return "Hello World"
	}, func() {
		finally = true
	})
	a.NilNow(err)
	a.EqualNow(out, []any{"Hello World"})
	a.TrueNow(finally)

	finally = false
	out, err = try.TryFinally(func() (string, error) {
		return "Hello World", expectedErr
	}, func() {
		finally = true
	})
	a.EqualNow(err, expectedErr)
	a.EqualNow(out, []any{"Hello World", expectedErr})
	a.TrueNow(finally)

	finally = false
	out, err = try.TryFinally(func() {
		panic(expectedErr)
	}, func() {
		finally = true
	})
	a.TrueNow(finally)
	a.EqualNow(err, expectedErr)
	a.EqualNow(out, []any{})
}

func TestTryCatch(t *testing.T) {
	a := assert.New(t)
	expectedErr := errors.New("expected error")
	caught := false

	out, err := try.TryCatch(func() string {
		return "Hello World"
	}, func(err error) {
		caught = true
	})
	a.NilNow(err)
	a.EqualNow(out, []any{"Hello World"})
	a.NotTrueNow(caught)

	caught = false
	out, err = try.TryCatch(func() (string, error) {
		return "Hello World", expectedErr
	}, func(err error) {
		a.EqualNow(err, expectedErr)
		caught = true
	})
	a.EqualNow(err, expectedErr)
	a.EqualNow(out, []any{"Hello World", expectedErr})
	a.TrueNow(caught)

	caught = false
	out, err = try.TryCatch(func() {
		panic(expectedErr)
	}, func(err error) {
		a.EqualNow(err, expectedErr)
		caught = true
	})
	a.EqualNow(err, expectedErr)
	a.EqualNow(out, []any{})
	a.TrueNow(caught)
}

func TestTryCatchFinally(t *testing.T) {
	a := assert.New(t)
	expectedErr := errors.New("expected error")
	caught := false
	finally := false

	out, err := try.TryCatchFinally(func() string {
		return "Hello World"
	}, func(err error) {
		caught = true
	}, func() {
		finally = true
	})
	a.NilNow(err)
	a.EqualNow(out, []any{"Hello World"})
	a.NotTrueNow(caught)
	a.TrueNow(finally)

	caught = false
	finally = false
	out, err = try.TryCatchFinally(func() (string, error) {
		return "Hello World", expectedErr
	}, func(err error) {
		a.EqualNow(err, expectedErr)
		caught = true
	}, func() {
		finally = true
	})
	a.EqualNow(err, expectedErr)
	a.EqualNow(out, []any{"Hello World", expectedErr})
	a.TrueNow(caught)
	a.TrueNow(finally)

	caught = false
	finally = false
	out, err = try.TryCatchFinally(func() {
		panic(expectedErr)
	}, func(err error) {
		a.EqualNow(err, expectedErr)
		caught = true
	}, func() {
		finally = true
	})
	a.EqualNow(err, expectedErr)
	a.EqualNow(out, []any{})
	a.TrueNow(caught)
	a.TrueNow(finally)
}

func TestInvalidTry(t *testing.T) {
	a := assert.New(t)

	a.PanicOfNow(func() {
		try.Try(nil)
		a.FailNow()
	}, try.ErrNilFunction)
	a.PanicOfNow(func() {
		try.Try("not a function")
		a.FailNow()
	}, try.ErrNotFunction)
}

func TestTryPanic(t *testing.T) {
	a := assert.New(t)
	expectedErr := errors.New("expected error")

	_, err := try.Try(func() {
		panic(expectedErr)
	})
	a.EqualNow(err, expectedErr)

	_, err = try.Try(func() {
		panic("expected error")
	})
	a.NotNilNow(err)
	a.EqualNow(err.Error(), "expected error")

	_, err = try.Try(func() {
		panic(123)
	})
	a.NotNilNow(err)
	a.EqualNow(err.Error(), "123")

	// Turn off catch panic mode
	try.CatchPanic = false
	a.PanicOfNow(func() {
		try.Try(func() {
			panic(expectedErr)
		})
		a.FailNow()
	}, expectedErr)
	// reset catch panic mode
	try.CatchPanic = true
}
