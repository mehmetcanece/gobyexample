package main

import "fmt"

func main(){

	var a ="initial"
	fmt.Println(a)

	var b,c int = 1,2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int  // e tanımlanmış ama değer verilmemiş o yüzden 0 olarak yazdırılır
	fmt.Println(e)

	var z string // z değişkeni tanımlanmış ama değer verilmemiş o yüzden boş olarak yazdırılır
	fmt.Println(z)

	f := "apple" // := kısayolu ile değişken tanımlama, bu sadece fonksiyon içinde kullanılabilir.
	fmt.Println(f)


}