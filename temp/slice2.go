package main 

import (
	"fmt"
)

func main() {
	a:=[]int{1,2,3,4,5,6}
	fmt.Printf("%3d\n",a)
	b := removeItems(a, 0,3)
	fmt.Printf("%3d\n",a)
	fmt.Printf("%3d\n",b)
}

func removeItems(s []int, start, end int) []int {
	result := make([]int, len(s)+start-end)
	at := copy(result, s[:start])
	copy(result[at:],s[end:])
	return result
}