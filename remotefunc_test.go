package remotefunc

import (
	"testing"
)

func TestCallfunc(t *testing.T) {
	params := "[\"test\"]"
	rf := RemoteFunc{}
	fun := func(s string) string {
		return s
	}
	result := rf.callfunc(params, fun)
	if result != "\"test\"" {
		t.Errorf("T1: %s != T2: %s", result, "\"test\"")
	}
}
