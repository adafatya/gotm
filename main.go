package main

import (
	"fmt"
	"turing/lib"
)

func main() {
	println("Simulasi Turing Machine")
	println("6. Modulo 2 bilangan positif")
	print("Masukkan pilihan menu : ")
	var x int
	fmt.Scanf("%d ", &x)
	switch x {
	case 6:
		{
			println("Anda memilih menu modulo")
			println("Silahkan masukkan 2 bilangan positif")
			print("Input : ")
			var n, m int
			fmt.Scanf("%d %d ", &n, &m)
			lib.Mod(n, m)
		}
	}
}