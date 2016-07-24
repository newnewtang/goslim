package employee

import (
	. "goslim/slimlist"
)

type EmployeePayRecordsRow struct {
}

func EmployeePayRecordsRow_Create(*SlimList) interface{} {
	return new(EmployeePayRecordsRow)
}

func (self *EmployeePayRecordsRow) Query(args *SlimList) string {
	id := SlimList_Create()
	pay := SlimList_Create()
	record1 := SlimList_Create()
	records := SlimList_Create()

	id.AddString("id")
	id.AddString("1")

	pay.AddString("pay")
	pay.AddString("1000")

	record1.AddList(id)
	record1.AddList(pay)

	records.AddList(record1)

	return records.Serialize()
}
