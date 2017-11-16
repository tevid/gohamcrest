package tug

import (
	"testing"
)


func TestGreaterThan(t *testing.T) {
	Assert(t,3,GreaterThan(2))
}