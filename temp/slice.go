package main 

import (
	"fmt"
)

type Product struct {
	name string
	price float64
}

func (product Product) String() string {
	return fmt.Sprintf("%9s (%6.2f)", product.name, product.price)
}

func main() {
	products := []*Product{{"Lego",45.99},{"Knikkers", 2.99}, {"Poppen", 12.54}}
	products = append(products,&Product{"Kleinigheidje", 0.99})
	fmt.Println(products)

	for _, product := range products {
		product.price += 0.50
	}
	fmt.Println(products)

	family := []string{"Clara","Remco"}
	kids := []string{"Willem","Pieter"}
	inM := insertAt(family, kids, 1)
	inM2 := insertAt2(family,kids, 1)
	family = append(family, kids...)

	b := []rune{'U','V'}
	l := []rune{'X','Z'}
	b  = append(b, l...)
	fmt.Printf("%v\n%q\n%v\n%c\n", family, inM, inM2, b)
}

func insertAt(slice, insertion []string, index int) []string {
	result := make([]string, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result	
}

func insertAt2(slice, insertion []string, index int) []string {
	return append(slice[:index], append(insertion, slice[index:]...)...)
}