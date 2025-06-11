package main

import "fmt"

func main() {

	if 7%2 == 0{
		fmt.Println("7 çifttir")
	} else {
		fmt.Println("7 tektir")
	}
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("8 ya da 7 çifttir")
	}

	if num := 9; num <0 {
		fmt.Println("num negatif")
	}	else if num <10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}