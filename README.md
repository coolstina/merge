# merge

Golang slice merge implement

## Installation

```shell
go get -u github.com/coolstina/merge
```

## Features

- Merge slice
- Custom merge handler

## Example

### [MergeStringSlice](example/integer/integer.go)

```go
package main

import (
	"fmt"

	"github.com/coolstina/merge"
)

func main() {
	var integer = merge.NewMerge()
	integer.Append("80", "90")
	integer.Append("90", "120")
	integer.Append("130", "160")

	actual := integer.Merged(merge.HandlerStringFunc)

	// ["80","90","120","130","160"]
	fmt.Printf("%+v\n", actual.String())
}
```

### [MergeStructSlice](example/integer/integer.go)

```go
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
```


## LICENSE 

- [Apache License](LICENSE)

## Author

- [@helloshaohua](https://github.com/helloshaohua)
