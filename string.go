package tug

import (
	"strings"
)

func toString(o interface{}) string  {
	return o.(string)
}

type startWith struct {
	BaseMatcher
}

func (this *startWith) Match(actual interface{}) bool {
	excepted := toString(this.Expected)
	if str, ok := actual.(string); ok {
		return strings.HasPrefix(str, excepted)
	}
	return false
}

func StartWith(expected string) Matcher {
	matcher := &startWith{}
	matcher.Expected=expected
	matcher.Reason="%v %s start with %v"
	return matcher
}

type endWith struct {
	BaseMatcher
}

func (this *endWith) Match(actual interface{}) bool {
	excepted := toString(this.Expected)
	if str, ok := actual.(string); ok {
		return strings.HasSuffix(str, excepted)
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
	excepted := toString(this.Expected)
	if str, ok := actual.(string); ok {
		return strings.Contains(str, excepted)
	}
	return false
}

func ContainString(expected string) Matcher {
	matcher := &containString{}
	matcher.Expected=expected
	matcher.Reason="%v %s contain %v"
	return matcher
}


