package errorlist

import (
  "errors"
  "fmt"
)

//==================================================================//
// Errorlist
//==================================================================//
type Errorlist struct {
  this []error
  failFast bool
}

// -------------------------------------------------------------- //
// Add
// ---------------------------------------------------------------//
func (el *Errorlist) Add(err error) error {
  if err != nil {
    el.this = append(el.this, err)
    if el.failFast {
      return el.Unwrap()
    }
  }
  return nil
}

//----------------------------------------------------------------//
// Addf
//----------------------------------------------------------------//
func (el *Errorlist) Addf(format string, args ...interface{}) error {
  return el.Add(fmt.Errorf(format, args...))
}

// -------------------------------------------------------------- //
// AddMiss
// ---------------------------------------------------------------//
func (el *Errorlist) AddMiss(ok bool, format string, args ...interface{}) error {
  if !ok {
    return el.Add(fmt.Errorf(format, args...))
  }
  return nil
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
  var caterr string
  last := len(el.this) - 1
  for i, err := range el.this {
    if i == last {
      caterr += fmt.Sprintf("%s", err.Error())
    } else {
      caterr += fmt.Sprintf("%s\n", err.Error())
    }
  }
  return caterr
}

// -------------------------------------------------------------- //
// HasOnly
// ---------------------------------------------------------------//
func (el *Errorlist) HasOnly(ekind error) bool {
  for _, err := range el.this {
    if !errors.Is(err, ekind) {
      return false
    }
  }
  return true
}

// -------------------------------------------------------------- //
// Reset
// ---------------------------------------------------------------//
func (el *Errorlist) Reset() {
  el.this = []error{}
}

// -------------------------------------------------------------- //
// SetFailMode
// ---------------------------------------------------------------//
func (el *Errorlist) SetFailMode(failfast bool) {
  el.failFast = failfast
}


// -------------------------------------------------------------- //
// Unwrap
// ---------------------------------------------------------------//
func (el *Errorlist) Unwrap() error {
  if ! el.Empty() {
    return errors.New(el.Error())
  }
  return nil 
}