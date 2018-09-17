package puid_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/Jose-R-Rodriguez/PUID/puid"
)

func TestPUIDString(t *testing.T) {
	puid := puid.NewPUID().String()
	matched, err := regexp.MatchString("(?:[A-Z2-7]{8})*(?:[A-Z2-7]{2}={6}|[A-Z2-7]{4}={4}|[A-Z2-7]{5}={3}|[A-Z2-7]{7}=)?", puid)
	if err != nil {
		t.Errorf("ERROR REG EXP MALFORMED")
	}
	if !matched {
		t.Errorf("Malformed base32 string")
	}
	fmt.Println(puid)
}
