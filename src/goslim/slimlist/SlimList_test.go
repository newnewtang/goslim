package slimlist

import (
	. "goslim/gotest"
	"testing"
)

var slimList *SlimList

func setup(cur *testing.T) {
	SetUp(cur)
	slimList = SlimList_Create()
}

func Test_SlimList_twoEmptyListsAreEqual(t *testing.T) {
	setup(t)

	list := SlimList_Create()

	EXPECT_SLIMLIST_EQ(slimList, list)
}

func Test_SlimList_twoDifferentLenghtListsAreNotEqual(t *testing.T) {
	setup(t)

	list := SlimList_Create()

	slimList.AddString("hello")

	EXPECT_SLIMLIST_NE(slimList, list)
}

func Test_SlimList_twoSingleElementListsWithDifferentElmementsAreNotEqual(t *testing.T) {
	setup(t)

	list := SlimList_Create()

	slimList.AddString("hello")
	list.AddString("goodbye")

	EXPECT_SLIMLIST_NE(slimList, list)
}

func Test_SlimList_twoIdenticalMultipleElementListsElmementsAreEqual(t *testing.T) {
	setup(t)

	list := SlimList_Create()

	slimList.AddString("hello")
	slimList.AddString("goodbye")
	list.AddString("hello")
	list.AddString("goodbye")

	EXPECT_SLIMLIST_EQ(slimList, list)
}

func Test_SlimList_twoNonIdenticalMultipleElementListsElmementsAreNotEqual(t *testing.T) {
	setup(t)

	list := SlimList_Create()

	slimList.AddString("hello")
	slimList.AddString("hello")
	list.AddString("hello")
	list.AddString("goodbye")

	EXPECT_SLIMLIST_NE(slimList, list)
}

func Test_SlimList_canGetElements(t *testing.T) {
	setup(t)

	slimList.AddString("element1")
	slimList.AddString("element2")

	EXPECT_STREQ("element1", slimList.GetStringAt(0))
	EXPECT_STREQ("element2", slimList.GetStringAt(1))
}

func Test_SlimList_canGetHashWithOneElement(t *testing.T) {
	setup(t)

	slimList.AddString("<table><tr><td>name</td><td>bob</td></tr></table>")

	hash := slimList.GetHashAt(0)
	twoElementList := hash.GetListAt(0)

	EXPECT_STREQ("name", twoElementList.GetStringAt(0))
	EXPECT_STREQ("bob", twoElementList.GetStringAt(1))
}

func Test_SlimList_canGetHashWithMultipleElements(t *testing.T) {
	setup(t)
	slimList.AddString("<table><tr><td>name</td><td>dough</td></tr><tr><td>addr</td><td>here</td></tr></table>")

	hash := slimList.GetHashAt(0)
	twoElementList := hash.GetListAt(1)

	EXPECT_STREQ("addr", twoElementList.GetStringAt(0))
	EXPECT_STREQ("here", twoElementList.GetStringAt(1))
}

func Test_SlimList_cannotGetElementThatAreNotThere(t *testing.T) {
	setup(t)

	slimList.AddString("element1")
	slimList.AddString("element2")

	EXPECT_STREQ("", slimList.GetStringAt(3))
}

func Test_SlimList_canReplaceString(t *testing.T) {
	setup(t)

	slimList.AddString("replaceMe")
	slimList.ReplaceAt(0, "WithMe")

	EXPECT_STREQ("WithMe", slimList.GetStringAt(0))
}

func Test_SlimList_canGetTail(t *testing.T) {
	setup(t)

	slimList.AddString("1")
	slimList.AddString("2")
	slimList.AddString("3")
	slimList.AddString("4")

	expected := SlimList_Create()
	expected.AddString("3")
	expected.AddString("4")

	tail := slimList.GetTailAt(2)

	EXPECT_SLIMLIST_EQ(expected, tail)
}

func Test_SlimList_getDouble(t *testing.T) {
	setup(t)

	slimList.AddString("2.3")

	EXPECT_FLOAT_EQ(2.3, slimList.GetDoubleAt(0), 0.1)
}

func Test_SlimList_ToStringForEmptyList(t *testing.T) {
	setup(t)

	EXPECT_STREQ("[]", slimList.ToString())
}

func Test_SlimList_toStringForSimpleList(t *testing.T) {
	setup(t)

	slimList.AddString("a")
	slimList.AddString("b")

	EXPECT_STREQ("[\"a\", \"b\"]", slimList.ToString())
}

func Test_SlimList_toStringDoesNotHaveASideEffectWhichChangesResultsFromPriorCalls(t *testing.T) {
	setup(t)

	priorString := slimList.ToString()

	slimList.AddString("a")
	listWithAnElementAsASting := slimList.ToString()

	EXPECT_STRNE(priorString, listWithAnElementAsASting)
}

func Test_SlimList_recursiveToString(t *testing.T) {
	setup(t)

	slimList.AddString("a")
	slimList.AddString("b")

	sublist := SlimList_Create()
	sublist.AddString("3")
	sublist.AddString("4")

	slimList.AddList(sublist)

	EXPECT_STREQ("[\"a\", \"b\", [\"3\", \"4\"]]", slimList.ToString())
}

func Test_SlimList_toStringForLongList(t *testing.T) {
	setup(t)

	entries := 128 //TODO: consider updating the size given its no longer as relevant

	for i := 0; i < entries; i++ {
		slimList.AddString("a")
	}

	slimList.ToString()
}

func Test_SlimList_CanPopHeadOnListWithOneEntry(t *testing.T) {
	setup(t)

	slimList.AddString("a")
	slimList.PopHead()

	EXPECT_EQ(0, slimList.GetLength())
}

func Test_SlimList_CanInsertAfterPoppingListWithEntries(t *testing.T) {
	setup(t)

	slimList.AddString("a")
	slimList.AddString("a")

	slimList.PopHead()
	slimList.PopHead()

	slimList.AddString("a")
	slimList.AddString("a")
	EXPECT_EQ(2, slimList.GetLength())
}

func Test_SlimList_iteratorDoesNotHaveAnItemWhenEmpty(t *testing.T) {
	setup(t)

	iterator := slimList.CreateIterator()
	EXPECT_TRUE(iterator == nil)
}

func Test_SlimList_iteratorHasItem(t *testing.T) {
	setup(t)

	slimList.AddString("a")

	iterator := slimList.CreateIterator()
	EXPECT_FALSE(iterator == nil)
}

func Test_SlimList_iteratorNext(t *testing.T) {
	setup(t)

	slimList.AddString("a")

	iterator := slimList.head
	iterator = iterator.Advance()
	EXPECT_FALSE(iterator != nil)
}

func Test_SlimList_iteratorGetString(t *testing.T) {
	setup(t)

	slimList.AddString("a")

	iterator := slimList.CreateIterator()
	EXPECT_TRUE(iterator.GetString() == "a")
}
