package test

import "testing"

func TestEqual(t *testing.T) {
	Assert(t,2,Equal(2))
	Assert(t,"joe",Equal("joe"))
}

func TestNotEqual(t *testing.T) {
	Assert(t,2,NotEqual(3))
	Assert(t,"joe",NotEqual("joe1"))
}

func TestNilVal(t *testing.T) {
	Assert(t,nil,NilVal())
}

func TestNotNilVal(t *testing.T) {
	Assert(t,"",NotNilVal())
}