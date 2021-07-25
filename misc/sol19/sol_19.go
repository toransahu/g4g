/*
sol_19.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package sol19
package sol19

type BSTNode struct {
	data   int
	left   *BSTNode
	right  *BSTNode
	parent *BSTNode
}

// Sol19v2 finds two swapped nodes - one can swap the data of those nodes to restore the BST.
// Approach: Inorder traversal (recursive) of the tree without extra space.
// Time complexity: O(n)
// Space complexity:
// - O(log n) if we consider the recursion stack
func Sol19v2(root *BSTNode) (*BSTNode, *BSTNode) {
	ctx := &Context{}
	InOrderTraversalv2(ctx, root)

	// fix the swapped nodes
	// if ctx.firstNode != nil && ctx.secondNode != nil {
	//     ctx.firstNode.data, ctx.secondNode.data = ctx.secondNode.data, ctx.firstNode.data
	// }

	return ctx.firstNode, ctx.secondNode
}

func InOrderTraversalv2(ctx *Context, root *BSTNode) {
	if root == nil {
		return
	}

	InOrderTraversalv2(ctx, root.left)

	if ctx.prevNode != nil && root.data < ctx.prevNode.data {

		if ctx.firstNode == nil {
			ctx.firstNode = ctx.prevNode
			ctx.secondNode = root
		} else {
			ctx.secondNode = root
		}

	}
	ctx.prevNode = root

	InOrderTraversalv2(ctx, root.right)
}

type Context struct {
	prevNode   *BSTNode
	firstNode  *BSTNode
	secondNode *BSTNode
}

// Sol19v1 finds two swapped nodes and restore their positions.
// Approach: Inorder traversal (recursive) of the tree with extra space.
// Time complexity: O(n)
// Space complexity:
// - O(n) [to store the inorder'd/sorted nodes]
// - O(log n) if we consider the recursion stack
func Sol19v1(root *BSTNode) (*BSTNode, *BSTNode) {
	arr := &[]*BSTNode{}
	InOrderTraversal(root, arr)

	if len(*arr) <= 1 {
		return nil, nil
	}

	// assuming there are only 2 misplaced/swapped nodes
	var first *BSTNode
	var second *BSTNode

	for i := 1; i < len(*arr); i++ {

		p1 := (*arr)[i-1]
		p2 := (*arr)[i]

		if p1.data <= p2.data {
			continue
		}

		if p2.data < p1.data {
			if first == nil {
				first = p1
				second = p2
			} else {
				second = p2
			}
		}
	}

	// fix the swapped nodes
	// if first != nil && second != nil {
	//     first.data, second.data = second.data, first.data
	// }

	return first, second
}

func InOrderTraversal(parent *BSTNode, arr *[]*BSTNode) {
	if parent.left == nil && parent.right == nil {
		*arr = append(*arr, parent)
		return
	}

	if parent.left != nil {
		InOrderTraversal(parent.left, arr)
	}
	*arr = append(*arr, parent)
	if parent.right != nil {
		InOrderTraversal(parent.right, arr)
	}
}
