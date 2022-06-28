// Copyright 2022 helloshaohua <wu.shaohua@foxmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"sync"

	"github.com/coolstina/merge"
)

type _student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// HandlerStringFunc Default support string merge.
var _handlerFunc merge.HandlerFunc = func(val interface{}, sets *sync.Map) bool {
	var exists bool

	sets.Range(func(key, value interface{}) bool {
		if value.(*_student).Id == val.(*_student).Id {
			exists = true
		}

		return true
	})

	return exists
}

func main() {
	var structure = merge.NewMerge()
	structure.Append(
		&_student{
			Id:   1,
			Name: "tom",
		},
		&_student{
			Id:   2,
			Name: "kitty",
		},
	)

	structure.Append(
		&_student{
			Id:   2,
			Name: "kitty",
		},
		&_student{
			Id:   4,
			Name: "helloshaohua",
		},
	)

	actual := structure.Merged(_handlerFunc)

	// [{"id":2,"name":"kitty"},{"id":4,"name":"helloshaohua"},{"id":1,"name":"tom"}]
	fmt.Printf("%+v\n", actual.String())
}
