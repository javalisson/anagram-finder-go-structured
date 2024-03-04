package main

import (
	"fmt"
	"sandbox/anagramfinder" // Update this import path based on your GOPATH
)

func main() {
	words := []string{"created", "catered", "reacted", "decimal", "claimed", "hello", "echo", "choice"}
	anagramGroups := anagramfinder.FindAnagrams(words)
	for _, group := range anagramGroups {
		fmt.Println(group)
	}
}
