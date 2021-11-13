/*
sol_41_test.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package misc
package misc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol41(t *testing.T) {

	tcs := []struct {
		given string
		want  int
	}{
		{"test", 19},
		{"good", 16},
		{"eacdaba", 73},
	}

	for _, tc := range tcs {
		got := Sol41(tc.given)
		assert.Equal(t, tc.want, got)
	}
}
