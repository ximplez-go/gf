// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gproperties_test

import (
	"strings"
	"testing"

	"github.com/ximplez-go/gf/encoding/gproperties"
	"github.com/ximplez-go/gf/test/gtest"
)

var pStr string = `
# 模板引擎目录
viewpath = "/home/www/templates/"
# redis数据库配置
redis.disk  = "127.0.0.1:6379,0"
redis.cache = "127.0.0.1:6379,1"
#SQL配置
sql.mysql.0.type = mysql
sql.mysql.0.ip = 127.0.0.1
sql.mysql.0.user = root
`

var errorTests = []struct {
	input, msg string
}{
	// unicode literals
	{"key\\u1 = value", "invalid unicode literal"},
	{"key\\u12 = value", "invalid unicode literal"},
	{"key\\u123 = value", "invalid unicode literal"},
	{"key\\u123g = value", "invalid unicode literal"},
	{"key\\u123", "invalid unicode literal"},

	// circular references
	{"key=${key}", `circular reference in:\nkey=\$\{key\}`},
	{"key1=${key2}\nkey2=${key1}", `circular reference in:\n(key1=\$\{key2\}\nkey2=\$\{key1\}|key2=\$\{key1\}\nkey1=\$\{key2\})`},

	// malformed expressions
	{"key=${ke", "malformed expression"},
	{"key=valu${ke", "malformed expression"},
}

func TestDecode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := make(map[string]interface{})
		m["properties"] = pStr
		res, err := gproperties.Encode(m)
		if err != nil {
			t.Errorf("encode failed. %v", err)
			return
		}
		decodeMap, err := gproperties.Decode(res)
		if err != nil {
			t.Errorf("decode failed. %v", err)
			return
		}
		t.Assert(decodeMap["properties"], pStr)
	})

	gtest.C(t, func(t *gtest.T) {
		for _, v := range errorTests {
			_, err := gproperties.Decode(([]byte)(v.input))
			if err == nil {
				t.Errorf("encode should be failed. %v", err)
				return
			}
			t.AssertIN(`Lib magiconair load Properties data failed.`, strings.Split(err.Error(), ":"))
		}
	})
}

func TestEncode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := make(map[string]interface{})
		m["properties"] = pStr
		res, err := gproperties.Encode(m)
		if err != nil {
			t.Errorf("encode failed. %v", err)
			return
		}
		decodeMap, err := gproperties.Decode(res)
		if err != nil {
			t.Errorf("decode failed. %v", err)
			return
		}
		t.Assert(decodeMap["properties"], pStr)
	})
}
