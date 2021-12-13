/*
sol_19.3.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package sol19
package sol19

import "math"

var prevNodeVal = math.MinInt64

func Sol19_3(node *BSTNode) bool {
	if node == nil {
		return true
	}

	// check the left subtree
	ok := Sol19_3(node.left)
	if !ok {
		return false
	}

	// check the root
	ok = node.data >= prevNodeVal
	if !ok {
		return false
	}
	prevNodeVal = node.data

	// check the right subtree
	ok = Sol19_3(node.right)
	if !ok {
		return false
	}
	return true
}
