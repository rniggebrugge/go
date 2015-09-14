package main
import "fmt"

func vals() (int, int){
  fmt.Println("..on it..")
  return 4,123
}

func main(){
  a,b := vals()
  fmt.Println(a)
  fmt.Println(b)


  _,c := vals()
  fmt.Println(c)

  vals()

}   
