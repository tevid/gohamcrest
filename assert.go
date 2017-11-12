package test

import (
	"testing"
	"runtime"
	"path/filepath"
)

func Assert(t *testing.T, actual interface{}, matcher Matcher) {
	if !matcher.Match(actual) {
		_, file, line, _ := runtime.Caller(1)
		t.Logf("\033[31m%s:%d:\n\n\t   %#v (expected)\n\n\t but : \n\n\t   %#v \033[39m\n\n",
			filepath.Base(file), line, matcher, matcher.FailReason(actual))
		t.FailNow()
	}
}
