package main

import (
	"fmt" 
	"strings"
)

func main(){
	phrases :=  []string{"Remco Clara Willem Pieter", "Remco\tClara\tWillem\tPieter","Remco*Pieter Clara|Willem"}
	for _, record := range phrases {
		fmt.Printf("%q -> %q\n",record, strings.FieldsFunc(record, func(char rune) bool {
			switch char {
				case ' ', '\t', '*', '|': 
					return true
				}
			return false
			}))
   }
	fmt.Println("-....-")
	for _, record := range phrases {
		fmt.Printf("%q -> %q\n", record, strings.Fields(record))
	}	
	f := func(char rune) bool {
		if char =='e' {
		return false
		}
		return true
	}
	fmt.Println("-....-")
	for _, record := range phrases {
		fmt.Printf("%q -> %q\n", record, strings.FieldsFunc(record,f))
	}	
}
