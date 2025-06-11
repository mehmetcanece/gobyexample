package main

import "fmt"


func main(){
	
	var a [5]int
	fmt.Println("emp:", a) //değer vermediğimizden 5 0 koyar

	a[4] = 100
	fmt.Println("set:", a) //değer verdik
	fmt.Println("get:", a[4]) //değerini alırız

	fmt.Println("len:", len(a))

	b:= [5]int {1,2,3,4,5}
	fmt.Println("dcl:", b) //declare

	c:= [...]int{1,2,3,4,5,6} //kendisi sayar kaç eleman olduğunu
	fmt.Println("dcl:",c)

	d:= [...]int{100, 3: 400,500} // 3. indexe 400 koy oraya kadar 0la doldur demek.
	fmt.Println("idx:", d) 

	m := [...]int{1,2, 5:45, 9:33,9,15}
	fmt.Println(m)

	var twoD [2][3]int
	for i:= range 2{
		for j:= range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	twoD = [2][3]int {
		{1,2,3},
		{1,2,3},
	}
	fmt.Println(twoD)









}