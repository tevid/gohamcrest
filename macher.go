package tug

import (
	"fmt"
)

type Matcher interface {
	//run match
	Match(actual interface{}) bool

	//get fail reason
	FailReason(actual interface{}) string

	//get negation fail reason
	NegationFailReason(actual interface{}) string
}


type BaseMatcher struct {
	Expected interface{}
	Reason string
}
func (this *BaseMatcher) Match(actual interface{}) bool {
	return true
}

func (this *BaseMatcher)FailReason(actual interface{})string {
	return fmt.Sprintf(this.Reason,actual,LOGIC_NOT,this.Expected)
}

func (this *BaseMatcher)NegationFailReason(actual interface{})string {
	return fmt.Sprintf(this.Reason,actual,this.Expected)
}