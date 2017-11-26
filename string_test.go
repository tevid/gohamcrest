package tug

import (
	"testing"
)

func TestStartWith(t *testing.T) {
	Assert(t,"joe",StartWith("j"))
	//Assert(t,"joe",StartWith("0"))
	//Assert(t,"joe",Not(StartWith("j")))
	Assert(t,"joe",Not(StartWith("0")))
}

func TestEndWith(t *testing.T) {
	Assert(t,"joe",EndWith("e"))
	//Assert(t,"joe",EndWith("0"))
	//Assert(t,"joe",Not(EndWith("e")))
	Assert(t,"joe",Not(EndWith("0")))
}


