package gomatcher

import (
	"reflect"
	"fmt"
)

const (
	EMPTY=""
	LOGIC_NOT=" not "
)

//validate parameters type
func validateParamType(actual interface{},expected interface{})  {
	actualType:=reflect.TypeOf(actual)
	expectedType:=reflect.TypeOf(expected)
	if actualType.Name()!=expectedType.Name(){
		panic(fmt.Sprintf("type not match actual type:%T,expected type:%T",actual,expected))
	}
}