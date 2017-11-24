package tug

import (
	"fmt"
)

type greaterThanMatcher struct {
	reason string
	expected interface{}
}

func (this *greaterThanMatcher) Match(actual interface{}) bool {
	validateParamType(actual,this.expected)

	switch actual.(type) { //多选语句switch
	case float64:
		return actual.(float64)>this.expected.(float64)
	case float32:
		return actual.(float32)>this.expected.(float32)
	case int:
		return actual.(int)>this.expected.(int)
	}

	return false
}

func (this *greaterThanMatcher)FailReason(actual interface{})string {
	return fmt.Sprintf(this.reason,actual,LOGIC_NOT,this.expected)
}

func (this *greaterThanMatcher)NegationFailReason(actual interface{})string {
	return fmt.Sprintf(this.reason,actual,EMPTY,this.expected)
}

func GreaterThan(expected interface{}) Matcher {
	return &greaterThanMatcher{
		expected:expected,
		reason:"%v is %s greater than %v(excepted)",
	}
}

func LessThanOrEquals(expected interface{}) Matcher {
	return Not(GreaterThan(expected))
}

type lessThanMatcher struct {
	reason string
	expected interface{}
}

func (this *lessThanMatcher) Match(actual interface{}) bool {
	validateParamType(actual,this.expected)

	switch actual.(type) { //多选语句switch
	case float64:
		return actual.(float64)<this.expected.(float64)
	case float32:
		return actual.(float32)<this.expected.(float32)
	case int:
		return actual.(int)<this.expected.(int)
	}

	return false
}

func (this *lessThanMatcher)FailReason(actual interface{})string {
	return fmt.Sprintf(this.reason,actual,LOGIC_NOT,this.expected)
}

func (this *lessThanMatcher)NegationFailReason(actual interface{})string {
	return fmt.Sprintf(this.reason,actual,EMPTY,this.expected)
}

func LessThan(expected interface{}) Matcher {
	return &lessThanMatcher{
		expected:expected,
		reason:"%v is %s less than %v(excepted)",
	}
}

func GreaterThanOrEquals(expected interface{}) Matcher {
	return Not(LessThan(expected))
}