package main

import (
	"fmt"
	"turing/lib"
)

func main() {
	println("Simulasi Turing Machine")
	println("6. Modulo 2 bilangan positif")
	println("8. Logaritma basis 2 dari n")
	print("Masukkan pilihan menu : ")
	var x int
	fmt.Scanf("%d ", &x)
	switch x {
		case 6:
			{
				println("Anda memilih menu menu no 6")
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