package exceptions

import ()

import (
	. "goslim/slimlist"
	"goslim/slimprotocol"
)

type Exceptions struct {
	m1     float32
	m2     float32
	result string
}

func Exceptions_Create(*SlimList) interface{} {
	return new(Exceptions)
}

func (self *Exceptions) SetTrouble(args *SlimList) string {
	return slimprotocol.EXCEPTION("You stink")
}
