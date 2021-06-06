/*
sol_11.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package misc
package misc

import "github.com/toransahu/goutils/adt"

// Sol11 removes duplicate nodes from the unsorted linkedlist by using auxilary space
// Approach: iterate over linkedlist & store the visited node values in a hash map if its absent in the map; if the node is found in the hash map, then that's a duplicate
// Time Complexity: O(N)
// Space Complexity: O(N)
func Sol11(l *adt.SLinkedlist) {
	// TODO
}
