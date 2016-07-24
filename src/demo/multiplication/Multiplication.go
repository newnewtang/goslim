package multiplication

import (
	. "goslim/slimlist"
	"strconv"
)

type Multiplication struct {
	m1     float32
	m2     float32
	result string
}

func Multiplication_Create(*SlimList) interface{} {
	return new(Multiplication)
}

func (self *Multiplication) SetMultiplicand1(args *SlimList) string {
	self.m1 = args.GetFloatAt(0)
	return ""
}

func (self *Multiplication) SetMultiplicand2(args *SlimList) string {
	self.m2 = args.GetFloatAt(0)
	return ""
}

func (self *Multiplication) Product(args *SlimList) string {
	return strconv.FormatFloat(float64(self.m1*self.m2), 'f', -1, 32)
}
