/*
sol_8.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package misc
package misc

import "github.com/toransahu/goutils/adt"

// Sol8_v2 solves the problem by using Floyd's cycle detection method.
// Approach: maintain 2 pointers, 1 slow (1 jump at a iteration), 1 fast (2 jumps at a iteration) - if both the pointers points to the same node, then there is a loop
// Idea:
func Sol8_v2(l *adt.SLinkedlist) (bool, *adt.SLinkedlist) {
	if l.IsEmpty() || l.Head.Next == nil {
		return false, l
	}
	if l.Head.Next == l.Head {
		l.Head.Next = nil
		return true, l
	}

	pS := l.Head
	pF := l.Head
	var meetingPoint *adt.SLLNode

	for {
		if pF.Next == nil || pF.Next.Next == nil {
			return false, l
		}

		if pS.Next == pF.Next.Next {
			meetingPoint = pS.Next
			break
		}

		pS = pS.Next
		pF = pF.Next.Next
	}

	// if we're here, then we've detected a loop in the linkedlist

	// check if the linkedlist is _circular_
	if meetingPoint == l.Head {
		// make it non-circular
		pF.Next = nil
		return true, l
	}

	// if we're here then the linkedlist a _loop_ and we need to find the start node of the loop

	pS = l.Head
	pF = meetingPoint

	for {
		if pS.Next == pF.Next {
			break
		}
		pS = pS.Next
		pF = pF.Next
	}

	pF.Next = nil
	return true, l
}

// Sol8 solves the problem by hashing the visited node's address and matching the current node's address against the hash map.
// TC: O(n)
// SC: O(n)
func Sol8(l *adt.SLinkedlist) bool {

	if l.IsEmpty() {
		return false
	}

	visited := map[*adt.SLLNode]bool{} // O(n)
	node := l.Head
	visited[node] = true

	for { // O(n)
		next := node.Next
		if next == nil {
			break
		}

		if visited[next] {
			return true
		}

		visited[next] = true
		node = node.Next
	}

	return false
}
