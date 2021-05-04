/*
sol_7.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package misc
package misc

import (
	"github.com/toransahu/goutils/adt"
)

type Node struct {
	i        int
	j        int
	distance int
}

func NewNode(i int, j int, distance int) *Node {
	return &Node{i: i, j: j, distance: distance}
}

// Approach: Treat the cells of the field as node of a directed graph, where the unsafe cells will be omited out from the possible nodes, and safe + reachable nodes (w.r.t. possible movements from a node) will be considered as adjacent nodes; Once a graph with all safe node is ready i.e. input matrix it self as adjacency matrix of the graph - run BFS from any cell of the first col - do it recursively; while doing do check if the visited node falls in the last column.
// Time complexity: O(M*N)
// Auxilary space: O(M*N) for the queue
func Sol7(input [][]int) int {

	m := len(input) // num of rows in [][]int
	n := 0
	if m > 0 {
		n = len(input[0])
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if input[i][j] == 0 {
				// iterate through adjacent nodes/cells
				for _, pos := range getAdjacentCellPos(i, j, m, n) {
					x := pos[0]
					y := pos[1]
					// mark adjacent nodes possibly unsafe (using -1)
					input[x][y] = -1
				}
			}
		}
	}

	// now find the shortest path (only distance) from nodes at first level to nodes at last level
	// in a un-weighted/equal-weighted directed graph

	// queue for BFS
	q := adt.NewQueue()

	// array/list to track visited nodes during BFS
	visited := make([][]bool, m)
	for idx := 0; idx < m; idx++ {
		visited[idx] = make([]bool, n)
	}

	// iterate through first column i.e in MxN iterate input[0, 0], input[1, 0], ... input[M, 0]
	for i := 0; i < m; i++ {
		// enqueue the nodes with distance from source as zero
		q.Enqueue(NewNode(i, 0, 0))
	}

	// while queue in not empty
	for {
		// if the queue is empty
		if q.IsEmpty() {
			// end the search of a destination node
			// no solution exists
			break
		}

		nodeI, _ := q.Dequeue() // get the first node at the left from the queue
		node := nodeI.(*Node)

		// if this node falls in the last column in the matrix
		if node.j == n-1 {
			// then this is the shortest distance from level 1 to level N
			return node.distance
		}

		// else run BFS for its adjacent nodes - to where traversal/movement is possible
		for _, pos := range getPossibleAdjacentNodes(node.i, node.j, m, n) {
			if isSafeNode(pos[0], pos[1], input) {
				visited[node.i][node.j] = true // mark this node as visited
				// enqueue the node
				q.Enqueue(NewNode(pos[0], pos[1], node.distance+1))
			}
		}
	}

	return -1 // if no solution exists
}

func getCellPos(nodeNum int, width int) (int, int) {
	row := int(nodeNum / width)
	col := nodeNum % width
	return row, col
}

// getAdjacentCellPos returns all the adjacent cells's position/location w.r.t the given node position
func getAdjacentCellPos(i int, j int, m int, n int) [][]int {
	adj := [][]int{}
	var (
		tl_x, t_x, tr_x = i - 1, i - 1, i - 1
		l_x, r_x        = i, i
		bl_x, b_x, br_x = i + 1, i + 1, i + 1

		tl_y, l_y, bl_y = j - 1, j - 1, j - 1
		t_y, b_y        = j, j
		tr_y, r_y, br_y = j + 1, j + 1, j + 1
	)

	x := []int{tl_x, t_x, tr_x, l_x, r_x, bl_x, b_x, br_x}
	y := []int{tl_y, t_y, tr_y, l_y, r_y, bl_y, b_y, br_y}

	for idx := 0; idx < 8; idx++ {
		if isValidPos(x[idx], y[idx], m, n) {
			adj = append(adj, []int{x[idx], y[idx]})
		}
	}
	return adj
}

// getPossibleAdjacentNodes returns possible adjacent node positions where movement (of 1 step) is possible
func getPossibleAdjacentNodes(i int, j int, m int, n int) [][]int {
	adjacents := [][]int{}
	// left, top, bottom, right
	x := []int{i, i - 1, i + 1, i}
	y := []int{j - 1, j, j, j + 1}

	for idx := 0; idx < 4; idx++ {
		if isValidPos(x[idx], y[idx], m, n) {
			adjacents = append(adjacents, []int{x[idx], y[idx]})
		}
	}
	return adjacents
}

// isValidPos tells if given i,j coordinates lies in an 2-d array of size MxN or not
func isValidPos(i int, j int, m int, n int) bool {
	if i < m && i >= 0 && j < n && j >= 0 {
		return true
	}
	return false
}

func isSafeNode(i int, j int, input [][]int) bool {
	isSafe := input[i][j] == 1
	return isSafe
}
