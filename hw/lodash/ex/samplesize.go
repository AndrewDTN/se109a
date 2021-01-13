package ex

import (
	"math/rand"
)
//[]interface{} 可視為一種type
func Samplesize (item []interface{}, n int) []interface{}{
	result := []interface{} {}

	for i:=0;i<n;i++{
		length := len(item)
		if length == 0 {
			break
		}	
		r := rand.Intn(length)
		result = append(result,item[r])
		item = append(item[:r],item[r+1:]...)
	}
	return result
} 