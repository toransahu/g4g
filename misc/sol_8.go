/*
sol_8.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package misc
package misc

import "github.com/toransahu/goutils/adt"

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
