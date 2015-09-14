package main
import "fmt"

func sum(nums ...int){
  fmt.Println(nums, " ")
  total := 0
  for _, num := range nums {
    total += num
  }
  fmt.Println(total)
}

func main(){
  sum(1,2)
  sum(1,2,3,4,5,6)

  fmt.Println("\n\n")
  nums := []int{1,2,3,4,5,6,7}
  fmt.Println(nums)
  sum(nums...)
}
