package test_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/illotum/go-test"
)

type mockPrinter struct{}

func (p mockPrinter) Helper()                       {}
func (p mockPrinter) Error(...interface{})          {}
func (p mockPrinter) Errorf(string, ...interface{}) {}

func TestEq(t *testing.T) {
	errA, errB := errors.New("a"), errors.New("b")
	for _, tc := range []struct {
		a, b interface{}
		want bool
	}{
		{1, 1, true},
		{1, 2, false},
		{nil, 2, false},
		{1, nil, false},
		{"a", "a", true},
		{"a", "b", false},
		{nil, "b", false},
		{"a", nil, false},
		{errA, errA, true},
		{errA, errB, false},
		{nil, errB, false},
		{errB, nil, false},
		{nil, nil, true},
	} {
		t.Run(fmt.Sprintf("%#v-%#v", tc.a, tc.b), func(t *testing.T) {
			tt := mockPrinter{}
			have := test.Eq(tt, tc.a, tc.b)
			if have != tc.want {
				t.Errorf("\nhave %+v\nwant %+v", have, tc.want)
			}
		})
	}
}
