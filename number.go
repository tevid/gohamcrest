package gohamcrest

type greaterThanMatcher struct {
	BaseMatcher
}

func (this *greaterThanMatcher) Match(actual interface{}) bool {
	validateParamType(actual,this.Expected)

	switch actual.(type) { //多选语句switch
	case float64:
		return actual.(float64)>this.Expected.(float64)
	case float32:
		return actual.(float32)>this.Expected.(float32)
	case int:
		return actual.(int)>this.Expected.(int)
	}

	return false
}

func GreaterThan(expected interface{}) Matcher {
	matcher := &greaterThanMatcher{}
	matcher.Expected=expected
	matcher.Reason="%v is %s greater than %v(excepted)"
	return matcher
}

func LessThanOrEquals(expected interface{}) Matcher {
	return Not(GreaterThan(expected))
}

type lessThanMatcher struct {
	BaseMatcher
}

func (this *lessThanMatcher) Match(actual interface{}) bool {
	validateParamType(actual,this.Expected)

	switch actual.(type) { //多选语句switch
	case float64:
		return actual.(float64)<this.Expected.(float64)
	case float32:
		return actual.(float32)<this.Expected.(float32)
	case int:
		return actual.(int)<this.Expected.(int)
	}

	return false
}

func LessThan(expected interface{}) Matcher {
	matcher := &lessThanMatcher{}
	matcher.Expected=expected
	matcher.Reason="%v is %s less than %v(excepted)"
	return matcher
}

func GreaterThanOrEquals(expected interface{}) Matcher {
	return Not(LessThan(expected))
}