package test

import (
	"fmt"
)

type not struct {
	trueMatcher Matcher
}

func (this *not)Match(expected interface{}) bool {
	return !this.trueMatcher.Match(expected)
}

func (this *not)FailReason(expected interface{})string {
	return fmt.Sprintf("not %s ",this.trueMatcher.FailReason(expected))
}

func Not(matcher Matcher) Matcher {
	return &not{
		trueMatcher:matcher,
	}
}