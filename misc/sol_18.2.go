/*
sol_18.2.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package misc
package misc

import "sort"

func Sol18_2(list []string) map[string][]string {

	// to hold the group of anagrams
	anagrams := map[string][]string{}

	for _, str := range list {
		slice := []rune(str)
		sort.SliceStable(slice, func(i int, j int) bool { return slice[i] < slice[j] })
		strSorted := string(slice)
		if val, ok := anagrams[strSorted]; ok {
			val = append(val, str)
			anagrams[strSorted] = val
		} else {
			anagrams[strSorted] = []string{}
		}
	}
	return anagrams
}
