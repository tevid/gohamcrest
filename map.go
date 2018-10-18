package gohamcrest

import (
	"reflect"
)

type hasMapValue struct {
	BaseMatcher
}

func (this *hasMapValue) Match(actual interface{}) bool {
	listValue := reflect.ValueOf(actual)

	flag := false
	if reflect.TypeOf(actual).Kind() == reflect.Map {
		mapKeys := listValue.MapKeys()
		for i := 0; i < len(mapKeys); i++ {
			if mapKeys[i].Interface() == this.Expected {
				flag = true
				break
			}
		}
	}

	return flag
}

func HasMapValue(expected interface{}) Matcher {
	matcher := &hasMapValue{}
	matcher.Expected=expected
	matcher.Reason="map: %v is %s contain %v(excepted)"
	return matcher
}
