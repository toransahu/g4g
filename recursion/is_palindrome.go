//
// is_palindrome.go
// Copyright (C) 2020 Toran Sahu <toran.sahu@yahoo.com>
//
// Distributed under terms of the MIT license.
//

package main

func is_palindrome(s string) bool {

	if len(s) == 0 {
		return false
	}

	start := 0
	end := len(s) - 1

	if start == end {
		return true
	}

	return s[start] == s[end] && is_palindrome(s[start+1:end])
}

// func main() {
//     s := "nitin"
//     fmt.Println(is_palindrome(s))
// }
