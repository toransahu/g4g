//
// factorial.go
// Copyright (C) 2020 Toran Sahu <toran.sahu@yahoo.com>
//
// Distributed under terms of the MIT license.
//

package main

func factorial(n int) int {
	if n <= 0 {
		return 1
	}

	return n * factorial(n-1)
}

// func main() {
//     fmt.Println(factorial(4))
// }
