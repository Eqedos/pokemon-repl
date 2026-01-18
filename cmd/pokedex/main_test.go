package main

import (
	"testing"

	"github.com/eqedos/repl/internal/pokeapi"
)

func TestCleanInput(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "trims whitespace and splits words",
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "handles multiple spaces between words",
			input:    "damn    that's crazzyy",
			expected: []string{"damn", "that's", "crazzyy"},
		},
		{
			name:     "converts to lowercase",
			input:    "HELLO World",
			expected: []string{"hello", "world"},
		},
		{
			name:     "handles empty input",
			input:    "",
			expected: []string{},
		},
		{
			name:     "handles only whitespace",
			input:    "   ",
			expected: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := cleanInput(tc.input)

			if len(actual) != len(tc.expected) {
				t.Errorf("length mismatch: expected %d, got %d (expected: %v, actual: %v)",
					len(tc.expected), len(actual), tc.expected, actual)
				return
			}

			for i, expectedWord := range tc.expected {
				if actual[i] != expectedWord {
					t.Errorf("word mismatch at index %d: expected %q, got %q",
						i, expectedWord, actual[i])
				}
			}
		})
	}
}

func TestMapCaching(t *testing.T) {
	// Initialize client and config
	client := pokeapi.NewClient()
	firstURL := client.GetFirstLocationAreasURL()

	cfg := &config{
		client:  client,
		nextURL: &firstURL,
		prevURL: nil,
	}

	// First map call - fetches from API and caches page 1
	err := commandMap(cfg, nil)
	if err != nil {
		t.Fatalf("first map failed: %v", err)
	}

	// Second map call - fetches from API and caches page 2
	err = commandMap(cfg, nil)
	if err != nil {
		t.Fatalf("second map failed: %v", err)
	}

	// mapb call - should retrieve page 1 from cache
	err = commandMapb(cfg, nil)
	if err != nil {
		t.Fatalf("mapb failed: %v", err)
	}

	// Call map again - should use cached page 2
	// (will print "using cached data" if working correctly)
	err = commandMap(cfg, nil)
	if err != nil {
		t.Fatalf("third map failed: %v", err)
	}
}
