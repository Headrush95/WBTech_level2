package dev04

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	name        string
	input       []string
	expectedOut map[string][]string
}

func TestFindAnagram(t *testing.T) {
	tests := []Test{
		{
			"OK",
			[]string{"тяпка", "пятак", "листок", "пятка", "слитОк", "тЯпкА", "столик"},
			map[string][]string{
				"листок": {"листок", "слиток", "столик"},
				"тяпка":  {"пятак", "пятка", "тяпка"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedOut, findAnagram(test.input))
		})
	}
}
