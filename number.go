package tug

import (
	"fmt"
	"reflect"
)

type greaterThanMatcher struct {
	reason string
	expected interface{}
}

func (this *greaterThanMatcher) Match(actual interface{}) bool {
	actualType:=reflect.TypeOf(actual)
	expectedType:=reflect.TypeOf(this.expected)
	if actualType.Name()!=expectedType.Name(){
		panic(fmt.Sprintf("type not match actual type:%T,expected type:%T",actual,this.expected))
	}

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