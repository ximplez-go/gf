// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gbuild_test

import (
	"testing"

	"github.com/ximplez-go/gf/frame/g"
	"github.com/ximplez-go/gf/os/gbuild"
	"github.com/ximplez-go/gf/test/gtest"
	"github.com/ximplez-go/gf/util/gconv"
)

func Test_Info(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gconv.Map(gbuild.Info()), g.Map{
			"GoFrame": "",
			"Golang":  "",
			"Git":     "",
			"Time":    "",
			"Version": "",
			"Data":    g.Map{},
		})
	})
}

func Test_Get(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gbuild.Get(`none`), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gbuild.Get(`none`, 1), 1)
	})
}

func Test_Map(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gbuild.Data(), map[string]interface{}{})
	})
}
