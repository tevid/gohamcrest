package test

import (
	"testing"
	"path/filepath"
	"runtime"
	"container/list"
)


func NotEmptyList(t *testing.T,list *list.List ) {
	if list==nil||list.Len() == 0 {
		_, file, line, _ := runtime.Caller(1)
		t.Logf("\033[31m%s:%d:\n\n\t   not empty (expected)\n\n",
			filepath.Base(file), line)
		t.FailNow()
	}
}

func HasStringItems(t *testing.T,list *list.List, obj interface{}) {
	NotEmptyList(t,list)

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


