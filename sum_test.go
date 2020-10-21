package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	testcases := []struct {
		input1 int
		intpu2 int
		expect int
	}{
		{1, 2, 3},
		{4, 5, 9},
	}

	for _, tc := range testcases {
		assert.Equal(t, sum(tc.input1, tc.intpu2), tc.expect)
	}
}
