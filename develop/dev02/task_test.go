package dev02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	name          string
	input         string
	expectedOut   string
	expectedError error
}

func TestUnpackString(t *testing.T) {
	tests := []Test{
		{
			name:          "OK1",
			input:         "a4bc2d5e",
			expectedOut:   "aaaabccddddde",
			expectedError: nil,
		},
		{
			name:          "OK2",
			input:         "abcd",
			expectedOut:   "abcd",
			expectedError: nil,
		},
		{
			name:          "Error. First Digit",
			input:         "45",
			expectedOut:   "",
			expectedError: invalidStringError,
		},
		{
			name:          "Error. Empty string",
			input:         "",
			expectedOut:   "",
			expectedError: invalidStringError,
		},
		{
			name:          "OK. Escape 1",
			input:         "qwe\\4\\5",
			expectedOut:   "qwe45",
			expectedError: nil,
		},
		{
			name:          "OK. Escape 2",
			input:         "qwe\\45",
			expectedOut:   "qwe44444",
			expectedError: nil,
		},
		{
			name:          "OK. Escape 3",
			input:         "qwe\\\\5",
			expectedOut:   "qwe\\\\\\\\\\",
			expectedError: nil,
		},
		{
			name:          "Error. Two numbers",
			input:         "qwe45",
			expectedOut:   "",
			expectedError: invalidStringError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out, err := UnpackString(test.input)
			assert.Equal(t, test.expectedOut, out)
			assert.Equal(t, test.expectedError, err)
		})
	}
}
