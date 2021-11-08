/*
sol_31.1.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package misc
package misc

import (
	"fmt"
	"math"

	"github.com/toransahu/goutils/adt"
)

func Sol31_1(str, pattern string) string {

	// if pattern is lengthier than the string, then no solution
	if len(pattern) > len(str) {
		return ""
	}

	// lets have two pointers
	// one pointing to start of the current window under examination
	startPointer := 0
	// other pointing to end of the window
	endPointer := 0

	// lets track the size of a possible minimal sized window so far
	soFarMinWindowSize := math.Inf(1)
	// also keep track of the start & end position contributing to the minimalWindow
	startOfWindow := 0
	endOfWindow := 0

	// to store the pattern in calculative manner
	patternMap := map[string]int{}

	// to store the count of required alphabet (given in pattern)
	strMap := map[string]int{}

	// initialize the maps
	for _, _rune := range pattern {
		ch := fmt.Sprintf("%c", _rune)
		patternMap[ch] = 0
		strMap[ch] = 0
	}

	// populate the patternMap with exact counts
	for _, r := range pattern {
		ch := fmt.Sprintf("%c", r)
		patternMap[ch]++
	}

	// lets track the requirements of the pattern - as the exact match is not required, thus order could differ
	matchesPattern := func() bool {
		for key := range patternMap {
			ch := key
			if strMap[ch] < patternMap[ch] {
				return false
			}
		}
		return true
	}

	inPattern := func(ch string) bool {
		// ch := fmt.Sprintf("%c", r)
		if patternMap[ch] > 0 {
			return true
		}
		return false
	}

	// a queue to store the seq in which we found the matching window
	q := adt.NewQueue()
	// info to store on that queue
	type Info struct {
		ch  string
		pos int
	}

	// lets start iterating the given string from beginning till we find a pattern matching window in the given str
	firstMatch := true
	for i, r := range str {
		// if pattern found early, then stop iterating
		if matchesPattern() {
			break
		}

		ch := fmt.Sprintf("%c", r)
		if firstMatch && inPattern(ch) {
			startPointer = i
			firstMatch = false
			q.Enqueue(&Info{
				ch:  ch,
				pos: i,
			})
			strMap[ch]++
			continue
		}

		if inPattern(ch) {
			// assign endPointer
			endPointer = i
			q.Enqueue(&Info{
				ch:  ch,
				pos: i,
			})
			strMap[ch]++
		}
	}

	// if the given string does not contain the pattern
	if !matchesPattern() {
		// then return an empty window
		return ""
	}

	// update the soFarMinWindowSize
	currWindowSize := float64(endPointer - startPointer + 1)
	if currWindowSize < soFarMinWindowSize {
		startOfWindow = startPointer
		endOfWindow = endPointer
		soFarMinWindowSize = currWindowSize
	}

	// soFarMinWindowSize = math.Min(float64(soFarMinWindowSize), float64(endPointer-startPointer+1))

	// lets start iterating the found window from beginning, and pop the first char & try to find the same in the rest of the string

	for {

		// keep doing it until there is no more possibility
		if endPointer >= len(str) {
			break
		}

		infoI, err := q.Dequeue()
		if err != nil {
			return str[startOfWindow : endOfWindow+1]
			// panic(err)
		}
		info := infoI.(*Info)
		ch := info.ch
		if q.IsEmpty() {
			return str[startOfWindow : endOfWindow+1]
		}
		nextInfo := (*q)[0].(*Info) // its fine, queue is not going to be empty anytime

		// if we have more than sufficient amount of that char, then no need to search; just point to the next candidate
		if strMap[ch] > patternMap[ch] {
			startPointer = nextInfo.pos // then next one
			strMap[ch]--
			// and update the soFarMinWindowSize
			// soFarMinWindowSize = math.Min(float64(soFarMinWindowSize), float64(endPointer-startPointer+1))
			currWindowSize := float64(endPointer - startPointer + 1)
			if currWindowSize < soFarMinWindowSize && matchesPattern() {
				startOfWindow = startPointer
				endOfWindow = endPointer
				soFarMinWindowSize = currWindowSize
			}
		} else {
			// loop over the str, start after endOfWindow
			j := endPointer + 1
			for i := j; i < len(str); i++ {
				nextCh := fmt.Sprintf("%c", str[i])
				endPointer = i

				// if the char is what we're looking for
				if ch == nextCh {
					q.Enqueue(&Info{
						ch:  nextCh,
						pos: i,
					})

					startPointer = nextInfo.pos // then next one
					currWindowSize := float64(endPointer - startPointer + 1)
					if currWindowSize < soFarMinWindowSize && matchesPattern() {
						startOfWindow = startPointer
						endOfWindow = endPointer
						soFarMinWindowSize = currWindowSize
					}

					break
				}

				// if the char is in pattern
				if inPattern(nextCh) {
					q.Enqueue(&Info{
						ch:  nextCh,
						pos: i,
					})

					strMap[nextCh]++
					// strMap[ch]--
				}

				// no solution case
				if i == len(str)-1 {
					return str[startOfWindow : endOfWindow+1] // TODO: +1?
				}
			}
		}
	}
	return ""
}
