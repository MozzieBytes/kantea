package main

import "testing"

type statusTest struct {
	input, want Status
}

var progressTests = []statusTest{
	{todo, inProgress},
	{inProgress, done},
	{done, todo},
}

func Test_Progress(t *testing.T) {
	for _, test := range progressTests {
		if got := test.input.Progress(); got != test.want {
			t.Errorf(
				"Input: %d. Result %d did not match expected result %d",
				test.input,
				got,
				test.want,
			)
		}
	}
}

var regressTests = []statusTest{
	{todo, done},
	{inProgress, todo},
	{done, inProgress},
}

func Test_Regress(t *testing.T) {
	for _, test := range regressTests {
		if got := test.input.Regress(); got != test.want {
			t.Errorf(
				"Input: %d. Result %d did not match expected result %d",
				test.input,
				got,
				test.want,
			)
		}
	}
}
