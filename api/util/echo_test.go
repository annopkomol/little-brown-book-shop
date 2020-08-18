package util

import (
	a "github.com/stretchr/testify/assert"
	"testing"
)

func TestGetServerAddress(t *testing.T) {
	testCases := []struct {
		name       string
		input      string
		expectAddr string
		expectErr  bool
	}{{
		name:       "valid input",
		input:      "7777",
		expectAddr: ":7777",
		expectErr:  false,
	}, {
		name:       "invalid input",
		input:      "a123bc",
		expectAddr: "",
		expectErr:  true,
	}, {
		name:       "port out of range",
		input:      "70000",
		expectAddr: "",
		expectErr:  true,
	}, {
		name:       "port out of range",
		input:      "-1",
		expectAddr: "",
		expectErr:  true,
	}}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualAddr, actualErr := GetServerAddress(tc.input)
			a.Equal(t, tc.expectAddr, actualAddr)
			if tc.expectErr {
				a.NotNil(t, actualErr)
			} else {
				a.Nil(t, actualErr)
			}
		})
	}
}
