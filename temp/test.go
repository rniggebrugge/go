package main 

import (
	"fmt"
)

func main(){
	phrase := "vått og tørt "
	fmt.Printf("string: \"%s\"\n", phrase)
	fmt.Println("index	rune	char	bytes")
	for index, char := range phrase {
		fmt.Printf("%-2d	%U	'%c'	% X\n", 
			index, char, char, []byte(string(char)))
	}

	for i:=0; i<len(phrase);i++ {
		fmt.Printf("%-2d : %c \n",i,phrase[i])
	}
}