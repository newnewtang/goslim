package slimexecutor

import (
	. "goslim/gotest"
	"testing"
)

var symbolTable *SymbolTable

func symbolTableSetup(cur *testing.T) {
	SetUp(cur)
	symbolTable = SymbolTable_Create()
}

func Test_SymbolTable_findNonExistentSymbolShouldReturnNull(t *testing.T) {
	symbolTableSetup(t)

	EXPECT_STREQ("", symbolTable.FindSymbol("Hey"))
}

func Test_SymbolTable_findSymbolShouldReturnSymbol(t *testing.T) {
	symbolTableSetup(t)

	symbolTable.SetSymbol("Hey", "You")
	EXPECT_STREQ("You", symbolTable.FindSymbol("Hey"))
}

func Test_SymbolTable_CanGetLengthOfSymbol(t *testing.T) {
	symbolTableSetup(t)

	symbolTable.SetSymbol("Hey", "1234567890")
	EXPECT_EQ(10, symbolTable.GetSymbolLength("Hey"))
}

func Test_SymbolTable_CanGetLengthOfNonExistentSymbol(t *testing.T) {
	symbolTableSetup(t)

	EXPECT_EQ(-1, symbolTable.GetSymbolLength("Hey"))
}
