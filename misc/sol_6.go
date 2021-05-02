/*
sol_6.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

package misc

import (
	"github.com/toransahu/goutils/adt"
)

// Assumption: say K linkedlists are there; for the simplicity of assymptotic calculations each linkedlists are of size N
// Constraint: allowed auxilary space: O(K)

// Approach: do it in `reduce` fashion, merge two linkedlist at one time.. repeat until there is only one linkedlist
// Time complexity: O(N^2) (as the size of the base linkedlist will grow in AP (Arithmatic Progression)
// Auxilary space: nil
func Sol_6_v0(input [][]int) {}

// Approach: delete the heads of each linkedlists and create another linkedlist of those node in sorted order
// Time complexity: O(K^2 * N) = O(K*N) (to delete the K heads N times) X O(K) (to compare with base array while sorting/inserting a head)
// Auxilary space: O(K) (to hold K heads)
func Sol_6_v1(input [][]int) []int {
	// K - number of linkedlists
	k := len(input)
	// a map to hold K linkedlists
	linkedLists := make(map[int]*adt.SLinkedlist, k) // max size would be K
	// iterate/parse the input matrix row by row
	for i, list := range input {
		// create a new linkedlist
		l := adt.NewSLinkedList()
		// iterate through each item in the row
		for _, item := range list {
			// add as a node to the linkedlist
			l.Append(item)
		}
		// add this linkedlist to the  t(holder)
		linkedLists[i] = &l
	}

	// resultant Linkedlist
	resultLL := adt.NewSLinkedList()

	// to store whether all the K linkedlists have been processed or not
	done := 0

	// while all the linkedlist are NON Empty - TC: O(N) * O(K) * O(K)
	for { // O(N)
		if done == k {
			break
		}
		// iterate through each linkedlists
		for i, l := range linkedLists { // O(K)
			if l.IsEmpty() {
				done++
				delete(linkedLists, i)
				break
			}
			// get the head
			head := l.Head
			// delete the head
			err := l.Remove(l.Head)
			if err != nil {
				break
			}
			// insert the head to sorted/result linkedlist
			insertInSortedOrder(&resultLL, head) // O(K)? or O(K*N)? or (K * logN)?
		}
	}

	// an array to return the sorted data
	res := []int{}
	// iterate through the result linkedlist
	for n := range resultLL.Iterate() {
		// append each node's data to the result array
		res = append(res, n.Data.(int))
	}
	return res
}

func insertInSortedOrder(l *adt.SLinkedlist, node *adt.SLLNode) {
	if l.IsEmpty() {
		l.Head = node
		node.Next = nil
		return
	}

	if l.Head.Data.(int) > node.Data.(int) {
		temp := l.Head
		l.Head = node
		node.Next = temp
		return
	}

	curr := l.Head
	for {
		if curr.Next == nil {
			curr.Next = node
			node.Next = nil
			break
		}

		if curr.Next.Data.(int) > node.Data.(int) {
			tmp := curr.Next
			curr.Next = node
			node.Next = tmp
			break
		}

		curr = curr.Next
	}
}

// Approach: delete the heads of each linkedlists and create a BST (Binary Search Tree), then print in preorder traversal fashion
// Time complexity: O(3K*N) = O(K*N) (for traversal of each linkedlists) + O(K*N) (for pre-order traversal of the BST) + O(K*N) (may be for insertion as well)
// Auxilary space: O(K*N)
func Sol_6_v2(input [][]int) {}

// Approach: delete the heads of each linkedlists and create a Min-Heap (priority queue), then pick the min element from the heap and append to the result linkedlist, replace that min element in the heap with with its next node in its linkedlist
// Time complexity: O(K*N)
// Auxilary space: O(K)
func Sol_6_v3(input [][]int) {

}
