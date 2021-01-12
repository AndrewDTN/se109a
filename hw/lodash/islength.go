package main

import "test/ex"

func main (){
	ex.IsLength(3)
	ex.IsLength(Number.MIN_VALUE)
	ex.IsLength(Infinity)
	ex.IsLength('3')
}