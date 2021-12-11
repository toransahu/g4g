/*
sol_18.2_test.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package misc
package misc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol18_2(t *testing.T) {
	tt := []struct {
		given []string
		want  map[string][]string
	}{
		{[]string{"abc", "cde", "bca", "bac"}, map[string][]string{"abc": {"bca", "bac"}, "cde": {}}},
	}

	for _, tc := range tt {
		got := Sol18_2(tc.given)
		assert.Equal(t, tc.want, got)
	}
}
