package gohamcrest

import (
	"reflect"
)

type isEqual struct {
	BaseMatcher
}

func (this *isEqual)Match(actual interface{}) bool {
	return reflect.DeepEqual(this.Expected, actual)
}

//Create a Matcher for match the actual object is equal excepted object
//example:
//int:gohamcrest.Assert(t,2,Equal(2))
//string:gohamcrest.Assert(t,"joe",Equal("joe"))
func Equal(expected interface{}) Matcher {
	matcher := &isEqual{}
	matcher.Expected=expected
	matcher.Reason="%v %s equal %v"
	return matcher
}


func NotEqual(expected interface{}) Matcher{
	return Not(Equal(expected))
}


func NotNilVal() Matcher{
	return Not(NilVal())
}

type isNil struct {
	BaseMatcher
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

//Create a Matcher for match the object is nil or not.
//example
func NilVal() Matcher {
	matcher := &isNil{}
	matcher.Reason="%v %s equal %v"
	return matcher
}