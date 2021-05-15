/*
sol_4_test.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package misc
package misc

import (
	"reflect"
	"testing"
)

func TestSol4(t *testing.T) {
	testcases := []struct {
		given []string
		want  []int
	}{
		{
			given: []string{"5", "0<3", "3<4", "4<1", "2<1", "2<0"},
			want:  []int{2, 0, 3, 4, 1},
		},
		{
			given: []string{"5", "0<3", "2<1", "3<4", "2<0", "2<3", "0<4", "0<1"},
			want:  []int{2, 0, 1, 3, 4},
		},
		{
			given: []string{"6", "2<3", "3<1", "4<0", "4<1", "5<2", "5<0"},
			want:  []int{5, 4, 2, 3, 1, 0},
		},
		{
			given: []string{"3", "0<1", "1<2", "2<0"},
			want:  []int{}, // cycle detected, so no tasks
		},
	}

	for _, testcase := range testcases {

		input := testcase.given
		want := testcase.want
		result := Sol4(input)
		if !reflect.DeepEqual(result, want) {
			t.Errorf("for input: %v , want: %v, got: %v", input, want, result)
		}
	}
}
