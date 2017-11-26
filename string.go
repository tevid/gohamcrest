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
	expected string
	reason string
}

func (this *endWith) Match(actual interface{}) bool {
	if str, ok := actual.(string); ok {
		return strings.HasSuffix(str, this.expected)
	}
	return false
}

func (this *endWith)FailReason(actual interface{})string {
	return fmt.Sprintf(this.reason,actual,LOGIC_NOT,this.expected)
}

func (this *endWith)NegationFailReason(actual interface{})string {
	return fmt.Sprintf(this.reason,actual,this.expected)
}

func EndWith(expected string) Matcher {
	return &endWith{
		expected:expected,
		reason:"%v %s end with %v",
	}
}

type containString struct {
	expected string
	reason string
}

func (this *containString) Match(actual interface{}) bool {
	if str, ok := actual.(string); ok {
		return strings.Contains(str, this.expected)
	}
	return false
}

func (this *containString)FailReason(actual interface{})string {
	return fmt.Sprintf(this.reason,actual,LOGIC_NOT,this.expected)
}

func (this *containString)NegationFailReason(actual interface{})string {
	return fmt.Sprintf(this.reason,actual,this.expected)
}

func ContainString(expected string) Matcher {
	return &endWith{
		expected:expected,
		reason:"%v %s contain %v",
	}
}


