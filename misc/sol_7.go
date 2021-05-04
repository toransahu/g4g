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
	n := 0          // num of columns - initially 0
	// if given matrix is not empty
	if m > 0 {
		n = len(input[0])
	}

	// iterate through each row of the matrix - O(M*N)
	for i := 0; i < m; i++ {
		// iterate through each column of the matrix
		for j := 0; j < n; j++ {
			// if the given cell has sensor implanted
			if input[i][j] == 0 {
				// then iterate through adjacent nodes/cells as well
				for _, pos := range getAdjacentCellPos(i, j, m, n) {
					x := pos[0]
					y := pos[1]
					// mark the adjacent node possibly unsafe (using -1)
					input[x][y] = -1
				}
			}
		}
	}

	// now find the shortest path (only distance) from nodes at first level to nodes at last level
	// in an un-weighted/equal-weighted directed graph

	// queue for BFS
	q := adt.NewQueue()

	// a matrix to track visited nodes during BFS
	visited := make([][]bool, m)
	// initialize all the positions with `false`
	for idx := 0; idx < m; idx++ {
		visited[idx] = make([]bool, n)
	}

	// iterate through first column i.e in MxN matrix, iterate like input[0, 0], input[1, 0], ... input[M, 0]
	for i := 0; i < m; i++ {
		// enqueue the nodes with distance from source as zero
		q.Enqueue(NewNode(i, 0, 0))
	}

	// while queue in not empty - O(M*N)
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

	// top_left, top, top_right, left, right, bottom_left, bottom, bottom_right
	x := []int{tl_x, t_x, tr_x, l_x, r_x, bl_x, b_x, br_x}
	y := []int{tl_y, t_y, tr_y, l_y, r_y, bl_y, b_y, br_y}

	// iterate 8 times - as 8 adjacents are possible (but may not be valid)
	for idx := 0; idx < 8; idx++ {
		// if the coordinate actually lies in a matrix of M*N
		if isValidPos(x[idx], y[idx], m, n) {
			// then append it to the list
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

	// iterate 4 times - as 4 adjacents movements are possible (but may not be valid)
	for idx := 0; idx < 4; idx++ {
		// if the coordinate actually lies in a matrix of M*N
		if isValidPos(x[idx], y[idx], m, n) {
			// then append it to the list
			adjacents = append(adjacents, []int{x[idx], y[idx]})
		}
	}
	return adjacents
}

// isValidPos tells if given i,j coordinates lies in an 2-d array of size MxN or not
func isValidPos(i int, j int, m int, n int) bool {
	// if the coordinate actually lies in a matrix of M*N
	if i < m && i >= 0 && j < n && j >= 0 {
		// then its a valid coordinate
		return true
	}
	return false
}

// isSafeNode returns is the cell at given coordinates is safe as per problem statement or not
func isSafeNode(i int, j int, input [][]int) bool {
	// if val is 1 (in another word, val is not 0, neither -1)
	return input[i][j] == 1
}
