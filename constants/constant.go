package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main(){
	fmt.Println(s)

	const n = 500000000 //constu var koyabildiğin her yere koyabilirsin.

/*
const ifadesi:
Go derleyicisi tarafından sınırsız hassasiyetle işlenir (çok büyük/küçük değerler desteklenir).
var ifadesi:
Belirli bir tipe bağlı olduğu için tipin sınırlarıyla kısıtlıdır. mesela işte float64 gibi.
*/
	const d = 3e20 / n 
	fmt.Println(d)
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

