package slimlist

import (
	. "goslim/gotest"
	"testing"
)

func Test_SlimListDeserializer_deserializeEmptyList(t *testing.T) {
	setup(t)

	deserializedList := SlimList_Deserialize("[000000:]")

	EXPECT_TRUE(deserializedList != nil)
	EXPECT_EQ(0, deserializedList.GetLength())
}

func Test_SlimListDeserializer_deserializeNull(t *testing.T) {
	setup(t)

	deserializedList := SlimList_Deserialize("")

	EXPECT_TRUE(deserializedList == nil)
}

func Test_SlimListDeserializer_deserializeEmptyString(t *testing.T) {
	setup(t)

	deserializedList := SlimList_Deserialize("")

	EXPECT_TRUE(deserializedList == nil)
}

func Test_SlimListDeserializer_MissingOpenBracketReturnsNull(t *testing.T) {
	setup(t)

	deserializedList := SlimList_Deserialize("hello")

	EXPECT_TRUE(deserializedList == nil)
}

func Test_SlimListDeserializer_MissingClosingBracketReturnsNull(t *testing.T) {
	setup(t)

	deserializedList := SlimList_Deserialize("[000000:")

	EXPECT_TRUE(deserializedList == nil)
}

func Test_SlimListDeserializer_canDeserializeCanonicalListWithOneElement(t *testing.T) {
	setup(t)

	canonicalList := "[000001:000008:Hi doug.:]"
	deserializedList := SlimList_Deserialize(canonicalList)
	EXPECT_TRUE(deserializedList != nil)

	EXPECT_EQ(1, deserializedList.GetLength())
	EXPECT_STREQ("Hi doug.", deserializedList.GetStringAt(0))
}

func Test_SlimListDeserializer_canDeserializeWithMultibyteCharacters(t *testing.T) {
	setup(t)

	canonicalList := "[000001:000008:Hi JRÜ€©:]"
	deserializedList := SlimList_Deserialize(canonicalList)
	EXPECT_TRUE(deserializedList != nil)

	EXPECT_EQ(1, deserializedList.GetLength())
	EXPECT_STREQ("Hi JRÜ€©", deserializedList.GetStringAt(0))
}

func Test_SlimListDeserializer_canDeSerializeListWithOneElement(t *testing.T) {
	setup(t)

	slimList.AddString("hello")
	serializedList := slimList.Serialize()
	deserializedList := SlimList_Deserialize(serializedList)

	EXPECT_TRUE(deserializedList != nil)
	EXPECT_SLIMLIST_EQ(slimList, deserializedList)
}

func Test_SlimListDeserializer_canDeSerializeListWithTwoElements(t *testing.T) {
	setup(t)

	slimList.AddString("hello")
	slimList.AddString("bob")
	serializedList := slimList.Serialize()
	deserializedList := SlimList_Deserialize(serializedList)

	EXPECT_TRUE(deserializedList != nil)
	EXPECT_SLIMLIST_EQ(slimList, deserializedList)
}

func Test_SlimListDeserializer_canAddSubList(t *testing.T) {
	setup(t)

	embeddedList := SlimList_Create()
	embeddedList.AddString("element")
	slimList.AddList(embeddedList)

	serializedList := slimList.Serialize()
	deserializedList := SlimList_Deserialize(serializedList)

	subList := deserializedList.GetListAt(0)
	EXPECT_SLIMLIST_EQ(embeddedList, subList)
}

func Test_SlimListDeserializer_getStringWhereThereIsAList(t *testing.T) {
	setup(t)

	embeddedList := SlimList_Create()
	embeddedList.AddString("element")
	slimList.AddList(embeddedList)

	serializedList := slimList.Serialize()
	deserializedList := SlimList_Deserialize(serializedList)
	str := deserializedList.GetStringAt(0)

	EXPECT_STREQ("[000001:000007:element:]", str)
}
