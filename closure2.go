package main
import "fmt"

func intSeq(start int) func() int {
  i:=start
  max:=10
  return func() int {
    i+=1
    for i>max {
      i-=max
    }
    return i
  }
}

func main(){

  nextInt := intSeq(0)
  for i:=0; i<30; i++ {
    fmt.Println(nextInt())
  }
  fmt.Println("==================================================")

  newInt := intSeq(6)
  for i:=0; i<30; i++ {
    fmt.Println(newInt())
  }
  fmt.Println("==================================================")

  othInt := intSeq(304)
  for i:=0; i<30; i++ {
    fmt.Println(othInt())
  }
  fmt.Println("==================================================")
}
