package main

import (
	"testing"
)

func TestCleanupInput(t *testing.T) {
	cases := []struct{
		input string 
		expected []string
		}{
			{
				input: "Kanto    -FIRE\t\t--fighting",
				expected: []string{
					"kanto", 
					"-fire", 
					"--fighting", 
				},
			},
			{
				input: "Foo\tBazz SHIZZ         dang",
				expected: []string{
					"foo", 
					"bazz", 
					"shizz", 
					"dang",
				},
			},
		}
	

	for i, cs := range(cases) {
		actual := cleanupInput(cs.input)

		if len(actual) != len(cs.expected) {
			t.Errorf(
				"\nError: different length for %dth case: %d (actual), %d (expected)", 
				i + 1, 
				len(actual), 
				len(cs.expected),
			)
			continue
		} 
		for i, word := range(cs.expected) {
			if actual[i] != word {
				t.Errorf(
					"\nError: different word for %dth case: %s (actual), %s (expected)", 
					i + 1, 
					actual[i], 
					word,
				)
			}
		}
	}
}