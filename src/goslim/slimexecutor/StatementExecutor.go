package slimexecutor

import (
	"fmt"
	. "goslim/slimlist"
	"reflect"
	"strings"
)

type Fixture func(*StatementExecutor)
type Constructor func(*SlimList) interface{}

type FixtureNode struct {
	next        *FixtureNode
	name        string
	constructor Constructor
	methods     map[string]string
}

type InstanceNode struct {
	next     *InstanceNode
	name     string
	instance interface{}
	fixture  *FixtureNode
}

type StatementExecutor struct {
	fixtures    *FixtureNode
	instances   *InstanceNode
	symbolTable *SymbolTable
	message     string
}

func StatementExecutor_Create() *StatementExecutor {
	self := new(StatementExecutor)
	self.symbolTable = SymbolTable_Create()
	return self
}

func (executor *StatementExecutor) GetInstanceNode(instanceName string) *InstanceNode {
	for instanceNode := executor.instances; instanceNode != nil; instanceNode = instanceNode.next {
		if compareNamesIgnoreUnderScores(instanceNode.name, instanceName) {
			return instanceNode
		}
	}
	return nil
}

func (executor *StatementExecutor) Make(instanceName string, className string, args *SlimList) string {
	fixtureNode := executor.findFixture(className)
	if fixtureNode != nil {
		instanceNode := new(InstanceNode)
		instanceNode.next = executor.instances
		instanceNode.name = instanceName
		instanceNode.fixture = fixtureNode

		executor.instances = instanceNode
		executor.symbolTable.replaceSymbols(args)

		instanceNode.instance = (fixtureNode.constructor)(args)
		if instanceNode.instance != nil {
			return "OK"
		} else {
			formatString := "__EXCEPTION__:message:<<COULD_NOT_INVOKE_CONSTRUCTOR %.32s .>>"
			return fmt.Sprintf(formatString, className)
		}
	}
	formatString := "__EXCEPTION__:message:<<NO_CLASS %.32s.>>"
	executor.message = fmt.Sprintf(formatString, className)
	return executor.message
}

func (executor *StatementExecutor) Call(instanceName string, methodName string, args *SlimList) string {
	instanceNode := executor.GetInstanceNode(instanceName)
	if instanceNode != nil {
		var realName string

		//1st. find alias methond.
		if aliasName, OK := instanceNode.fixture.methods[methodName]; OK {
			realName = aliasName
		} else {
			realName = methodName
		}

		//method first name upper.
		MethodName := strings.ToUpper(string(realName[0])) + string([]byte(realName)[1:])

		//2nd. find reflect methond.
		v := reflect.ValueOf(instanceNode.instance)
		if v.MethodByName(MethodName).IsValid() != true {
			formatString := "__EXCEPTION__:message:<<NO_METHOD_IN_CLASS %.32s[%d] %.32s.>>"
			executor.message = fmt.Sprintf(formatString, methodName, args.GetLength(), instanceNode.fixture.name)
			return executor.message
		}

		executor.symbolTable.replaceSymbols(args)

		return v.MethodByName(MethodName).Call([]reflect.Value{reflect.ValueOf(args)})[0].String()
	}
	formatString := "__EXCEPTION__:message:<<NO_INSTANCE %.32s.>>"
	executor.message = fmt.Sprintf(formatString, instanceName)
	return executor.message
}

func (executor *StatementExecutor) Instance(instanceName string) interface{} {
	instanceNode := executor.GetInstanceNode(instanceName)
	if instanceNode != nil {
		return instanceNode.instance
	}
	return nil
}

func (executor *StatementExecutor) AddFixture(fixture Fixture) {
	(fixture)(executor)
}

func (executor *StatementExecutor) RegisterFixture(className string, constructor Constructor) {
	fixtureNode := executor.findFixture(className)
	if fixtureNode == nil {
		fixtureNode = new(FixtureNode)
		fixtureNode.next = executor.fixtures
		executor.fixtures = fixtureNode
		fixtureNode.name = className
	}

	fixtureNode.constructor = constructor
}

func (executor *StatementExecutor) findFixture(className string) *FixtureNode {
	var fixtureNode *FixtureNode
	for fixtureNode = executor.fixtures; fixtureNode != nil; fixtureNode = fixtureNode.next {
		if compareNamesIgnoreUnderScores(fixtureNode.name, className) {
			break
		}
	}
	return fixtureNode
}

func (executor *StatementExecutor) RegisterMethod(className string, aliasMethodName string, realMethodName string) {
	fixtureNode := executor.findFixture(className)
	if fixtureNode == nil {
		return
	}

	if fixtureNode.methods == nil {
		fixtureNode.methods = make(map[string]string)
	}

	fixtureNode.methods[aliasMethodName] = realMethodName
	return
}

func (executor *StatementExecutor) SetSymbol(symbol, value string) {
	executor.symbolTable.SetSymbol(symbol, value)
}

func StatementExecutor_FixtureError(message string) string {
	formatString := "__EXCEPTION__:message:<<%.100s.>>"
	return fmt.Sprintf(formatString, message)
}

var Null_Create Constructor = func(args *SlimList) interface{} {
	return nil
}

func compareNamesIgnoreUnderScores(name1, name2 string) bool {
	i, j := 0, 0
	for i < len(name1) && j < len(name2) {
		if name1[i] == name2[j] {
			i++
			j++
		} else if name1[i] == '_' {
			i++
		} else if name2[j] == '_' {
			j++
		} else {
			return false
		}
	}
	if i != len(name1) || j != len(name2) {
		return false
	}
	return true
}
