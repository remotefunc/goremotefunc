package remotefunc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCallfuncString(t *testing.T) {
	params := `["test"]`
	rf := RemoteFunc{}
	fun := func(s string) string {
		return s
	}
	result := rf.callfunc(params, fun)
	assert.Equal(t, `"test"`, result)
}

func TestCallfuncStruct(t *testing.T) {
	type Test struct {
		A string
		B int
	}

	rf := RemoteFunc{}

	fun := func(t Test) Test {
		return t
	}

	params := `[{"A":"test","B":1}]`

	result := rf.callfunc(params, fun)

	assert.Equal(t, `{"A":"test","B":1}`, result)
}

func TestCallfuncMultipleValues(t *testing.T) {
	type Test struct {
		A string
		B int
	}

	rf := RemoteFunc{}

	fun := func(t Test, s string, i int) Test {
		t.A = s
		t.B = i
		return t
	}

	params := `[{"A":"test","B":1}, "testing", 3]`

	result := rf.callfunc(params, fun)

	assert.Equal(t, `{"A":"testing","B":3}`, result)

}

func TestCallfunNoValues(t *testing.T) {
	fun := func() {}

	rf := RemoteFunc{}

	result := rf.callfunc("[]", fun)
	assert.Empty(t, "", result)
}

func TestCallfunNoValuesNoList(t *testing.T) {
	fun := func() {}

	rf := RemoteFunc{}

	result := rf.callfunc("", fun)
	assert.Empty(t, "", result)
}
