package test

import (
	"testing"
	"container/list"
)

func TestNotEmptyList(t *testing.T) {
	l := list.New()
	l.PushBack("joe")
	NotEmptyList(t,l)
}

func TestHasStringItems(t *testing.T) {
	l := list.New()
	l.PushBack("joe")
	HasStringItems(t,l,"joe")
}


