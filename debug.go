// Package debug contains a bool [Debug] that can be toggled via a build tag.
// To turn on debug mode, use `go build --tags debug ...`.
package debug

import "fmt"

// Err is a non nil error passed to [Error].
type Err struct{ error error }

func (e *Err) Error() string { return fmt.Sprintf("debug: %s", e.error.Error()) }
func (e *Err) Unwrap() error { return e.error }

// Error return the given error if [Debug] is true.
// Returns nil if err is nil or [Debug] is false.
func Error(err error) error {
	if Debug && err != nil {
		return &Err{error: err}
	}
	return nil
}
