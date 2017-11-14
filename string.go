package test

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


