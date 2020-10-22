package test

import (
	"reflect"
	"testing"
)

func Eq(t testing.TB, have, want interface{}) bool {
	t.Helper()
	eq := reflect.DeepEqual(have, want)
	if !eq {
		t.Errorf("\nhave %+v\nwant %+v", have, want)
	}
	return eq
}
