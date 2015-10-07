package main 

import "fmt"
import "time"

func main() {
	
	start := time.Now()

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go func(){
		time.Sleep(time.Second *1 )
		c1 <- "one"
	}()

	go func(){
		time.Sleep(time.Second * 4)
		c3 <- "three"
	}()

	go func(){
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i:=0; i<3; i++ {
		fmt.Println(i)
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case msg3 := <-c3:
			fmt.Println("received", msg3)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("\n\nIt took %s \n\n", elapsed)

}