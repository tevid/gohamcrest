package tug

import (
	"fmt"
	"strings"
)

type startWith struct {
	expected string
	reason string
}

func (this *startWith) Match(actual interface{}) bool {
	if str, ok := actual.(string); ok {
		return strings.HasPrefix(str, this.expected)
	}
	return false
}

func (this *startWith)FailReason(actual interface{})string {
	return fmt.Sprintf(this.reason,actual,LOGIC_NOT,this.expected)
}

func (this *startWith)NegationFailReason(actual interface{})string {
	return fmt.Sprintf(this.reason,actual,this.expected)
}

func StartWith(expected string) Matcher {
	return &startWith{
		expected:expected,
		reason:"%v %s start with %v",
		}
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


