package add

import (
	"time"
	"strings"
)

func Add(n int, m int) {
	var s string = "B"
	for i := 0; i < n; i++ {
		s += "0"
	}
	s += "C"
	for i := 0; i < m; i++ {
		s += "0"
	}
	s += "B"
	q0(s, 1)
}

func q0(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q0")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		s = s[:pos]+"B"+s[pos+1:]
		q1(s, pos+1)
	} else if(s[pos] == 'C') {
		s = s[:pos]+"B"+s[pos+1:]
		q5(s, pos+1)
	} else {
		println("Halted")
	}
}

func q1(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q1")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q1(s, pos+1)
	} else if(s[pos] == 'C') {
		q2(s, pos+1)
	} else {
		println("Halted")
	}
}

func q2(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q2")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q2(s, pos+1)
	} else if(s[pos] == 'B') {
		s = s[:pos]+"0B"
		q3(s, pos-1)
	} else {
		println("Halted")
	}
}

func q3(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q3")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q3(s, pos-1)
	} else if(s[pos] == 'C') {
		q4(s, pos-1)
	} else {
		println("Halted")
	}
}

func q4(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q4")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q4(s, pos-1)
	} else if(s[pos] == 'B') {
		q0(s, pos+1)
	} else {
		println("Halted")
	}
}

func q5(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q5")
	println("Finished, result =", strings.Count(s, "0"))
}