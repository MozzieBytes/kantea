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

func Test_FilterValue(t *testing.T) {
	tsk := Task{title: "want"}
	if got := tsk.FilterValue(); got != "want" {
		t.Errorf(
			"Getter value %q did not match underlying field value %q",
			got,
			"want",
		)
	}
}

func Test_Title(t *testing.T) {
	tsk := Task{title: "want"}
	if got := tsk.Title(); got != "want" {
		t.Errorf(
			"Getter value %q did not match underlying field value %q",
			got,
			"want",
		)
	}
}

func Test_Description(t *testing.T) {
	tsk := Task{description: "want"}
	if got := tsk.Description(); got != "want" {
		t.Errorf(
			"Getter value %q did not match underlying field value %q",
			got,
			"want",
		)
	}
}
