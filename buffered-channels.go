package main
import "fmt"

func main() {
	messages := make(chan interface{},5)

	messages <- "buffered"
	messages <- "channels"
	messages <- "rule"
	messages <- "the world"
	messages <- 2015

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)	
}