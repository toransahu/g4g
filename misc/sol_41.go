/*
sol_41.go
Copyright (C) 2021 Toran Sahu <toran.sahu@yahoo.com>

Distributed under terms of the MIT license.
*/

// Package misc
package misc

/*
Sol41 breaks the problem in 2 cases/parts and returns the result.
e.g.

s=eacdaba
  0123456

i	s[i]	suffixes ending with s[i]					strength of the suffixes/substrings
-------------------------------------------------------------------------------------------
0	e		[e]											1								 1
1	a		[ea, a]										2 + 1							 3
2	c		[eac, ac, c]								3 + 2 + 1						 6
3	d		[eacd, acd, cd, d]							4 + 3 + 2 + 1					10
4	a		[eacda, acda, cda, da, a]					4 + 3 + 3 + 2 + 1				13
5	b		[eacdab, acdab, cdab, dab, ab, b]			5 + 4 + 4 + 3 + 2 + 1			19
6	a		[eacdaba, acdaba, cdaba, daba, aba, ba, a]	5 + 4 + 4 + 3 + 2 + 2 + 1		21
																					   -----
																						73
*/
func Sol41(password string) int {
	// total strength of the password
	strength := 0
	// strength calculated so far
	strength_from_prev_char := 0
	// a map to store the last occurrence & position of characters in the password
	var last_index_map map[rune]int = map[rune]int{}
	var strength_from_current_char int

	// iterate through the password
	for i, r := range password {
		// handle first char
		if i == 0 {
			strength_from_current_char = 1
			strength++                                           // increment the strength by 1
			strength_from_prev_char = strength_from_current_char // update the strength so far
			last_index_map[r] = i                                // store the position
			continue                                             // handle the second char
		}

		// CASE-1: IF THE CURRENT CHAR IS THE REPEATING ONE
		if last_index, ok := last_index_map[r]; ok {
			// calculate additional strength we get by including the current char
			// i.e. strength_from_current_char := strength so far + strength due to # of chars between last occurrence and current occurrence of the char + strength of current char (i.e. 1)
			strength_from_current_char = strength_from_prev_char + (i - last_index - 1) + 1
			strength = strength + strength_from_current_char
			strength_from_prev_char = strength_from_current_char
			last_index_map[r] = i
			continue
		}

		// CASE-2: IF THE CURRENT CHAR IS NON-REPEATING ONE
		// i.e. strength_from_current_char := strength from previous char addition + strength due to # of chars between  last occurrence (which was never, thus from index=0) and current occurrence of the char + strength of current char (i.e. 1)
		strength_from_current_char = strength_from_prev_char + i + 1
		strength = strength + strength_from_current_char
		strength_from_prev_char = strength_from_current_char
		last_index_map[r] = i

	}
	return strength
}
