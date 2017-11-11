package test

import "testing"

func TestEqual(t *testing.T) {
	AssertThat(t,2,Equal(2))
	AssertThat(t,"joe",Equal("joe"))
}