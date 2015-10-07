package main

import (
	"fmt"
)

type composer struct{
	name		string
	birthYear	int
}

func main(){
	antonio := composer{"Antonio Teixeira", 1707}
	agnes := new(composer)
	agnes.name, agnes.birthYear = "Agnes Zimmermann", 1845
	julia := &composer{}
	julia.name, julia.birthYear = "Julia Ward Howe", 1819
	augusta := &composer{"Augusta Holmes", 1847}
	remco := &composer{}

	fmt.Printf("%T = %v\n", antonio, antonio)
	fmt.Printf("%T = %v\n", agnes, agnes)
	fmt.Printf("%T = %v\n", julia, julia)
	fmt.Printf("%T = %v\n", augusta, augusta)
	fmt.Printf("%T = %v\n", remco, remco)

}
