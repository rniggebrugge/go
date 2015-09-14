package main
import "fmt"

type person struct {
  name string
  age int
}

func changeName(per person){
  per.name = "Remco"
}

func changeNamePTR(per *person){
  per.name = "Clara"
}

func main(){

  fmt.Println(person{"Bob",20})
  fmt.Println(person{name:"Alice", age:30})
  fmt.Println(person{name:"Fred"})
  fmt.Println(&person{name:"Ann",age:40})

  s:= person{name:"Sean", age:50}
  fmt.Println(s.name)

  sp := &s
  fmt.Println(sp.age)

  sp.age = 51
  fmt.Println(sp.age)

  sp.age = 40
  changeName(s)
  fmt.Println(s)
  changeNamePTR(&s)
  fmt.Println(s)
}
