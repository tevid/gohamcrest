package test

type Matcher interface {

	//run match
	Match(expected interface{}) bool

	//get fail reason
	FailReason(expected interface{})string
}