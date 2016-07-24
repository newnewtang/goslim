package slimlist

import (
	. "goslim/gotest"
)

func EXPECT_SLIMLIST_EQ(expected *SlimList, actual *SlimList) {
	EXPECT_TRUE(actual.Equals(expected))
}
func EXPECT_SLIMLIST_NE(expected *SlimList, actual *SlimList) {
	EXPECT_FALSE(actual.Equals(expected))
}
