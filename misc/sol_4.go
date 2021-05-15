/*
sol_4.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package misc
package misc

import (
	"strconv"
	"strings"

	"github.com/toransahu/goutils/adt"
)

func Sol4(input []string) []int {
	firstLine := input[0]
	remainingLines := input[1:]

	numOfTasks, err := strconv.Atoi(firstLine)
	if err != nil {
		panic(err)
	}

	dependencies := [][]int{}
	for _, line := range remainingLines {
		dep := strings.Split(line, "<")
		u, err := strconv.Atoi(dep[0])
		if err != nil {
			panic(err)
		}
		v, err := strconv.Atoi(dep[1])
		if err != nil {
			panic(err)
		}
		dependencies = append(dependencies, []int{u, v})
	}

	g := adt.NewGraph(numOfTasks)
	for _, dep := range dependencies {
		g.AddEdge(dep[0], dep[1])
	}
	result, hasCycle := g.TopoSort()
	if hasCycle {
		return []int{}
	}
	return result
}
