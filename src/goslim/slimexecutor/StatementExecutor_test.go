package slimexecutor

import (
	"demo/testslim"
	"fmt"
	. "goslim/gotest"
	. "goslim/slimlist"
	"testing"
)

var (
	statementExecutor *StatementExecutor
	args              *SlimList
	empty             *SlimList
	expectedTestSlim  *testslim.TestSlim
)

func registerTestSlim() {
	statementExecutor.RegisterFixture("TestSlim", testslim.TestSlim_Create)
	statementExecutor.RegisterMethod("TestSlim", "echo", "OneArg")

	statementExecutor.RegisterFixture("TestSlimAgain", testslim.TestSlim_Create)
	statementExecutor.RegisterMethod("TestSlimAgain", "setArgAgain", "setArg")
	statementExecutor.RegisterMethod("TestSlimAgain", "getArgAgain", "getArg")

	statementExecutor.RegisterFixture("TestSlimDeclaredLate", testslim.TestSlim_Create)
	statementExecutor.RegisterMethod("TestSlimDeclaredLate", "echo", "OneArg")

	statementExecutor.RegisterMethod("TestSlimUndeclared", "echo", "oneArg")
}

func statementExecutorSetup(cur *testing.T) {
	SetUp(cur)

	args = SlimList_Create()
	empty = SlimList_Create()

	statementExecutor = StatementExecutor_Create()
	registerTestSlim()

	EXPECT_EQ("OK", statementExecutor.Make("test_slim", "TestSlim", empty))

	expectedTestSlim = testslim.TestSlim_Create(args).(*testslim.TestSlim)

	fmt.Println("test started...")
}

func Test_StatementExecutor_canCallFunctionWithNoArguments(t *testing.T) {
	statementExecutorSetup(t)

	statementExecutor.Call("test_slim", "noArgs", args)
	testSlim := statementExecutor.Instance("test_slim").(*testslim.TestSlim)
	EXPECT_TRUE(testSlim.NoArgsCalled() == 1)
}

func Test_StatementExecutor_cantCallFunctionThatDoesNotExist(t *testing.T) {
	statementExecutorSetup(t)

	result := statementExecutor.Call("test_slim", "noSuchMethod", args)
	EXPECT_EQ("__EXCEPTION__:message:<<NO_METHOD_IN_CLASS noSuchMethod[0] TestSlim.>>", result)

	result = statementExecutor.Call("test_slim", "noOtherSuchMethod", args)
	EXPECT_EQ("__EXCEPTION__:message:<<NO_METHOD_IN_CLASS noOtherSuchMethod[0] TestSlim.>>", result)
}

func Test_StatementExecutor_shouldTruncateReallyLongNamedFunctionThatDoesNotExistTo32(t *testing.T) {
	statementExecutorSetup(t)

	result := statementExecutor.Call("test_slim", "noOtherSuchMethod123456789022345678903234567890423456789052345678906234567890", args)
	EXPECT_TRUE(len(result) < 120)
	EXPECT_EQ("__EXCEPTION__:message:<<NO_METHOD_IN_CLASS noOtherSuchMethod123456789022345[0] TestSlim.>>", result)
}

func Test_StatementExecutor_shouldKnowNumberofArgumentsforNonExistantFunction(t *testing.T) {
	statementExecutorSetup(t)

	args.AddString("BlahBlah")
	result := statementExecutor.Call("test_slim", "noSuchMethod", args)
	EXPECT_EQ("__EXCEPTION__:message:<<NO_METHOD_IN_CLASS noSuchMethod[1] TestSlim.>>", result)
}

func Test_StatementExecutor_shouldNotAllowACallToaNonexistentInstance(t *testing.T) {
	statementExecutorSetup(t)

	result := statementExecutor.Call("noSuchInstance", "noArgs", args)
	EXPECT_EQ("__EXCEPTION__:message:<<NO_INSTANCE noSuchInstance.>>", result)
}

func Test_StatementExecutor_shouldNotAllowAMakeOnANonexistentClass(t *testing.T) {
	statementExecutorSetup(t)

	result := statementExecutor.Make("instanceName", "NoSuchClass", empty)
	EXPECT_EQ("__EXCEPTION__:message:<<NO_CLASS NoSuchClass.>>", result)
}

func Test_StatementExecutor_canCallaMethodThatReturnsAValue(t *testing.T) {
	statementExecutorSetup(t)

	result := statementExecutor.Call("test_slim", "returnValue", args)
	EXPECT_EQ("value", result)
}

