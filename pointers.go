package main
import "fmt"

func zeroval(ival int){
  ival = 0
}

func zeroptr(iptr *int){
  *iptr = 666 
}

func main(){
  i := 1
  fmt.Println("initial: ",i)

  zeroval(i)
  fmt.Println("zeroval: ",i)

  zeroptr(&i)
  fmt.Println("zeroptr: ",i)

  fmt.Println("pointer: ", &i)

  fmt.Println("------------------==================------------------")
  arr := []int{1,2,3,4}
  for idx, val := range arr {
     fmt.Printf("%d -> %d \n", idx, val)
     fmt.Println("pointer: ", &arr[idx])
  }
}
