package test

import (
	"testing"
	"path/filepath"
	"runtime"
	"container/list"
	"fmt"
)

type emptyList struct {
	reason string
}

func (this *emptyList) Match(actual interface{}) bool {
	if list, ok := actual.(*list.List); ok {
		return list==nil||list.Len()== 0
	}
	return false
}

func (this *emptyList)FailReason(actual interface{})string {
	return fmt.Sprintf(this.reason,actual,LOGIC_NOT)
}

func (this *emptyList)NegationFailReason(actual interface{})string {
	return fmt.Sprintf(this.reason,actual)
}

func EmptyList() Matcher {
	return &emptyList{
		reason:"list is %s nil, %v(actual)",
	}
}

func HasStringItems(t *testing.T,list *list.List, obj interface{}) {
	//NotEmptyList(t,list)

	flag := false

	for e := list.Front(); e != nil; e = e.Next() {
		if e.Value == obj {
			flag = true
			break
		}
	}

	if !flag {
		_, file, line, _ := runtime.Caller(1)
		t.Logf("\033[31m%s:%d:\n\n\t   contain %#v (expected)\n\n\t not contain %#v (actual)\033[39m\n\n",
			filepath.Base(file), line, obj, obj)
		t.FailNow()
	}
}


