/*
sol_12.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package misc
package misc

import (
	"github.com/toransahu/goutils/adt/heap"
	"github.com/toransahu/goutils/adt/heap/maxheap"
)

// Sol12_v2 returns the median of the streaming integers.
// Approach: By storing the streaming numbers in two different data structure. One min-heap and another max-heap. Then on-demand calculates the median
// Time complexity: O(log n) - for any intermediate median calculation
// Space complexity: O(n)
func Sol12_v2(ctx *ContextSol12v2, n int, returnMedian bool) float32 {

	/*
		Q. Why 2 heaps and not single heap?
		A. Because we need middle and half of the time 2 middle items in the sorted list, which a heap (represented using array) is not capable of telling.

		Min-heap for: storing all the items more than the effective median
		Max-heap for: storing all the items less than the effective median

		e.g.


		n			max		  median			min
		-------------------------------------------

		3			 3			3


		0			 0			1.5				 3


		4			 0			 3				 3
												/
											   4


		2			 2			 2.5		     3
					/							/
				   0						   4
	*/

	if ctx.sizeLeftHeap == 0 && ctx.sizeRightHeap == 0 {
		InsertToLeftHeap(ctx, n)
		ctx.median = float32(n)
		return ctx.median
	}

	// if the integer is less than the effective median AND size of the left heap is greater than the right heap
	if float32(n) <= ctx.median && ctx.sizeLeftHeap > ctx.sizeRightHeap {
		/*
			move the top of left heap to right heap
			and insert the given number to left heap
		*/

		// replace top of left heap with the given number
		leftTop := maxheap.Replace(ctx.leftHeap, n)

		// insert the replaced item to right heap AND increase the size
		InsertToRightHeap(ctx, leftTop)
	} else if float32(n) > ctx.median && ctx.sizeLeftHeap < ctx.sizeRightHeap {
		/*
			move the top of right heap to left heap
			and insert the given number to right heap
		*/

		// replace top of right heap with the given number
		rightTop := heap.Replace(ctx.rightHeap, n)

		// insert the replaced item to left heap AND increase the size
		InsertToLeftHeap(ctx, rightTop)
	} else if float32(n) > ctx.median {
		// insert the integer to right heap AND increase the size
		InsertToRightHeap(ctx, n)
	} else if float32(n) <= ctx.median {
		// insert the integer to right heap AND increase the size
		InsertToLeftHeap(ctx, n)
	}

	// calculate median - check even/odd
	ctx.median = CalcMedian(ctx)
	return ctx.median
}

type ContextSol12v2 struct {
	leftHeap      *heap.IntArray // max-heap to store the numbers less than effective median
	sizeLeftHeap  int            // size of the left heap
	rightHeap     *heap.IntArray // min-heap to store the numbers more than effective median
	sizeRightHeap int            // size of the right heap
	median        float32        // median of the received integers
}

func NewContextSol12v2() *ContextSol12v2 {
	return &ContextSol12v2{
		leftHeap:      &heap.IntArray{},
		sizeLeftHeap:  0,
		rightHeap:     &heap.IntArray{},
		sizeRightHeap: 0,
	}
}

func CalcMedian(ctx *ContextSol12v2) float32 {
	if ctx.sizeLeftHeap == ctx.sizeRightHeap {
		return float32(IntVal(maxheap.Top(ctx.leftHeap))+IntVal(heap.Top(ctx.rightHeap))) / 2
	} else if ctx.sizeLeftHeap > ctx.sizeRightHeap {
		return float32(IntVal(maxheap.Top(ctx.leftHeap)))
	} else {
		return float32(IntVal(heap.Top(ctx.rightHeap)))
	}
}

func InsertToLeftHeap(ctx *ContextSol12v2, n interface{}) {
	maxheap.Insert(ctx.leftHeap, IntVal(n))
	ctx.sizeLeftHeap++
}

func InsertToRightHeap(ctx *ContextSol12v2, n interface{}) {
	heap.Insert(ctx.rightHeap, IntVal(n))
	ctx.sizeRightHeap++
}

func IntVal(i interface{}) int {
	return i.(int)
}

// Sol_12 returns the median of the streaming integers.
// Approach: By storing the streaming numbers in a sorted order (say min-heap implemented using array/list i.e. insertion sort) and then on-demand calculates the median
// Time complexity: O(n)
// Space complexity: O(n)
func Sol_12(ctx *Context, n int, returnMedian bool) float32 {

	ctx.size++

	// get or create a list
	if ctx.l == nil {
		ctx.l = NewList()
	}

	// append in sorted order
	ctx.l = ctx.l.InsertInSortedOrder(n)

	var m float32 // median of the list of the numbers received
	if ctx.size == 1 {
		m = float32(n)
	} else if ctx.size == 2 {
		m = float32((*ctx.l)[(ctx.size/2)-1]+n) / 2
	} else if ctx.size%2 == 0 {
		m = float32((*ctx.l)[(ctx.size/2)-1]+(*ctx.l)[ctx.size/2]) / 2
	} else {
		m = float32((*ctx.l)[(ctx.size-1)/2.0])
	}
	if returnMedian {
		return m
	}
	return -1.0
}

type List []int

func NewList() *List {
	return &List{}
}

func (l *List) InsertInSortedOrder(n int) *List {
	if l == nil {
		l = NewList()
	}
	if len(*l) == 0 {
		*l = append(*l, n)
		return l
	}

	nIsLargest := false
	for i, val := range *l {
		if val <= n {
			if i == len(*l)-1 {
				nIsLargest = true
			}
			continue
		}

		// increase the length thus capacity by appending a dummy element
		*l = append(*l, -1)
		// shift the elements starting from index 2 by 1 position
		copy((*l)[i+1:], (*l)[i:]) // len(src) > len(dst) thus -1 will be omitted out
		(*l)[i] = n
		break
	}
	if nIsLargest {
		*l = append(*l, n)
	}
	return l
}

type Context struct {
	l    *List // sorted list of the received numbers
	size int   // size of the list
}

func NewContext() *Context {
	return &Context{
		l:    nil,
		size: 0,
	}
}
