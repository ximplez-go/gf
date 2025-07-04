// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package builtin

import (
	"errors"

	"github.com/ximplez-go/gf/internal/json"
)

// RuleArray implements `array` rule:
// Value should be type of array.
//
// Format: array
type RuleArray struct{}

func init() {
	Register(RuleArray{})
}

func (r RuleArray) Name() string {
	return "array"
}

func (r RuleArray) Message() string {
	return "The {field} value `{value}` is not of valid array type"
}

func (r RuleArray) Run(in RunInput) error {
	if in.Value.IsSlice() {
		return nil
	}
	if json.Valid(in.Value.Bytes()) {
		value := in.Value.String()
		if len(value) > 1 && value[0] == '[' && value[len(value)-1] == ']' {
			return nil
		}
	}
	return errors.New(in.Message)
}
