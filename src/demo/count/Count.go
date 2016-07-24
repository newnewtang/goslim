package count

import (
	. "goslim/slimlist"
	"strconv"
)

type Count struct {
	count int
}

func Count_Create(*SlimList) interface{} {
	return new(Count)
}

func (self *Count) Count(args *SlimList) string {
	self.count++
	return ""
}

func (self *Count) Counter(args *SlimList) string {
	return strconv.Itoa(self.count)
}
