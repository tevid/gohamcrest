package tug

import (
	"strings"
)

type startWith struct {
	BaseMatcher
}

func (this *startWith) Match(actual interface{}) bool {
	if str, ok := actual.(string); ok {
		return strings.HasPrefix(str, this.Expected)
	}
	return false
}

func StartWith(expected string) Matcher {
	matcher := &endWith{}
	matcher.Expected=expected
	matcher.Reason="%v %s start with %v"
	return matcher
}

type endWith struct {
	BaseMatcher
}

func (this *endWith) Match(actual interface{}) bool {
	if str, ok := actual.(string); ok {
		return strings.HasSuffix(str, this.Expected)
	}
	return false
}

func EndWith(expected string) Matcher {
	matcher := &endWith{}
	matcher.Expected=expected
	matcher.Reason="%v %s end with %v"
	return matcher
}

type containString struct {
	BaseMatcher
}

func (this *containString) Match(actual interface{}) bool {
	if str, ok := actual.(string); ok {
		return strings.Contains(str, this.Expected)
	}
	return false
}

func ContainString(expected string) Matcher {
	matcher := &containString{}
	matcher.Expected=expected
	matcher.Reason="%v %s contain %v"
	return matcher
}


