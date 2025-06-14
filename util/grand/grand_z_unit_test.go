// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// go test *.go -bench=".*"

package grand_test

import (
	"strings"
	"testing"
	"time"

	"github.com/ximplez-go/gf/test/gtest"
	"github.com/ximplez-go/gf/text/gstr"
	"github.com/ximplez-go/gf/util/grand"
)

func Test_Intn(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 1000000; i++ {
			n := grand.Intn(100)
			t.AssertLT(n, 100)
			t.AssertGE(n, 0)
		}
		for i := 0; i < 1000000; i++ {
			n := grand.Intn(-100)
			t.AssertLE(n, 0)
			t.Assert(n, -100)
		}
	})
}

func Test_Meet(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(grand.Meet(100, 100), true)
		}
		for i := 0; i < 100; i++ {
			t.Assert(grand.Meet(0, 100), false)
		}
		for i := 0; i < 100; i++ {
			t.AssertIN(grand.Meet(50, 100), []bool{true, false})
		}
	})
}

func Test_MeetProb(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(grand.MeetProb(1), true)
		}
		for i := 0; i < 100; i++ {
			t.Assert(grand.MeetProb(0), false)
		}
		for i := 0; i < 100; i++ {
			t.AssertIN(grand.MeetProb(0.5), []bool{true, false})
		}
	})
}

func Test_N(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(grand.N(1, 1), 1)
		}
		for i := 0; i < 100; i++ {
			t.Assert(grand.N(0, 0), 0)
		}
		for i := 0; i < 100; i++ {
			t.AssertIN(grand.N(1, 2), []int{1, 2})
		}
	})
}

func Test_D(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(grand.D(time.Second, time.Second), time.Second)
		}
		for i := 0; i < 100; i++ {
			t.Assert(grand.D(0, 0), time.Duration(0))
		}
		for i := 0; i < 100; i++ {
			t.AssertIN(
				grand.D(1*time.Second, 3*time.Second),
				[]time.Duration{1 * time.Second, 2 * time.Second, 3 * time.Second},
			)
		}
	})
}

func Test_Rand(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(grand.N(1, 1), 1)
		}
		for i := 0; i < 100; i++ {
			t.Assert(grand.N(0, 0), 0)
		}
		for i := 0; i < 100; i++ {
			t.AssertIN(grand.N(1, 2), []int{1, 2})
		}
		for i := 0; i < 100; i++ {
			t.AssertIN(grand.N(-1, 2), []int{-1, 0, 1, 2})
		}
	})
}

func Test_S(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(grand.S(5)), 5)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(grand.S(5, true)), 5)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(len(grand.S(0)), 0)
	})
}

func Test_B(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			b := grand.B(5)
			t.Assert(len(b), 5)
			t.AssertNE(b, make([]byte, 5))
		}
	})
	gtest.C(t, func(t *gtest.T) {
		b := grand.B(0)
		t.AssertNil(b)
	})
}

func Test_Str(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(grand.S(5)), 5)
		}
	})
}

func Test_RandStr(t *testing.T) {
	str := "我爱GoFrame"
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 10; i++ {
			s := grand.Str(str, 100000)
			t.Assert(gstr.Contains(s, "我"), true)
			t.Assert(gstr.Contains(s, "爱"), true)
			t.Assert(gstr.Contains(s, "G"), true)
			t.Assert(gstr.Contains(s, "o"), true)
			t.Assert(gstr.Contains(s, "F"), true)
			t.Assert(gstr.Contains(s, "r"), true)
			t.Assert(gstr.Contains(s, "a"), true)
			t.Assert(gstr.Contains(s, "m"), true)
			t.Assert(gstr.Contains(s, "e"), true)
			t.Assert(gstr.Contains(s, "w"), false)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(grand.Str(str, 0), "")
	})
	gtest.C(t, func(t *gtest.T) {
		list := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
		str := ""
		for _, s := range list {
			tmp := ""
			for i := 0; i < 15; i++ {
				tmp += tmp + s
			}
			str += tmp
		}
		t.Assert(len(grand.Str(str, 300)), 300)
	})
}

func Test_Digits(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(grand.Digits(5)), 5)
		}
	})
}

func Test_RandDigits(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(grand.Digits(5)), 5)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(len(grand.Digits(0)), 0)
	})
}

func Test_Letters(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(grand.Letters(5)), 5)
		}
	})
}

func Test_RandLetters(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(grand.Letters(5)), 5)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(len(grand.Letters(0)), 0)
	})
}

func Test_Perm(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.AssertIN(grand.Perm(5), []int{0, 1, 2, 3, 4})
		}
	})
}

func Test_Symbols(t *testing.T) {
	symbols := "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			syms := []byte(grand.Symbols(5))
			for _, sym := range syms {
				t.AssertNE(strings.Index(symbols, string(sym)), -1)
			}
		}
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(grand.Symbols(0), "")
	})
}
