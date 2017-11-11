package test

import (
	"testing"
	"runtime"
	"path/filepath"
)

func AssertThat(t *testing.T, expected interface{}, matcher Matcher) {
	if !matcher.Match(expected) {
		_, file, line, _ := runtime.Caller(1)
		t.Logf("\033[31m%s:%d:\n\n\t   %#v (expected)\n\n\t but : \n\n\t   %#v \033[39m\n\n",
			filepath.Base(file), line, expected, matcher.FailReason(expected))
		t.FailNow()
	}
}
