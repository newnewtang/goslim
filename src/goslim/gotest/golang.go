package gotest

import (
	"math"
	"strings"
	"testing"
)

var t *testing.T

func SetUp(cur *testing.T) {
	t = cur
}

func EXPECT_EQ(expected interface{}, actual interface{}) {
	if expected != actual {
		t.Error("\n expected = ", expected, "\n actual   = ", actual, "\n")
	}
}

func EXPECT_TRUE(condition bool) {
	if condition != true {
		t.Fail()
	}
}

func EXPECT_FALSE(condition bool) {
	EXPECT_TRUE(!condition)
}

func EXPECT_STREQ(expected string, actual string) {
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("expected = %s, actual = %s \n", expected, actual)
	}
}

func EXPECT_STRNE(expected string, actual string) {
	if strings.Compare(expected, actual) == 0 {
		t.Fail()
	}
}

func EXPECT_FLOAT_EQ(expected float64, actual float64, p float64) {
	if math.Dim(expected, actual) >= p {
		t.Fail()
	}
}
