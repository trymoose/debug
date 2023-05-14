// Package debug contains a bool [Debug] that can be toggled via a build tag.
// To turn on debug mode, use `go build --tags debug ...`.
package debug

import (
	"fmt"
	"runtime"
)

// Err is a non nil error passed to [Error].
type Err struct {
	File  string
	Line  uint
	error error
}

func (e *Err) Unwrap() error { return e.error }
func (e *Err) Error() string {
	return fmt.Sprintf("[DEBUG] { %s:%d }: %s", e.File, e.Line, e.error.Error())
}

// Error return the given error if [Debug] is true.
// Returns nil if err is nil or [Debug] is false.
func Error(err error) error {
	if Debug && err != nil {
		_, fn, ln, _ := runtime.Caller(1)
		return &Err{File: fn, Line: uint(ln), error: err}
	}
	return nil
}
