package main

import (
	"test/ex"
	"fmt"
)

type item []interface{}

func main (){
	in := item{1, 2, 3}
	fmt.Println(ex.Samplesize(in , 2))
	fmt.Println(ex.Samplesize(in , 4))
	fmt.Println(ex.Samplesize(in , 4))
}
