// Package test implements simple equality and diff routines for unit testing
// in Go.
package test

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Eq(t testing.TB, have, want interface{}) bool {
	t.Helper()
	eq := reflect.DeepEqual(have, want)
	if !eq {
		t.Errorf("\nhave %+v\nwant %+v", have, want)
	}
	return eq
}

func Diff(t testing.TB, have, want interface{}) bool {
	t.Helper()
	eq := cmp.Equal(have, want)
	if !eq {
		t.Error(cmp.Diff(want, have))
	}
	return eq
}
