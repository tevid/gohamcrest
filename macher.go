package test

type Matcher interface {
	//run match
	Match(actual interface{}) bool

	//get fail reason
	FailReason(actual interface{}) string

	//get negation fail reason
	NegationFailReason(actual interface{}) string
}