package test

import (
	"reflect"
	"fmt"
)

type isEqual struct {
	expected interface{}
}

func (this *isEqual)Match(actual interface{}) bool {
	return reflect.DeepEqual(this.expected, actual)
}

func (this *isEqual)FailReason(actual interface{})string {
	return fmt.Sprintf("%v(expected)!=%v(actual)",this.expected,actual)
}

func (this *isEqual)NegationFailReason(actual interface{})string {
	return fmt.Sprintf("%v(expected)=%v(actual)",this.expected,actual)
}

//Create a Matcher for match the actual object is equal excepted object
//example:
//int:tug.Assert(t,2,Equal(2))
//string:tug.Assert(t,"joe",Equal("joe"))
func Equal(expected interface{}) Matcher {
	return &isEqual{expected:expected}
}


func NotEqual(expected interface{}) Matcher{
	return Not(Equal(expected))
}


func NotNilVal() Matcher{
	return Not(NilVal())
}

type isNil struct {
	expected interface{}
}

func (this *isNil)Match(actual interface{}) bool {
	if actual == nil {
		return true
	}

	value := reflect.ValueOf(actual)
	kind := value.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
		return true
	}

	return false
}

func (this *isNil)FailReason(actual interface{})string {
	return fmt.Sprintf("%v(expected)!=%v(actual)",this.expected,actual)
}

func (this *isNil)NegationFailReason(actual interface{})string {
	return fmt.Sprintf("%v(expected)=%v(actual)",this.expected,actual)
}

//Create a Matcher for match the object is nil or not.
//example
func NilVal() Matcher {
	return &isNil{}
}