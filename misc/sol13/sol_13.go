/*
sol_13.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package sol13
package sol13

// Sol13_v3 solves this problem using dynamic programming approach (bottom-up/tabulation approach - smaller to larger).
// Time complexity: O(n)
// Space complexity: O(n)
func Sol13_v3(n int) int {
	// init the table with base knowledge
	table := map[int]int{
		0: 0, // assume there is no way to take 0 steps
		1: 1,
		2: 2,
		3: 4,
	}

	// lets start populating the table for unknowns (till `n`)
	for i := 4; i <= n; i++ {
		table[i] = table[i-1] + table[i-2] + table[i-3]

	}
	// return the value of `n` from the table
	return table[n]
}

// Sol13_v2 solves this problem by using recurrence relation
// Time complexity: O(3^^n)
// Space complexity: O(n) (for inbuilt stack?)
func Sol13_v2(n int) int {

	// base case
	if n < 1 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if n == 3 {
		return 4
	}

	// else try solve the recurrence relation
	return Sol13_v2(n-1) + Sol13_v2(n-2) + Sol13_v2(n-3)
}

// Sol13 solves this problem by taking n-ary tree approach - more verbose style.
// Time complexity: O(3^^n)
// Space complexity: O(n)
func Sol13(n int) int {
	// init the no. of possible ways
	ctx := &Context{ways: 0, N: n}
	// lets start approaching the top of the stairs from ground
	root := NewNode(0)
	// from ground, try to explore all 3 possible ways (i.e. take 1 step, 2 step, or 3 step)
	Explore(ctx, root)

	// return the possible no. ways
	return ctx.ways
}

// Explore explores all the possible (based on problem statement) path from the given position/node.
// Time complexity:
// Space complexity:
func Explore(ctx *Context, node *Node) {
	// if the given node/position already lead us to the top of the stair
	if node.sumVal == ctx.N {
		// then stop the exploration
		return
	}

	// else start exploring by trying step-1,2 and 3 from current position
	for i := 1; i < 4; i++ {
		// select a next step
		child := NewNode(i)
		// if that step seems useful (i.e. may lead us to top of the stair)
		if node.addChild(ctx, child) {
			// explore that further
			Explore(ctx, child)
		}
	}
}

type Context struct {
	ways int
	N    int
}

type Node struct {
	val      int
	children []*Node
	sumVal   int
}

func NewNode(val int) *Node {
	return &Node{
		val:      val,
		children: []*Node{},
		sumVal:   val,
	}
}

// addChild adds a given node as a child to the current node - iff adding it makes any sense - and returns true if child gets added
// Time complexity: O(1)
// Space complexity: O(1)
func (n *Node) addChild(ctx *Context, node *Node) bool {
	// add the accumulated sum of parent to child's
	node.sumVal += n.sumVal
	// append the child to the list (optional)
	n.children = append(n.children, node)
	// if this child/action leads us to reach the top of the stair
	if node.sumVal == ctx.N {
		// increment the 'ways' count
		ctx.ways++
	}

	// if this child/action is useful
	if node.sumVal <= ctx.N {
		return true
	} else {
		return false
	}
}
