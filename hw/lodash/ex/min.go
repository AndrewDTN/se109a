package ex

import "fmt"

func Min(x []int) {
	var k int
	if len(x) == 0 {
		fmt.Println("undefind")
	} else {
		for i:=1;i<len(x);i++{
			if x[i] < x[0]{
				k = x[i]
			}
		}
		fmt.Println(k)
	}
}