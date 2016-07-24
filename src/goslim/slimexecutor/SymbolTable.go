package slimexecutor

import (
	. "goslim/slimlist"
	"strings"
)

type SymbolNode struct {
	next  *SymbolNode
	name  string
	value string
}

type SymbolTable struct {
	head *SymbolNode
}

func SymbolTable_Create() *SymbolTable {
	return new(SymbolTable)
}

func (self *SymbolTable) FindSymbol(name string) string {
	for node := self.head; node != nil; node = node.next {
		if strings.Compare(node.name, name) == 0 {
			return node.value
		}
	}
	return ""
}

func (self *SymbolTable) SetSymbol(symbol, value string) {
	symbolNode := new(SymbolNode)
	symbolNode.name = symbol
	symbolNode.value = value
	symbolNode.next = self.head
	self.head = symbolNode
}

func (self *SymbolTable) GetSymbolLength(symbol string) int {
	symbolValue := self.FindSymbol(symbol)
	if symbolValue == "" {
		return -1
	}
	return len(symbolValue)
}

func (symbolTable *SymbolTable) replaceSymbols(list *SlimList) {
	for iterator := list.CreateIterator(); iterator != nil; iterator = iterator.Advance() {
		str := iterator.GetString()
		embeddedList := SlimList_Deserialize(str)
		if embeddedList == nil {
			replacedString := symbolTable.replaceString(str)
			iterator.Replace(replacedString)
		} else {
			symbolTable.replaceSymbols(embeddedList)
			serializedReplacedList := embeddedList.Serialize()
			iterator.Replace(serializedReplacedList)
		}
	}
}

func (symbolTable *SymbolTable) replaceString(str string) string {
	return symbolTable.replaceStringFrom(str, str)
}

func (symbolTable *SymbolTable) replaceStringFrom(str string, from string) string {
	dollarSign := strings.IndexByte(from, '$')
	if dollarSign != -1 {
		length := lengthOfSymbol(from, dollarSign+1)
		symbolValue := symbolTable.FindSymbol(string([]byte(from)[dollarSign+1 : dollarSign+1+length]))
		if symbolValue != "" {
			newString := string([]byte(str)[:dollarSign+len(str)-len(from)])
			newString += symbolValue
			newString += string([]byte(from)[dollarSign+1+length:])

			recursedString := symbolTable.replaceStringFrom(newString, newString)
			return recursedString
		} else {
			if dollarSign == len(from) {
				return str
			}

			return symbolTable.replaceStringFrom(str, string([]byte(from)[dollarSign+1:]))
		}
	}
	return str
}

func lengthOfSymbol(str string, start int) int {
	isalnum := func(c byte) bool {
		return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
	}

	length := 0
	for i := start; i < len(str) && isalnum(str[i]); i++ {
		length++
	}
	return length
}
