package main

import (
	"fmt"
)

func main(){

	x := 300
	y := 22
	pi := &x
	x++
	*pi++
	pi = &y
	*pi++
	ppi := &pi
	**ppi++
	fmt.Printf("%d %d %d %v %v\n", x, y, *pi, pi, &x)
	product := 0
	swapAndProduct(&x,&y,&product)
	fmt.Printf("%d * %d =  %d\n\n", x,y,product)
}

func swapAndProduct(x,y,product *int){
	if *x > *y {
		*x, *y = *y, *x
	}
	*product = *x**y
}
