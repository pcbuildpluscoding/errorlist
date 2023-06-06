package errorlist

// -------------------------------------------------------------- //
// NewErrorlist
// ---------------------------------------------------------------//
func NewErrorlist(failFast bool) Errorlist {
  elist := Errorlist{
    this: []error{},
    failFast: failFast,
  }
  return elist
}
