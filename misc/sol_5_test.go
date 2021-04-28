/*
sol_5_test.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

package misc

import "testing"

func TestSol5(t *testing.T) {
	testcases := []struct {
		given [][]int
		want  int
	}{
		{given: [][]int{
			{0, 0, 1, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 0},
			{0, 0, 1, 0}},
			want: 2,
		},
		{given: [][]int{
			{0, 0},
			{0, 0},
		},
			want: -1,
		},
		{given: [][]int{
			{0, 1},
			{1, 0},
		},
			want: -1,
		},
		{given: [][]int{
			{0},
		},
			want: -1,
		},
	}

	for _, testcase := range testcases {
		got := Sol5_V1(len(testcase.given), testcase.given)
		want := testcase.want

		if got != want {
			t.Errorf("for input: %v , want: %v, got: %v", testcase.given, want, got)
		}
	}
}
