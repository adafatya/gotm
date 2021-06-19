package main

import (
	"fmt"
	"turing/lib"
)

func main() {
	println("Simulasi Turing Machine")
	println("1. Penjumlahan 2 bilangan positif")
	println("2. Pengurangan 2 bilangan")
	println("3. Perkalian 2 bilangan positif")
	println("4. Pembagian 2 bilangan positif")
	println("5. Faktorial dari sebuah bilangan")
	println("6. Modulo 2 bilangan positif")
	println("8. Logaritma basis 2 dari n")
	print("Masukkan pilihan menu : ")
	var x int
	fmt.Scanf("%d ", &x)
	switch x {
		case 1:
			{
				println("Anda memilih menu no 1")
				println("Silahkan masukkan 2 bilangan positif")
				print("input : ")
				var n, m int
				fmt.Scanf("%d %d ", &n, &m)
				lib.Add(n, m)
			}
		case 2:
			{
				println("Anda memilih menu no 2")
				println("Silahkan masukkan 2 bilangan")
				print("Input : ")
				var n, m int
				fmt.Scanf("%d %d ", &n, &m)
				lib.Sub(n, m)
			}
		case 3:
			{
				println("Anda memilih menu no 3")
				println("Silahkan masukkan 2 bilangan positif")
				print("Input : ")
				var n, m int
				fmt.Scanf("%d %d ", &n, &m)
				lib.Mul(n, m)
			}
		case 4:
			{
				println("Anda memilih menu no 4")
				println("Silahkan masukkan 2 bilangan positif")
				print("Input : ")
				var n, m int
				fmt.Scanf("%d %d ", &n, &m)
				lib.Div(n, m)
			}
		case 5:
			{
				println("Anda memilih menu no 5")
				println("Silahkan masukkan bilangan positif")
				print("input : ")
				var n int
				fmt.Scanf("%d ", &n)
				lib.Fac(n)
			}
		case 6:
			{
				println("Anda memilih menu no 6")
				println("Silahkan masukkan 2 bilangan positif")
				print("Input : ")
				var n, m int
				fmt.Scanf("%d %d ", &n, &m)
				lib.Mod(n, m)
			}
		case 8:
			{
				println("Anda memilih menu no 8")
				print("Input : ")
				var n int
				fmt.Scanf("%d ", &n)
				lib.Log2(n)
			}
	}
}