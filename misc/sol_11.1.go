/*
sol_11.1.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package misc
package misc

import (
	"github.com/toransahu/goutils/adt"
)

// Sol11_1 removes duplicate nodes from the unsorted linkedlist by WITHOU using auxilary space
// Approach: start sorting the linkedlist using quick/merge sort & while doing so, discard the duplicate node
func Sol11_1(l *adt.SLinkedlist) *adt.SLLNode {
	return MergeSort(l.Head)
}

// MergeSort sorts the LL using merge-sort technique & while doing so, it removes the duplicate nodes;
// and returns the head of the sorted LL
func MergeSort(head *adt.SLLNode) *adt.SLLNode {
	if head == nil || head.Next == nil {
		return head
	}

	middle := middleNode(head)
	nextToMiddle := middle.Next
	middle.Next = nil
	left := MergeSort(head)
	right := MergeSort(nextToMiddle)
	// TODO check left/right nil?
	sorted := merge(left, right)
	return sorted
}

// middleNode finds & returns the middle node of the linkedlist
func middleNode(head *adt.SLLNode) *adt.SLLNode {
	// this logic is already handled in MergeSort
	// if head == nil || head.Next == nil {
	// 	return nil
	// }

	slowP := head
	fastP := head

	// util fast pointer reaches the tail
	for {
		if fastP.Next == nil || fastP.Next.Next == nil {
			break
		}

		fastP = fastP.Next.Next // advance by 2 step
		slowP = slowP.Next      // advance by 1 step
	}
	return slowP
}

// merge merges the sorted left- & right-linkedlist using their given heads & returns the head of the merged LL
func merge(leftH *adt.SLLNode, rightH *adt.SLLNode) *adt.SLLNode {
	var head *adt.SLLNode
	if leftH == nil && rightH == nil {
		return nil
	} else if leftH == nil {
		return rightH
	} else if rightH == nil {
		return leftH
	}

	leftP := leftH
	rightP := rightH

	head = leftP

	if leftP.Data.(int) == rightP.Data.(int) {
		rightP = rightP.Next
	} else if leftP.Data.(int) > rightP.Data.(int) {
		tmp := rightP.Next
		rightP.Next = leftP
		leftP = rightP
		head = leftP

		rightP = tmp
	}
	// while we reach the tail/end (node.next == nil)
	for {
		// if any of the LL have NO more nodes to visit
		if leftP.Next == nil || rightP == nil {
			// then break
			break
		}

		if leftP.Next.Data.(int) == rightP.Data.(int) {
			// if data is same
			// increment the right pointer
			rightP = rightP.Next
		} else if leftP.Next.Data.(int) < rightP.Data.(int) {
			// if current item in leftP LL is less than current item in rightP LL
			// then increment the left pointer
			leftP = leftP.Next
		} else {
			// else if curr node in right LL is less than curr node in left LL
			// then shift+fit that right LL node to left LL
			tmp := rightP.Next

			rightP.Next = leftP.Next
			leftP.Next = rightP

			rightP = tmp
		}
	}
	if leftP.Next == nil {
		// then append all the nodes from the other LL
		leftP.Next = rightP
	} else if rightP == nil {
		// we're trying to move things from rightP LL to leftP LL; so nothing to move; so do nothing
	}

	return head
}
