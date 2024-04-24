package main

import "testing"

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
