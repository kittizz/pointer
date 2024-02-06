# pointer

[![Go Reference](https://pkg.go.dev/badge/github.com/kittizz/pointer.svg)](https://pkg.go.dev/github.com/kittizz/pointer)

Go package `pointer` provides a convenient way to work with pointers, especially when dealing with basic data types and generics in Go. This package simplifies the creation and manipulation of pointers for various types, leveraging Go's generics feature introduced in Go 1.18. With pointer, developers can easily generate pointers for any given value without the need to manually allocate memory or deal with the verbosity of referencing and dereferencing values.

```
go get github.com/kittizz/pointer
```

example

```go
package main

import "github.com/kittizz/pointer"

func main() {
	pointer := pointer.Of[string]("hello")
	println(pointer, *pointer)
}

```