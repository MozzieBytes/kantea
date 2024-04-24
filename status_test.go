package main

import "testing"

type wrapTest struct {
	input, want Status
}

var wrapTests = []wrapTest{
	{todo - 1, done},
	{done + 1, todo},
	{inProgress, inProgress},
}

func Test_Wrap(t *testing.T) {
	for _, test := range wrapTests {
		if got := test.input.Wrap(); got != test.want {
			t.Errorf(
				"Input: %q. Result %q did not match expected result %q",
				test.input,
				got,
				test.want,
			)
		}
	}
}
