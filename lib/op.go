package lib

import (
	"turing/lib/mod"
	"turing/lib/log2"
	"turing/lib/add"
	"turing/lib/sub"
	"turing/lib/mul"
	"turing/lib/fac"
	"turing/lib/div"
)

func Mod(n int, m int)  {
	mod.Mod(n, m)
}

func Log2(n int)  {
	log2.Log2(n)
}

func Add(n int, m int) {
	add.Add(n, m)
}

func Sub(n int, m int)  {
	sub.Sub(n, m)
}

func Mul(n int, m int) {
	mul.Mul(n, m)
}

func Fac(n int) {
	fac.Fac(n)
}

func Div(n int, m int) {
	div.Div(n, m)
}