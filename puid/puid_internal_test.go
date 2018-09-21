package puid

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

var testSplit = []struct {
	input     int64
	expected1 uint32
	expected2 uint32
}{
	{0x0FFFFFFF0FABCDEA, 0x0FFFFFFF, 0x0FABCDEA},
	{0x0000000010F00FFF, 0x00000000, 0x10F00FFF},
	{0x0F00FFFF00000000, 0x0F00FFFF, 0x00000000},
	{0x20F0F0F010F0F0F0, 0x20F0F0F0, 0x10F0F0F0},
	{0x11FBBFFA4FFEEF11, 0x11FBBFFA, 0x4FFEEF11},
	{0x3A019DE1632300FA, 0x3A019DE1, 0x632300FA},
}

func TestSplitInt64Time(t *testing.T) {
	for _, test := range testSplit {
		lowSide, highSide := splitInt64Time(test.input)
		switch {
		case lowSide != test.expected1:
			t.Errorf("Low side is incorrect")
		case highSide != test.expected2:
			t.Errorf("High side is incorrect")
		}
	}
}

/*
This controlled test generates a PUID based off the first 12 "predictable" bytes the last 4 bytes are random and therefore no need to be tested
*/
var testUUID = []struct {
	iUnixTime    int64
	iUnixnsTtime int64
	iProcessID   uint32
	num          [12]byte
}{
	{0x00000000FFFFFFFF, 0x00000000FFFFFFFF, 0x0000AABB, [12]byte{0xFF, 0xFF, 0xFF, 0xFF, 0x4A, 0x52, 0xAA, 0xBB, 0xFF, 0xFF, 0xFF, 0xFF}},
}

// Equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

func TestGeneratePUID(t *testing.T) {
	//puid := generatePUID(time.Now().Unix(), time.Now().UnixNano(), uint32(os.Getpid()))
	for _, test := range testUUID {
		puid := generatePUID(test.iUnixTime, test.iUnixnsTtime, test.iProcessID)
		for i, expected := range test.num {
			equals(t, puid.number[i], expected)
		}
	}
}
