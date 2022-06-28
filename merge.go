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
)

type Merge struct {
	store []interface{}
}

func (merge *Merge) Append(values ...interface{}) *Merge {
	merge.store = append(merge.store, values...)

	return merge
}

func (merge *Merge) Merged(fn HandlerFunc) Merged {
	var sets = &sync.Map{}
	var merged Merged

	// Filter
	for idx, item := range merge.store {
		// If HandlerFunc return false, then save store
		if !fn(item, sets) {
			sets.Store(idx, item)
		}
	}

	// Range sync map
	sets.Range(func(key, value interface{}) bool {
		if merged == nil {
			merged = make(Merged, 0)
		}

		merged = append(merged, value)

		return true
	})

	return merged
}

func NewMerge() *Merge {
	return &Merge{store: make([]interface{}, 0)}
}
