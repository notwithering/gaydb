# gaydb

the worlds gayest database for go that is just messagepack tied to a file

gaydb provides a total of 2 methods, function Get to read the gaydb file into a struct and function Put to put a struct into a gaydb file

## example

this example shows a program that adds 1 to a counter in the database every time its ran

```go
package main

import (
	"fmt"

	"github.com/notwithering/gaydb"
)

func main() {
	var e = struct {
		Counter int
	}{Counter: 0}

	if err := gaydb.Get("database.gay", &e); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(e.Counter)
	e.Counter++

	if err := gaydb.Put("database.gay", &e); err != nil {
		fmt.Println(err)
		return
	}
}

```