package SilentiumIO_test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	. "github.com/nnguyent/testcicd/SilentiumIO"

	"github.com/stretchr/testify/assert"
)

func captureOutput(f func(n int), n int) string {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f(n)
	w.Close()
	os.Stdout = stdout
	out, _ := ioutil.ReadAll(r)
	return string(out)
}
func TestPrintNumber(t *testing.T) {
	log.Println("TestPrintNumber running")

	testcases := []struct {
		Value  int
		Expect []string
	}{
		{3, []string{}},
		{4, []string{"Silentium"}},
		{7, []string{"IO"}},
		{7, []string{"Silentium", "IO"}},
		{28, []string{"Silentium", "IO", "SilentiumIO"}},
	}

	for _, tc := range testcases {
		output := captureOutput(PrintNumber, tc.Value)
		// fmt.Println("Expect:", strings.Join(tc.Expect, ","), "Output:", output)
		for _, exp := range tc.Expect {
			assert.Contains(t, output, exp)
		}
	}
}
