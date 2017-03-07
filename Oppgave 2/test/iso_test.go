package test

import (
	"testing"
	"../iso"
)

func TestGreetingExtendedASCII(t *testing.T) {
	testExtendedAscii := iso.GreetingExtendedASCII()
	if !(isExtendedASCII(testExtendedAscii)){
		t.Fail()
	}
}

func isExtendedASCII(s string) bool {
	for _, c := range s {
		if c < 127 && c > 255 {
			return false
		}
	}
	return true
}