package main

import "github.com/kittizz/pointer"

func main() {
	pointer := pointer.Of[string]("hello")
	println(pointer, *pointer)
}
