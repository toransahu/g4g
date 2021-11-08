/*
sol_31.1_test.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package misc
package misc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol31_1(t *testing.T) {
	tcs := []struct {
		given []string
		want  string
	}{
		{[]string{"this is a test string.", "tist"}, "t stri"},
	}

	for _, tc := range tcs {
		str := tc.given[0]
		pattern := tc.given[1]
		got := Sol31_1(str, pattern)

		assert.Equal(t, tc.want, got)
	}
}
