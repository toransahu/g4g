/*
sol_5.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// In a party of N people, only one person is known to everyone. Such a person may be present in the party, if yes, (s)he doesn’t know anyone in the party. We can only ask questions like “does A know B? “. Find the stranger (celebrity) in the minimum number of questions.

package misc

import (
	"github.com/toransahu/goutils/adt"
)

// Sol5_V1 solves the problem with:
// Approach: Linear with a stack of N people
// Time Complexity: O(N)
// Space Complexity: O(N)
func Sol5_V1(numOfPeople int, input [][]int) int {
	// if there is no or one person in the party
	if numOfPeople <= 1 {
		// then no one is celebrity
		return -1
	}

	// create a stack to hold N people
	stack := adt.NewStack()

	// put all N people in the stack
	for person := 0; person < numOfPeople; person++ {
		stack.Push(person)
	}

	// while the stack is not empty
	for {
		isEmpty, err := stack.IsEmpty()
		if err != nil {
			panic(err)
		}
		if isEmpty {
			break
		}

		// pop a person from top of the stack
		A, err := stack.Pop()
		if err != nil {
			panic(err)
		}
		personA := A.(int)

		// try pop another person from top of the stack
		B, err := stack.Pop()
		// if the stack was empty
		if err != nil {
			// then A was the only person left in the stack and is the celebrity
			return personA
		}
		personB := B.(int)

		aKnowsB := input[personA][personB] == 1 // whether A knows B?
		bKnowsA := input[personB][personA] == 1 // whether B knows A?

		// if A & B knows each other
		if aKnowsB && bKnowsA {
			// then both of them are not celebrity; continue the search
			continue
			// if A knows B & A does not know B
		} else if aKnowsB {
			// then B might be the celebrity; let's treat him as celebrity candidate
			// and push back to the stack
			stack.Push(personB)
			// if B knows A & B does not know A
		} else if bKnowsA {
			// then A might be the celebrity; let's treat him as celebrity candidate
			// and push back to the stack
			stack.Push(personA)
		}
	}

	return -1
}
