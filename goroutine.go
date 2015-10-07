package main 
import "fmt"

func f(from string) {
	for i:=0;i<100;i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")
	go f("goroutine")
	go f("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}