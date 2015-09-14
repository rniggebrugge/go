package main
import "fmt"

func main(){

  m := make(map[string]int)
  fmt.Println("emp: ",m)

  m["k1"]=7
  m["k2"]=13

  fmt.Println("map: ",m)

  v1 := m["k1"]
  fmt.Println("v1:  ",v1)
  fmt.Println("len: ",len(m))

  delete(m, "k2")
  fmt.Println("map: ",m)

  valueAtk2, prs := m["k2"]
  fmt.Println("v at k2: ",valueAtk2)
  fmt.Println("prs: ",prs)

  n := map[string]int{"foo":1,"bar":2}
  fmt.Println("map: ",n)

  n["other"]=101
  fmt.Println("map: ",n)
}
