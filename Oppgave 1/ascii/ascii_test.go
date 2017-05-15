package ascii

import (
	"testing"
)

func TestGreetingASCII(t *testing.T) {
	tempAscii := GreetingASCII()
	if !(isASCII(tempAscii)){
		t.Fail()
	}
}

//Hentet fra https://play.golang.org/p/hnZzfnbXeF
func isASCII(s string) bool {
	for _, c := range s {
		if c > 127 {
			return false
		}
	}
	return true
}