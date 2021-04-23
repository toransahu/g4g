/*
sol_3.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

package misc

import (
	"fmt"
	"strconv"

	"github.com/toransahu/goutils/adt"
)

/*
3. Given an expression string expression of parenthesis ( assume numbers from 0 to 9 each represent a type of parenthesis), write a program to examine whether the pairs and the orders are valid in expressions.

'-9-999' -> valid

'-1-1-22-4411' -> valid
*/

func Sol3(input string) string {
	// create an empty stack
	var s adt.Stack = adt.NewStack()
	// flag to check negative numbers
	var isNegativeNum bool = false

	// iterate through the input string
	for _, rune := range input {
		// convert `rune` to `char`
		ch := fmt.Sprintf("%c", rune)
		// if current char is "-"
		if ch == "-" {
			// then treat digit next to/after it as a negative number
			// i.e. an open bracket equivalent
			isNegativeNum = true
			continue // nothing to do next; process the next char
		}
		// convert the char to integer
		item, err := strconv.Atoi(ch)
		// if current item in the input string is a negative number
		if isNegativeNum {
			// then push it to the stack
			s.Push(item)
			isNegativeNum = false // reset the flag
			continue              // nothing to do next; process the next char
		}

		// if we're here, we've encountered a positive number
		// fetch the top element at the stack
		top, err := s.Top()
		if err != nil {
			return "INVALID"
		}
		// if the current item/number the couter part of the element at top of the stack
		if top.(int)-item == 0 {
			// remove the element from top of the stack
			s.Pop()
		} else {
			return "INVALID"
		}
	}
	// after processing all the digits in the given input, check if stack is empty or not
	isEmpty, err := s.IsEmpty()

	if err != nil {
		panic(err)
	}

	// if the stack is empty
	if isEmpty {
		// then the pairs & orders of the digits in the given string were VALID
		return "VALID"
	}
	// else the pairs & orders of the digits in the given string were INVALID
	return "INVALID"
}
