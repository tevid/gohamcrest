package test

import (
	"testing"
	"container/list"
)

func TestEmptyList(t *testing.T) {
	l := list.New()
	Assert(t,l,EmptyList())
}

func TestNotEmptyList(t *testing.T) {
	l := list.New()
	l.PushBack("joe")
	Assert(t,l,Not(EmptyList()))
}

func TestHasStringItems(t *testing.T) {
	l := list.New()
	l.PushBack("joe")
	HasStringItems(t,l,"joe")
}


