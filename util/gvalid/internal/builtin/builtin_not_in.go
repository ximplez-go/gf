// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package builtin

import (
	"errors"
	"strings"

	"github.com/ximplez-go/gf/text/gstr"
)

// RuleNotIn implements `not-in` rule:
// Value should not be in: value1,value2,...
//
// Format: not-in:value1,value2,...
type RuleNotIn struct{}

func init() {
	Register(RuleNotIn{})
}

func (r RuleNotIn) Name() string {
	return "not-in"
}

func (r RuleNotIn) Message() string {
	return "The {field} value `{value}` must not be in range: {pattern}"
}

func (r RuleNotIn) Run(in RunInput) error {
	var (
		ok    = true
		value = in.Value.String()
	)
	for _, rulePattern := range gstr.SplitAndTrim(in.RulePattern, ",") {
		if in.Option.CaseInsensitive {
			ok = !strings.EqualFold(value, strings.TrimSpace(rulePattern))
		} else {
			ok = strings.Compare(value, strings.TrimSpace(rulePattern)) != 0
		}
		if !ok {
			return errors.New(in.Message)
		}
	}
	return nil
}
