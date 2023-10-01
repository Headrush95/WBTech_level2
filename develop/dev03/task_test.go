package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	name        string
	input       []string
	expectedOut []string
}

func TestReverse(t *testing.T) {
	tests := []Test{
		{
			"Even Length",
			[]string{"a", "b", "c", "d"},
			[]string{"d", "c", "b", "a"},
		},
		{
			"Odd Length",
			[]string{"a", "b", "c", "d", "e"},
			[]string{"e", "d", "c", "b", "a"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reverse(test.input)
			assert.Equal(t, test.expectedOut, test.input)
		})
	}
}

func TestDeleteDuplicates(t *testing.T) {
	tests := []Test{
		{
			"Has Duplicates",
			[]string{"a", "a", "a", "b", "c", "d"},
			[]string{"a", "b", "c", "d"},
		},
		{
			"Has no Duplicates",
			[]string{"a", "b", "c", "d"},
			[]string{"a", "b", "c", "d"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedOut, deleteDuplicates(test.input))
		})
	}
}

func TestNumericSort(t *testing.T) {
	tests := []Test{
		{
			"Only Numbers",
			[]string{"3", "9", "1", "20", "0"},
			[]string{"0", "1", "3", "9", "20"},
		},
		{
			"With Words",
			[]string{"c", "20", "a", "1", "15", "b", "8"},
			[]string{"1", "8", "15", "20", "a", "b", "c"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			numericSort(test.input)
			assert.Equal(t, test.expectedOut, test.input)
		})
	}
}

func TestSortByColumn(t *testing.T) {
	tests := []Test{
		{
			"Only by column index",
			[]string{"sad a qwe", "ads", "q23 c 12qw afasd", "qr23r b"},
			[]string{"sad a qwe", "qr23r b", "q23 c 12qw afasd", "ads"},
		},
		{
			"By column index + Numeric",
			[]string{"fsadf fsad3", "qw 10 fqwe 3r", "ad 1 fas3", "fw 4 asdfa3", "sda", "2"},
			[]string{"ad 1 fas3", "fw 4 asdfa3", "qw 10 fqwe 3r", "fsadf fsad3", "sda", "2"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sortByColumn(test.input, 1, true)
			assert.Equal(t, test.expectedOut, test.input)
		})
	}
}
