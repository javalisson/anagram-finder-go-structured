package anagramfinder

import (
	"reflect"
	"sort"
	"testing"
)

// helper function to sort slice of slices for consistent comparison
func sortSliceOfSlices(sos [][]string) {
	for _, s := range sos {
		sort.Strings(s)
	}
	sort.Slice(sos, func(i, j int) bool {
		return sos[i][0] < sos[j][0]
	})
}

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected [][]string
	}{
		{
			name:  "standard test",
			words: []string{"created", "catered", "reacted", "decimal", "claimed", "hello"},
			expected: [][]string{
				{"catered", "created", "reacted"},
				{"claimed", "decimal"},
			},
		},
		{
			name:     "no anagrams",
			words:    []string{"unique", "words", "only"},
			expected: [][]string{},
		},
		{
			name:  "with duplicates",
			words: []string{"listen", "silent", "enlist", "listen"},
			expected: [][]string{
				{"enlist", "listen", "listen", "silent"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindAnagrams(tt.words)
			sortSliceOfSlices(got) // Sorting for consistent comparison
			sortSliceOfSlices(tt.expected)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("TestFindAnagrams failed for %v: expected %v, got %v", tt.name, tt.expected, got)
			}
		})
	}
}
