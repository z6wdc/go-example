package main

import "fmt"

const big = 9223372036854775807 + 1 // untyped constant

// var big2 int = 9223372036854775807 + 1

func main() {
	fmt.Println(big - 1)
}
