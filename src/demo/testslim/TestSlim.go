package testslim

import (
	"goslim/slimlist"
	"goslim/slimprotocol"
)

//static local variables
type TestSlim struct {
	noArgsCalled    int
	arg             string
	constructionArg string
	echoBuf         string
}

func TestSlim_Create(args *slimlist.SlimList) interface{} {
	if args.GetLength() > 1 {
		return nil
	}

	self := new(TestSlim)

	if args.GetLength() == 1 {
		self.constructionArg = args.GetStringAt(0)
	}

	return self
}

func (self *TestSlim) NoArgsCalled() int {
	return self.noArgsCalled
}

func (self *TestSlim) NoArgs(args *slimlist.SlimList) string {
	self.noArgsCalled = 1
	return "/__VOID__/"
}

func (self *TestSlim) ReturnValue(args *slimlist.SlimList) string {
	return "value"
}

func (self *TestSlim) OneArg(args *slimlist.SlimList) string {
	return args.GetStringAt(0)
}

func (self *TestSlim) Add(args *slimlist.SlimList) string {
	return args.GetStringAt(0) + args.GetStringAt(1)
}

func (self *TestSlim) Null(args *slimlist.SlimList) string {
	return ""
}

func (self *TestSlim) SetArg(args *slimlist.SlimList) string {
	self.arg = args.GetStringAt(0)
	return "/__VOID__/"
}

func (self *TestSlim) GetArg(args *slimlist.SlimList) string {
	return self.arg
}

func (self *TestSlim) GetArgFromFunctionWithUnderscores(args *slimlist.SlimList) string {
	return self.arg
}

func (self *TestSlim) GetConstructionArg(args *slimlist.SlimList) string {
	return self.constructionArg
}

func (self *TestSlim) ReturnError(args *slimlist.SlimList) string {
	return slimprotocol.EXCEPTION("my exception")
}
