package main

import (
	"fmt"
)

func main(){
	type polar struct{ radius, θ float64 }
	p := polar{8.32, .49}
	fmt.Print(-18.5, 17, "Elephant", -8+.7i, 0x3C7, '\u03C7', "a", "b", p)
	fmt.Println()
	fmt.Println(-18.5, 17, "Elephant", -8+.7i, 0x3C7, '\u03C7', "a", "b", p)
	c := -232+42.2i
	fmt.Println(c)
}