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

package merge

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestMergeSuite(t *testing.T) {
	suite.Run(t, &MergeSuite{})
}

type MergeSuite struct {
	suite.Suite
	intMerge    *Merge
	structMerge *Merge
}

func (suite *MergeSuite) BeforeTest(suiteName, testName string) {
	suite.intMerge = NewMerge()
	suite.structMerge = NewMerge()
}

func (suite *MergeSuite) Test_IntMerge() {
	var _ = `["120","130","160","80","90"]`

	suite.intMerge.Append("80", "90")
	suite.intMerge.Append("90", "120")
	suite.intMerge.Append("130", "160")

	actual := suite.intMerge.Merged(HandlerStringFunc)
	assert.Len(suite.T(), actual.Value(), 5)
}

func (suite *MergeSuite) Test_StringMerge() {
	suite.intMerge.Append("tom", "kitty")
	suite.intMerge.Append("kitty", "darwin")
	suite.intMerge.Append("tom", "windows", "linux")
	suite.intMerge.Append("darwin", "windows")

	actual := suite.intMerge.Merged(HandlerStringFunc)
	assert.Len(suite.T(), actual.Value(), 5)
}

type _student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// HandlerStringFunc Default support string merge.
var _HandlerStructFunc HandlerFunc = func(val interface{}, sets *sync.Map) bool {
	var exists bool

	sets.Range(func(key, value interface{}) bool {
		if value.(*_student).Id == val.(*_student).Id {
			exists = true
		}

		return true
	})

	return exists
}

func (suite *MergeSuite) Test_StructMerge() {
	suite.structMerge.Append(
		&_student{
			Id:   1,
			Name: "tom",
		},
		&_student{
			Id:   2,
			Name: "kitty",
		},
	)

	suite.structMerge.Append(
		&_student{
			Id:   2,
			Name: "kitty",
		},
		&_student{
			Id:   4,
			Name: "helloshaohua",
		},
	)

	actual := suite.structMerge.Merged(_HandlerStructFunc)
	assert.Len(suite.T(), actual.Value(), 3)
}
