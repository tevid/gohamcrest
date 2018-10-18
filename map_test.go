package gohamcrest

import (
	"testing"
)

func TestHasMapValue(t *testing.T) {
	m := make(map[string]string)
	m["a"]="a"
	m["joe"]="joe"
	Assert(t,m,HasMapValue("joe"))
}
