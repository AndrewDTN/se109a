package main

import (
	"fmt"
	"test/ex"
)

func main() {
	fmt.Println(ex.Clamp(-10,-5,5))
	fmt.Println(ex.Clamp(10,-5,5))
	fmt.Println(ex.Clamp(1,7,8))
}