func Test_StatementExecutor_canCallaMethodThatTakesASlimList(t *testing.T) {
	statementExecutorSetup(t)

	args.AddString("hello world")

	result := statementExecutor.Call("test_slim", "echo", args)
	EXPECT_EQ("hello world", result)
}

func Test_StatementExecutor_WhereCalledFunctionHasUnderscoresSeparatingNameParts(t *testing.T) {
	statementExecutorSetup(t)

	args.AddString("hello world")

	statementExecutor.Call("test_slim", "setArg", args)
	result := statementExecutor.Call("test_slim", "getArgFromFunctionWithUnderscores", empty)
	EXPECT_EQ("hello world", result)
}

func Test_StatementExecutor_canCallTwoInstancesOfTheSameFixture(t *testing.T) {
	statementExecutorSetup(t)

	args2 := SlimList_Create()
	args.AddString("one")
	args2.AddString("two")

	statementExecutor.Make("test_slim2", "TestSlim", empty)
	statementExecutor.Call("test_slim", "setArg", args)
	statementExecutor.Call("test_slim2", "setArg", args2)
	one := statementExecutor.Call("test_slim", "getArg", empty)
	two := statementExecutor.Call("test_slim2", "getArg", empty)
	EXPECT_EQ("one", one)
	EXPECT_EQ("two", two)
}

func Test_StatementExecutor_canCreateTwoDifferentFixtures(t *testing.T) {
	statementExecutorSetup(t)

	args2 := SlimList_Create()
	args.AddString("one")
	args2.AddString("two")

	statementExecutor.Make("test_slim2", "TestSlimAgain", empty)
	statementExecutor.Call("test_slim", "setArg", args)
	statementExecutor.Call("test_slim2", "setArgAgain", args2)
	one := statementExecutor.Call("test_slim", "getArg", empty)
	two := statementExecutor.Call("test_slim2", "getArgAgain", empty)
	EXPECT_EQ("one", one)
	EXPECT_EQ("two", two)
}

func Test_StatementExecutor_canReplaceSymbolsWithTheirValue(t *testing.T) {
	statementExecutorSetup(t)

	statementExecutor.SetSymbol("v", "bob")
	args.AddString("hi $v.")
	result := statementExecutor.Call("test_slim", "echo", args)
	EXPECT_EQ(len("hi bob."), len(result))
	EXPECT_EQ("hi bob.", result)
}

func Test_StatementExecutor_canReplaceSymbolsInTheMiddle(t *testing.T) {
	statementExecutorSetup(t)

	statementExecutor.SetSymbol("v", "bob")
	args.AddString("hi $v whats up.")
	EXPECT_EQ("hi bob whats up.", statementExecutor.Call("test_slim", "echo", args))
}

func Test_StatementExecutor_canReplaceSymbolsWithOtherNonAlphaNumeric(t *testing.T) {
	statementExecutorSetup(t)

	statementExecutor.SetSymbol("v2", "bob")
	args.AddString("$v2=why")
	EXPECT_EQ("bob=why", statementExecutor.Call("test_slim", "echo", args))
}

func Test_StatementExecutor_canReplaceMultipleSymbolsWithTheirValue(t *testing.T) {
	statementExecutorSetup(t)

	statementExecutor.SetSymbol("v", "bob")
	statementExecutor.SetSymbol("e", "doug")
	args.AddString("hi $v. Cost:  $12.32 from $e.")
	EXPECT_EQ("hi bob. Cost:  $12.32 from doug.", statementExecutor.Call("test_slim", "echo", args))
}

func Test_StatementExecutor_canHandlestringWithJustADollarSign(t *testing.T) {
	statementExecutorSetup(t)

	statementExecutor.SetSymbol("v2", "bob")
	args.AddString("$")
	EXPECT_EQ("$", statementExecutor.Call("test_slim", "echo", args))
}

func Test_StatementExecutor_canHandleDollarSignAtTheEndOfTheString(t *testing.T) {
	statementExecutorSetup(t)

	statementExecutor.SetSymbol("v2", "doug")
	args.AddString("hi $v2$")
	EXPECT_EQ("hi doug$", statementExecutor.Call("test_slim", "echo", args))
}

func Test_StatementExecutor_canReplaceSymbolsInSubLists(t *testing.T) {
	statementExecutorSetup(t)

	statementExecutor.SetSymbol("v2", "doug")
	subList := SlimList_Create()
	subList.AddString("Hi $v2.")
	args.AddList(subList)
	result := statementExecutor.Call("test_slim", "echo", args)
	EXPECT_TRUE(result != "")
	returnedList := SlimList_Deserialize(result)
	EXPECT_TRUE(nil != returnedList)
	EXPECT_EQ(1, returnedList.GetLength())
	element := returnedList.GetStringAt(0)
	EXPECT_EQ("Hi doug.", element)
}

