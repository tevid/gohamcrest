package tug

import (
	"testing"
)


func TestGreaterThan(t *testing.T) {
	Assert(t,3.2,GreaterThan(2))
}