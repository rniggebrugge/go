package main
import "fmt"

func plus(a, b int) int {
  return a+b
}

func plusPlus(a, b, c int) int {
  return a+b+c
}

func plusAll(args ...int) int {
  sum := 0
  for _, i := range args {
    sum += i
  }
  return sum
}

func main(){
  res := plus(1,2)
  fmt.Println("1 + 2 = ",res)

  res = plusPlus(1,2,3)
  fmt.Println("1 + 2 + 3 = ",res)

  res = plusAll(34,32,775,23,53,-101010,4,5,6,3,6)
  fmt.Println("sum all = ",res)

}