func Test_StatementExecutor_canReplaceSymbolsInSubSubLists(t *testing.T) {
	statementExecutorSetup(t)

	statementExecutor.SetSymbol("v2", "doug")
	subList := SlimList_Create()
	subSubList := SlimList_Create()
	subSubList.AddString("Hi $v2.")
	subList.AddList(subSubList)
	args.AddList(subList)
	result := statementExecutor.Call("test_slim", "echo", args)
	EXPECT_TRUE(result != "")
	returnedSubList := SlimList_Deserialize(result)
	EXPECT_TRUE(nil != returnedSubList)
	EXPECT_EQ(1, returnedSubList.GetLength())
	returnedSubSubList := returnedSubList.GetListAt(0)
	EXPECT_TRUE(nil != returnedSubSubList)
	EXPECT_EQ(1, returnedSubSubList.GetLength())
	element := returnedSubSubList.GetStringAt(0)
	EXPECT_TRUE("" != element)
	EXPECT_EQ("Hi doug.", element)
}

func Test_StatementExecutor_canCreateFixtureWithArguments(t *testing.T) {
	statementExecutorSetup(t)

	constructionArgs := SlimList_Create()
	constructionArgs.AddString("hi")
	statementExecutor.Make("test_slim", "TestSlim", constructionArgs)
	result := statementExecutor.Call("test_slim", "getConstructionArg", empty)
	EXPECT_EQ("hi", result)

}

func Test_StatementExecutor_canCreateFixtureWithArgumentsThatHaveSymbols(t *testing.T) {
	statementExecutorSetup(t)

	statementExecutor.SetSymbol("name", "doug")
	constructionArgs := SlimList_Create()
	constructionArgs.AddString("hi $name")
	statementExecutor.Make("test_slim", "TestSlim", constructionArgs)
	result := statementExecutor.Call("test_slim", "getConstructionArg", empty)
	EXPECT_EQ("hi doug", result)

}

func Test_StatementExecutor_canCreateFixtureWithArgumentsThatHaveMultipleSymbols(t *testing.T) {
	statementExecutorSetup(t)

	statementExecutor.SetSymbol("fname", "doug")
	statementExecutor.SetSymbol("lname", "bradbury")

	constructionArgs := SlimList_Create()
	constructionArgs.AddString("hi $fname $lname")
	statementExecutor.Make("test_slim", "TestSlim", constructionArgs)
	result := statementExecutor.Call("test_slim", "getConstructionArg", empty)
	EXPECT_EQ("hi doug bradbury", result)

}

func Test_StatementExecutor_fixtureConstructionFailsWithNoUserErrorMessage(t *testing.T) {
	statementExecutorSetup(t)

	constructionArgs := SlimList_Create()
	constructionArgs.AddString("hi doug")
	constructionArgs.AddString("ho doug")

	result := statementExecutor.Make("test_slim", "TestSlim", constructionArgs)
	EXPECT_EQ("__EXCEPTION__:message:<<COULD_NOT_INVOKE_CONSTRUCTOR TestSlim .>>", result)

}

func Test_StatementExecutor_fixtureCanReturnError(t *testing.T) {
	statementExecutorSetup(t)

	result := statementExecutor.Call("test_slim", "returnError", args)
	EXPECT_EQ("__EXCEPTION__:message:<<my exception.>>", result)
}

func Test_StatementExecutor_canCallFixtureDeclaredBackwards(t *testing.T) {
	statementExecutorSetup(t)

	statementExecutor.Make("backwardsTestSlim", "TestSlimDeclaredLate", empty)
	args.AddString("hi doug")
	result := statementExecutor.Call("backwardsTestSlim", "echo", args)
	EXPECT_EQ("hi doug", result)
}

func Test_StatementExecutor_canNotCallFixtureNotDeclared(t *testing.T) {
	statementExecutorSetup(t)

	statementExecutor.Make("undeclaredTestSlim", "TestSlimUndeclared", empty)
	args.AddString("hi doug")
	result := statementExecutor.Call("undeclaredTestSlim", "echo", args)
	EXPECT_EQ("__EXCEPTION__:message:<<NO_INSTANCE undeclaredTestSlim.>>", result)
}

func Test_StatementExecutor_canHaveNullResult(t *testing.T) {
	statementExecutorSetup(t)

	result := statementExecutor.Call("test_slim", "null", args)
	EXPECT_EQ("", result)
}
