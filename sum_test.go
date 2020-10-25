package main_test

// package main

import (
	"log"
	"os"
	"testing"

	. "github.com/nnguyent/testcicd"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	log.Println("TestMain running")
	os.Exit(m.Run())
}

// func TestRunMain(t *testing.T) {
// 	log.Println("TestRunMain running")
// 	main()
// }

func TestSum(t *testing.T) {
	log.Println("TestSum running")

	testcases := []struct {
		input1 int
		intpu2 int
		expect int
	}{
		{1, 2, 3},
		{4, 5, 9},
	}

	for _, tc := range testcases {
		assert.Equal(t, Sum(tc.input1, tc.intpu2), tc.expect)
	}
}
