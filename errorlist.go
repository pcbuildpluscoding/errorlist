package errorlist

import (
	"errors"
	"fmt"
)

// ==================================================================//
// Errorlist
// ==================================================================//
type Errorlist struct {
	this     []error
	failFast bool
}

// -------------------------------------------------------------- //
// Add
// ---------------------------------------------------------------//
func (el *Errorlist) Add(err error) {
	if err != nil {
		el.this = append(el.this, err)
		if el.failFast {
			panic(err)
		}
	}
}

// ----------------------------------------------------------------//
// Addf
// ----------------------------------------------------------------//
func (el *Errorlist) Addf(format string, args ...interface{}) {
	el.Add(fmt.Errorf(format, args...))
}

// -------------------------------------------------------------- //
// AddMiss
// ---------------------------------------------------------------//
func (el *Errorlist) AddMiss(ok bool, format string, args ...interface{}) {
	if !ok {
		el.Add(fmt.Errorf(format, args...))
	}
}

// -------------------------------------------------------------- //
// Empty
// ---------------------------------------------------------------//
func (el Errorlist) Empty() bool {
	return len(el.this) == 0
}

// -------------------------------------------------------------- //
// Error
// ---------------------------------------------------------------//
func (el *Errorlist) Error() string {
	if el.Empty() {
		return ""
	}
	return el.unwrap().Error()
}

// -------------------------------------------------------------- //
// HasOnly
// ---------------------------------------------------------------//
func (el *Errorlist) Is(errKind error) bool {
	for _, err := range el.this {
		if errors.Is(err, errKind) {
			return true
		}
	}
	return false
}

// -------------------------------------------------------------- //
// Reset
// ---------------------------------------------------------------//
func (el *Errorlist) Reset() error {
	err := el.Unwrap()
	el.this = []error{}
	return err
}

// -------------------------------------------------------------- //
// SetFailMode
// ---------------------------------------------------------------//
func (el *Errorlist) SetFailMode(failfast bool) {
	el.failFast = failfast
}

// -------------------------------------------------------------- //
// unwrap
// ---------------------------------------------------------------//
func (el *Errorlist) unwrap() error {
	var caterr string
	last := len(el.this) - 1
	for i, err := range el.this {
		if i == last {
			caterr += err.Error()
		} else {
			caterr += fmt.Sprintf("%s\n", err.Error())
		}
	}
	return errors.New(caterr)
}

// -------------------------------------------------------------- //
// Unwrap
// ---------------------------------------------------------------//
func (el *Errorlist) Unwrap() error {
	if !el.Empty() {
		return el.unwrap()
	}
	return nil
}
