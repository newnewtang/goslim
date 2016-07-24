package slimlist

import (
	. "goslim/gotest"
	"testing"
)

func Test_SlimListSerializer_SerializeAListWithNoElements(t *testing.T) {
	setup(t)

	EXPECT_STREQ("[000000:]", slimList.Serialize())
}

func Test_SlimListSerializer_SerializeAListWithOneElements(t *testing.T) {
	setup(t)

	slimList.AddString("hello")
	EXPECT_STREQ("[000001:000005:hello:]", slimList.Serialize())
}

func Test_SlimListSerializer_SerializeAListWithTwoElements(t *testing.T) {
	setup(t)

	slimList.AddString("hello")
	slimList.AddString("world")

	EXPECT_STREQ("[000002:000005:hello:000005:world:]", slimList.Serialize())
}

func Test_SlimListSerializer_ListCopysItsString(t *testing.T) {
	setup(t)

	str := "Hello"
	slimList.AddString(str)
	str = "Goodbye"

	EXPECT_STREQ("[000001:000005:Hello:]", slimList.Serialize())
}

func Test_SlimListSerializer_canCopyAList(t *testing.T) {
	setup(t)

	slimList.AddString("123456")
	slimList.AddString("987654")

	copyList := SlimList_Create()
	for i := 0; i < slimList.GetLength(); i++ {
		copyList.AddString(slimList.GetStringAt(i))
	}

	EXPECT_STREQ(copyList.Serialize(), slimList.Serialize())
}

func Test_SlimListSerializer_SerializeNestedList(t *testing.T) {
	setup(t)

	embeddedList := SlimList_Create()
	embeddedList.AddString("element")
	slimList.AddList(embeddedList)

	EXPECT_STREQ("[000001:000024:[000001:000007:element:]:]", slimList.Serialize())
}

func Test_SlimListSerializer_serializedLength(t *testing.T) {
	setup(t)

	slimList.AddString("12345")

	EXPECT_EQ(5+17, slimList.SerializedLength())

	slimList.AddString("123456")
	EXPECT_EQ(9+(5+8)+(6+8), slimList.SerializedLength())

	EXPECT_EQ(9+(5+8)+(6+8), len(slimList.Serialize()))
}

func Test_SlimListSerializer_serializeNull(t *testing.T) {
	setup(t)

	slimList.AddString("")
	EXPECT_STREQ("[000001:000004:null:]", slimList.Serialize())

}

func Test_SlimListSerializer_serializeMultibyteCharacters(t *testing.T) {
	setup(t)

	slimList.AddString("Ü€©phewÜ€©")
	EXPECT_STREQ("[000001:000018:Ü€©phewÜ€©:]", slimList.Serialize())
}
