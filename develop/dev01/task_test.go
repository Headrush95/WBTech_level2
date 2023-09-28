package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	invalidHost = "0.beevik-ntp.ntp.og"
	correctHost = "0.beevik-ntp.pool.ntp.org"
)

type Test struct {
	name           string
	inputHost      string
	expectedOutput time.Time
	expectedError  string
}

func TestGetTime(t *testing.T) {
	tests := []Test{
		{
			name:           "OK",
			inputHost:      correctHost,
			expectedError:  "",
			expectedOutput: time.Now(),
		},
		{
			name:           "invalid host",
			inputHost:      invalidHost,
			expectedOutput: time.Time{},
			expectedError:  fmt.Sprintf("lookup %s: no such host", invalidHost),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expTime, expErr := GetTime(test.inputHost)
			if test.expectedError != "" {
				assert.Equal(t, test.expectedError, expErr.Error())
				return
			}
			assert.Condition(t, func() bool {
				fmt.Println(expTime.Sub(test.expectedOutput))
				return expTime.Sub(test.expectedOutput) < 2*time.Second && expTime.Sub(test.expectedOutput) > -2*time.Second
			},
			)
		})
	}

}
