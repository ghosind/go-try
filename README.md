# go-try

![test](https://github.com/ghosind/go-try/workflows/test/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/ghosind/go-try)](https://goreportcard.com/report/github.com/ghosind/go-try)
[![codecov](https://codecov.io/gh/ghosind/go-try/branch/main/graph/badge.svg)](https://codecov.io/gh/ghosind/go-try)
![Version Badge](https://img.shields.io/github/v/release/ghosind/go-try)
![License Badge](https://img.shields.io/github/license/ghosind/go-try)
[![Go Reference](https://pkg.go.dev/badge/github.com/ghosind/go-try.svg)](https://pkg.go.dev/github.com/ghosind/go-try)

The `try...catch` statement alternative in Golang.

## Installation

```
go get -u github.com/ghosind/go-try
```

## Getting Started

There is the simplest example to run a function and handle the error in the `catch` function that the `try` function returned.

```go
out, err := try.TryCatch(func () (error) {
  // Do something
  return err
}, func (err error) {
  // The function will not executing if err is nil.
  // Handle error...
})
```

The try functions will return two values. The first value is the all values the `try` function returned, and the second value is the error that the `try` function returns or it panics.

```go
out, err := try.Try(func () (string, error) {
  return "Hello world", errors.New("expected error")
})
fmt.Println(out)
fmt.Println(err)
// [Hello world, expected error]
// expected error
```

The package provides four forms for the `try` statement alternative:

- `Try`
- `TryFinally`
- `TryCatch`
- `TryCatchFinally`

In default cases, the try functions will catch the error that the `try` function panics, and you can set the `CatchPanic` variable to `false` to disable it.

```go
try.CatchPanic = false
try.Try(func () {
  panic("panic error")
})
// panic: expected error
```
