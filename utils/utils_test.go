package utils_test

import (
	"errors"
	"os"
	"os/exec"
	"testing"

	"github.com/Jose-R-Rodriguez/WDIO_CommandLineUtility/utils"
)

// Code and idea by Andrew Gerrand https://talks.golang.org/2014/testing.slide#23
func TestLogger(t *testing.T) {
	if os.Getenv("TEST_ERRORS") == "1" {
		utils.Log(errors.New("This is failing"), "")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestLogger")
	cmd.Env = append(os.Environ(), "TEST_ERRORS=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

var TestSlices = []struct {
	slice1         []string
	slice2         []string
	expectedResult bool
}{
	{
		[]string{"testing"},
		[]string{"testing"},
		true,
	},
	{
		[]string{"testing", "wat"},
		[]string{"testing", "fail"},
		false,
	},
	{
		[]string{"testing", "testing2", "unit"},
		[]string{"testing", "testing2", "unit"},
		true,
	},
}

func TestCompareStringSliceEquality(t *testing.T) {
	for _, pair := range TestSlices {
		result := utils.CompareStringSliceEquality(pair.slice1, pair.slice2)
		if result != pair.expectedResult {
			t.Errorf("Error pair comparison incorrect %v -- %v", pair.slice1, pair.slice2)
		}
	}
}
