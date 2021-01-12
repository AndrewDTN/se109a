package ex

import "fmt"

func IsLength(x interface{}) {
	switch x.(type) {
	case int64:
		fmt.Println("true")
	case string:
		fmt.Println("dalse")
	}
}