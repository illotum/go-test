// Package test implements simple equality and diff routines for unit testing
// in Go.
package test

import (
	"errors"
	"reflect"

	"github.com/google/go-cmp/cmp"
)

func Eq(t ErrorPrinter, have, want interface{}) bool {
	t.Helper()
	eq := eq(have, want)
	if !eq {
		t.Errorf("\nhave %+v\nwant %+v", have, want)
	}
	return eq
}

func Neq(t ErrorPrinter, have, want interface{}) bool {
	t.Helper()
	neq := !eq(have, want)
	if !neq {
		t.Errorf("\nhave %+v\nwant !%+v", have, want)
	}
	return neq
}

func Nil(t ErrorPrinter, have interface{}, must bool) bool {
	t.Helper()
	if must {
		return Eq(t, have, nil)
	}
	return Neq(t, have, nil)
}

func eq(have, want interface{}) bool {
	haveE, okHaveE := have.(error)
	wantE, okWantE := want.(error)
	var eq bool
	switch {
	case okHaveE && okWantE:
		eq = errors.Is(haveE, wantE)
	default:
		eq = reflect.DeepEqual(have, want)
	}
	return eq
}

func Diff(t ErrorPrinter, have, want interface{}) bool {
	t.Helper()
	eq := cmp.Equal(have, want)
	if !eq {
		t.Error(cmp.Diff(want, have))
	}
	return eq
}

type ErrorPrinter interface {
	Helper()
	Error(...interface{})
	Errorf(string, ...interface{})
}
