package main 
import "fmt"
import "time"

func adjust(getal *int){
	*getal++
}

func worker(done chan bool){
	fmt.Println("working...")
	time.Sleep(1*time.Second)
	fmt.Println("done!")
	done<-true
}

func main(){
	getal := 10
	adjust(&getal)
	fmt.Println(getal)
	done := make(chan bool, 1)
	go worker(done)
	<-done
}