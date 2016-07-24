package main

import (
	"demo/count"
	"demo/division"
	"demo/employee"
	"demo/exceptions"
	"demo/multiplication"
	"demo/testslim"
	"goslim"
)

func main() {
	goslim.RegisterFixture("Division", division.Division_Create)
	goslim.RegisterFixture("Count", count.Count_Create)
	goslim.RegisterFixture("EmployeePayRecordsRow", employee.EmployeePayRecordsRow_Create)
	goslim.RegisterFixture("Multiplication", multiplication.Multiplication_Create)
	goslim.RegisterFixture("ExceptionsExample", exceptions.Exceptions_Create)
	goslim.RegisterFixture("TestSlim", testslim.TestSlim_Create)
	goslim.RegisterMethod("TestSlim", "echo", "OneArg")
	goslim.WAIT_AND_RUN()
}
