package slimexecutor

import (
	"fmt"
	. "goslim/slimlist"
)

type slimExecutor struct {
	executor *StatementExecutor
}

var slimExecutorSingleton *slimExecutor

func Instance() *slimExecutor {
	if slimExecutorSingleton == nil {
		slimExecutorSingleton = new(slimExecutor)
		slimExecutorSingleton.executor = StatementExecutor_Create()
	}
	return slimExecutorSingleton
}

func AddResult(list *SlimList, id string, result string) {
	pair := SlimList_Create()
	pair.AddString(id)
	pair.AddString(result)
	list.AddList(pair)
}

func InvalidCommand(instruction *SlimList) string {
	id := instruction.GetStringAt(0)
	command := instruction.GetStringAt(1)
	return fmt.Sprintf("__EXCEPTION__:message:<<INVALID_STATEMENT: [\"%s\", \"%s\"].>>", id, command)
}

func MalformedInstruction(instruction *SlimList) string {
	return fmt.Sprintf("__EXCEPTION__:message:<<MALFORMED_INSTRUCTION %s.>>", instruction.ToString())
}

func Import() string {
	return "OK"
}

func (self *slimExecutor) Make(instruction *SlimList) string {
	instanceName := instruction.GetStringAt(2)
	className := instruction.GetStringAt(3)
	args := instruction.GetTailAt(4)
	return self.executor.Make(instanceName, className, args)
}

func (self *slimExecutor) Call(instruction *SlimList) string {
	if instruction.GetLength() < 4 {
		return MalformedInstruction(instruction)
	}
	instanceName := instruction.GetStringAt(2)
	methodName := instruction.GetStringAt(3)
	args := instruction.GetTailAt(4)
	return self.executor.Call(instanceName, methodName, args)
}

func (self *slimExecutor) CallAndAssign(instruction *SlimList) string {
	if instruction.GetLength() < 5 {
		return MalformedInstruction(instruction)
	}
	symbolName := instruction.GetStringAt(2)
	instanceName := instruction.GetStringAt(3)
	methodName := instruction.GetStringAt(4)
	args := instruction.GetTailAt(5)
	result := self.executor.Call(instanceName, methodName, args)
	self.executor.SetSymbol(symbolName, result)
	return result
}

func (self *slimExecutor) Dispatch(instruction *SlimList) string {
	command := instruction.GetStringAt(1)
	switch command {
	case "import":
		return Import()
	case "make":
		return self.Make(instruction)
	case "call":
		return self.Call(instruction)
	case "callAndAssign":
		return self.CallAndAssign(instruction)
	default:
		return InvalidCommand(instruction)
	}
}

func (self *slimExecutor) Execute(instructions *SlimList) *SlimList {
	results := SlimList_Create()

	for iterator := instructions.CreateIterator(); iterator != nil; iterator = iterator.Advance() {
		instruction := iterator.GetList()
		id := instruction.GetStringAt(0)
		result := self.Dispatch(instruction)
		AddResult(results, id, result)
	}

	return results
}

func (self *slimExecutor) RegisterFixture(className string, constructor Constructor) {
	self.executor.RegisterFixture(className, constructor)
}

func (self *slimExecutor) RegisterMethod(className string, aliasMethodName string, realMethodName string) {
	self.executor.RegisterMethod(className, aliasMethodName, realMethodName)
}
