# Generics

Package `generics` contains type-safe golang struct based on standard library struct.

- [SyncMap](syncmap.go) contains a type safe implementation of golang [`sync.Map`](https://pkg.go.dev/sync?utm_source=godoc#Map).


## Usage


### `SyncMap`


```go
package main

import (
    "fmt"

    "github.com/quentinlesceller/generics"
)

func main() {
    // Instantiate a new SyncMap
    var m generics.SyncMap[string, int32]

    // Store some data
	m.Store("hello", 1)
	m.Store("world", 2)

    // Range over it
	m.Range(func(k string, v int32) bool {
		fmt.Println(k, v)
		return true
	})
 }