/*
sol_3_test.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package misc
package misc

import "testing"

func TestSol3(t *testing.T) {
	testcases := []struct {
		given string
		want  string
	}{
		{given: "-9-999", want: "VALID"},
		{given: "-1-1-22-4411", want: "VALID"},
		{given: "-11-11", want: "VALID"},
		{given: "-1-1-22-441", want: "INVALID"},
		{given: "1111", want: "INVALID"},
		{given: "-01-01", want: "INVALID"},
	}

	for _, testcase := range testcases {

		input := testcase.given
		want := testcase.want
		result := Sol3(input)
		if result != want {
			t.Errorf("for input: %v , want: %v, got: %v", input, want, result)
		}
	}
}
