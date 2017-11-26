package tug

import (
	"container/list"
	"fmt"
)

type emptyList struct {
	reason string
}

func (this *emptyList) Match(actual interface{}) bool {
	list, ok := actual.(*list.List)
	if ok {
		return list==nil||list.Len()== 0
	}else{
		panic(fmt.Sprintf("type not match List.list,%T(actual)",list))
	}
	return false
}

func (this *emptyList)FailReason(actual interface{})string {
	return fmt.Sprintf(this.reason,actual,LOGIC_NOT)
}

func (this *emptyList)NegationFailReason(actual interface{})string {
	return fmt.Sprintf(this.reason,actual)
}

func EmptyList() Matcher {
	return &emptyList{
		reason:"list is %s nil, %v(actual)",
	}
}

type hasItems struct {
	BaseMatcher
}

func (this *hasItems) Match(actual interface{}) bool {
	list, ok := actual.(*list.List)
	if !ok {
		panic(fmt.Sprintf("type not match List.list,%T(actual)",list))
	}

	flag := false

	for e := list.Front(); e != nil; e = e.Next() {
		if e.Value == this.Expected {
			flag = true
			break
		}
	}

	return flag
}

func (this *hasItems)FailReason(actual interface{})string {
	return fmt.Sprintf(this.Reason,EMPTY,this.Expected)
}

func (this *hasItems)NegationFailReason(actual interface{})string {
	return fmt.Sprintf(this.Reason,LOGIC_NOT,this.Expected)
}

func HasItems(expected interface{}) Matcher {
	matcher := &hasItems{}
	matcher.Expected=expected
	matcher.Reason="list is %s contain %v(excepted)"
	return matcher
}
