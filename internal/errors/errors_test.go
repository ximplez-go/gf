// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package errors_test

import (
	"testing"

	"github.com/ximplez-go/gf/internal/errors"
	"github.com/ximplez-go/gf/test/gtest"
)

func Test_IsStackModeBrief(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(errors.IsStackModeBrief(), true)
	})
}
