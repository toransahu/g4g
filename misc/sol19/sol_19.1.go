/*
sol_19.1.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package sol19
package sol19

var k int
var counter int = 0
var kth *int

func Sol19_1(node *BSTNode, K int) int {

	kth = nil
	counter = 0
	k = K
	sol19_1(node)
	return *kth
}

func sol19_1(node *BSTNode) {
	if node == nil {
		return
	}

	// early return
	if kth != nil {
		return
	}

	// traverse the right subtree
	sol19_1(node.right)

	// check if the current
	counter++
	if counter == k {
		kth = &node.data
		return
	}

	// traverse the left subtree
	sol19_1(node.left)
}
