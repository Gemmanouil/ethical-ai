package main

import (
	"fmt"
)

func main() {
	// List of example strings to test
	words := []string{
		"racecar",
		"hello",
		"A man a plan a canal Panama",
		"apostopa",
	}

	// Go through each string in the list
	for _, w := range words {

		// Build a cleaned version: only A–Z and a–z, all lowercase
		clean := ""
		for i := 0; i < len(w); i++ {
			ch := w[i]

			// Check if it's a letter
			if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') {
				// Convert uppercase to lowercase
				if ch >= 'A' && ch <= 'Z' {
					ch = ch + 32
				}
				clean += string(ch)
			}
		}

		// Check if the cleaned string is a palindrome
		ok := true
		for i := 0; i < len(clean)/2; i++ {
			if clean[i] != clean[len(clean)-1-i] {
				ok = false
				fmt.Println("The palidrome breaks in position:", i+1)
				break
			}
		}

		// Print the result and the cleaned string
		fmt.Println(ok, "|", clean)
	}
}
