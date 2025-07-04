// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconv_test

import (
	"testing"

	"github.com/ximplez-go/gf/test/gtest"
	"github.com/ximplez-go/gf/util/gconv"
)

var (
	boolTestTrueValue  = true
	boolTestFalseValue = false
)

var boolTests = []struct {
	value  interface{}
	expect bool
}{
	{true, true},
	{false, false},

	{0, false},
	{1, true},

	{[]byte(""), false},

	{"", false},
	{"0", false},
	{"1", true},
	{"123.456", true},
	{"true", true},
	{"false", false},
	{"on", true},
	{"off", false},

	{complex(1, 2), true},
	{complex(123.456, 789.123), true},

	{[3]int{1, 2, 3}, true},
	{[]int{1, 2, 3}, true},

	{map[int]int{1: 1}, true},
	{map[string]string{"Earth": "印度洋"}, true},

	{struct{}{}, true},
	{&struct{}{}, true},
	{nil, false},
	{(*bool)(nil), false},

	{&boolTestTrueValue, true},
	{&boolTestFalseValue, false},

	{myBool(true), true},
	{myBool(false), false},
	{(*myBool)(&boolTestTrueValue), true},
	{(*myBool)(&boolTestFalseValue), false},

	{(*myBool)(nil), false},
}

func TestBool(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for _, test := range boolTests {
			t.AssertEQ(gconv.Bool(test.value), test.expect)
		}
	})
}
