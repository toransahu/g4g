/*
sol_12.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package misc
package misc

// Sol12_v2 returns the median of the streaming integers.
// Approach: By storing the streaming numbers in two different data structure. One min-heap and another max-heap. Then on-demand calculates the median
// Time complexity: O(log n)
// Space complexity: O(n)

func Sol12_v2(n int, returnMedian bool) {

}

// Sol_12 returns the median of the streaming integers.
// Approach: By storing the streaming numbers in a sorted order (say min-heap implemented using array/list) and then on-demand calculates the median
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
