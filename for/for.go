package main

import "fmt"

func main(){

	i := 1
	for i<=3 {
		fmt.Println(i)
		i++
	}

	for j := 0; j<3; j++{
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for m := 7 ; m<= 121 ; m++{
		if m%7 == 0 && m%5 == 3{
			fmt.Println(m+2)
		}
		fmt.Println(m)
  }
	for n := range 6 {
		if n%2 == 0{
			continue
		}
		fmt.Println(2)
	}
	


}