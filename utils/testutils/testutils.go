package testutils

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

/*
Testing Functions code by https://github.com/benbjohnson/testing
*/

// Assert fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// Ok fails the test if an err is not nil.
func Ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// Equals fails the test if exp is not equal to act.
func Equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

/*
CreateDummyDirs creates directories under a list of test directories
*/
func CreateDummyDirs(testDirs, directories []string) (err error) {
	for _, testingDir := range testDirs {
		for _, dummyName := range directories {
			path := filepath.Join(testingDir, dummyName)
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

/*
ClearDummyDirs clears all the directories sent in as string
*/
func ClearDummyDirs(testingDirs []string) (err error) {
	for _, testingDir := range testingDirs {
		err = os.RemoveAll(testingDir)
		if err != nil {
			return err
		}
	}
	return nil
}
