package gohamcrest

import (
	"testing"
)


func TestGreaterThan(t *testing.T) {
	Assert(t,3,GreaterThan(2))
}

func TestLessThanOrEquals(t *testing.T) {
	Assert(t,3,LessThanOrEquals(3))
	Assert(t,2,LessThanOrEquals(3))
}

func TestLessThan(t *testing.T) {
	Assert(t,2,LessThan(3))
}

func TestGreaterThanOrEquals(t *testing.T) {
	Assert(t,3,GreaterThanOrEquals(3))
	Assert(t,3,GreaterThanOrEquals(2))
}