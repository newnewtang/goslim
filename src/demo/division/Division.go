package division

import (
	. "goslim/slimlist"
	"goslim/slimprotocol"
	"strconv"
)

type Division struct {
	numerator   float32
	denominator float32
	result      string
}

func Division_Create(*SlimList) interface{} {
	return new(Division)
}

func (self *Division) SetNumerator(args *SlimList) string {
	self.numerator = args.GetFloatAt(0)
	return ""
}

func (self *Division) SetDenominator(args *SlimList) string {
	self.denominator = args.GetFloatAt(0)
	if self.denominator == 0.0 {
		return slimprotocol.EXCEPTION("You shouldn't divide by zero now should ya?")
	}
	return ""
}

func (self *Division) Quotient(args *SlimList) string {
	quotient := self.numerator / self.denominator
	return strconv.FormatFloat(float64(quotient), 'f', -1, 32)
}

//These are optional.  If they aren't declared, they are ignored
func (self *Division) Execute(args *SlimList) string {
	return ""
}

func (self *Division) Reset(args *SlimList) string {
	self.denominator = 0.0
	self.numerator = 0.0
	return ""
}
func (self *Division) Table(args *SlimList) string {
	return ""
}

func (self *Division) String(args *SlimList) string {
	return "Division"
}
