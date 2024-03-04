package anagramfinder

import (
	"sort"
	"strings"
)

// NormalizeWord returns a string with the letters of the word sorted alphabetically
func NormalizeWord(word string) string {
	chars := strings.Split(word, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

// FindAnagrams finds all anagrams in the given list of words
func FindAnagrams(words []string) [][]string {
	anagramsMap := make(map[string][]string)
	for _, word := range words {
		normalized := NormalizeWord(word)
		anagramsMap[normalized] = append(anagramsMap[normalized], word)
	}

	result := make([][]string, 0)
	for _, group := range anagramsMap {
		if len(group) > 1 {
			result = append(result, group)
		}
	}
	return result
}
