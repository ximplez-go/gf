// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with gm file,
// You can obtain one at https://github.com/gogf/gf.

package gset_test

import (
	"fmt"

	"github.com/ximplez-go/gf/container/gset"
)

func ExampleSet_AddIfNotExist() {
	var set gset.Set
	fmt.Println(set.AddIfNotExist(1))
	fmt.Println(set.AddIfNotExist(1))
	fmt.Println(set.Slice())

	// Output:
	// true
	// false
	// [1]
}

func ExampleSet_Pop() {
	var set gset.Set
	set.Add(1, 2, 3, 4)
	fmt.Println(set.Pop())
	fmt.Println(set.Pops(2))
	fmt.Println(set.Size())

	// May Output:
	// 1
	// [2 3]
	// 1
}

func ExampleSet_Pops() {
	var set gset.Set
	set.Add(1, 2, 3, 4)
	fmt.Println(set.Pop())
	fmt.Println(set.Pops(2))
	fmt.Println(set.Size())

	// May Output:
	// 1
	// [2 3]
	// 1
}

func ExampleSet_Join() {
	var set gset.Set
	set.Add("a", "b", "c", "d")
	fmt.Println(set.Join(","))

	// May Output:
	// a,b,c,d
}

func ExampleSet_Contains() {
	var set gset.StrSet
	set.Add("a")
	fmt.Println(set.Contains("a"))
	fmt.Println(set.Contains("A"))
	fmt.Println(set.ContainsI("A"))

	// Output:
	// true
	// false
	// true
}

func ExampleSet_ContainsI() {
	var set gset.StrSet
	set.Add("a")
	fmt.Println(set.Contains("a"))
	fmt.Println(set.Contains("A"))
	fmt.Println(set.ContainsI("A"))

	// Output:
	// true
	// false
	// true
}
