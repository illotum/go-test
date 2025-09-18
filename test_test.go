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

func TestNeq(t *testing.T) {
	errA, errB := errors.New("a"), errors.New("b")
	for _, tc := range []struct {
		a, b interface{}
		want bool
	}{
		{1, 1, false},
		{1, 2, true},
		{nil, 2, true},
		{1, nil, true},
		{"a", "a", false},
		{"a", "b", true},
		{nil, "b", true},
		{"a", nil, true},
		{errA, errA, false},
		{errA, errB, true},
		{nil, errB, true},
		{errB, nil, true},
		{nil, nil, false},
	} {
		t.Run(fmt.Sprintf("%#v-%#v", tc.a, tc.b), func(t *testing.T) {
			tt := mockPrinter{}
			have := test.Neq(tt, tc.a, tc.b)
			if have != tc.want {
				t.Errorf("\nhave %+v\nwant %+v", have, tc.want)
			}
		})
	}
}

func TestNil(t *testing.T) {
	errA := errors.New("a")
	for _, tc := range []struct {
		a    interface{}
		must bool
		want bool
	}{
		{1, false, true},
		{1, true, false},
		{errA, false, true},
		{errA, true, false},
		{"", false, true},
		{"", true, false},
		{nil, false, false},
		{nil, true, true},
	} {
		t.Run(fmt.Sprintf("%#v-%#v", tc.a, tc.must), func(t *testing.T) {
			tt := mockPrinter{}
			have := test.Nil(tt, tc.a, tc.must)
			if have != tc.want {
				t.Errorf("\nhave %+v\nwant %+v", have, tc.want)
			}
		})
	}
}
