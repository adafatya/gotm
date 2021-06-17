package mod

import (
	"time"
	"strings"
)

func Mod(n int, m int) {
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
		q0(s, pos+1)
	} else if(s[pos] == 'C') {
		q1(s, pos+1)
	} else {
		println("Halted")
	}
}

func q1(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q1")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q1(s, pos+1)
	} else if(s[pos] == 'B') {
		q2(s, pos-1)
	} else {
		println("Halted")
	}
}

func q2(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q2")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		s = s[:pos]+"X"+s[pos+1:]
		q3(s, pos-1)
	} else if(s[pos] == 'C') {
		q7(s, pos+1)
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
	if(s[pos] == 'Y') {
		q4(s, pos-1)
	} else if(s[pos] == '0') {
		s = s[:pos]+"Y"+s[pos+1:]
		q5(s, pos+1)
	} else if(s[pos] == 'B') {
		q8(s, pos+1)
	} else {
		println("Halted")
	}
}

func q5(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q5")
	time.Sleep(time.Second)
	if(s[pos] == 'Y') {
		q5(s, pos+1)
	} else if(s[pos] == 'C') {
		q6(s, pos+1)
	} else {
		println("Halted")
	}
}

func q6(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q6")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q6(s, pos+1)
	} else if(s[pos] == 'X') {
		q2(s, pos-1)
	} else {
		println("Halted")
	}
}

func q7(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q7")
	time.Sleep(time.Second)
	if(s[pos] == 'X') {
		s = s[:pos]+"0"+s[pos+1:]
		q7(s, pos+1)
	} else if(s[pos] == 'B') {
		q2(s, pos-1)
	} else {
		println("Halted")
	}
}

func q8(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q8")
	time.Sleep(time.Second)
	if(s[pos] == 'Y') {
		s = s[:pos]+"B"+s[pos+1:]
		q8(s, pos+1)
	} else if(s[pos] == 'C') {
		s = s[:pos]+"B"+s[pos+1:]
		q9(s, pos+1)
	} else {
		println("Halted")
	}
}

func q9(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q9")
	time.Sleep(time.Second)
	if(s[pos] == 'X') {
		s = s[:pos]+"0"+s[pos+1:]
		q9(s, pos+1)
	} else if(s[pos] == '0') {
		s = s[:pos]+"B"+s[pos+1:]
		q9(s, pos+1)
	} else if(s[pos] == 'B') {
		q10(s, pos-1)
	} else {
		println("Halted")
	}
}

func q10(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q10")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		s = s[:pos]+"B"+s[pos+1:]
		q11(s, pos+1)
	} else {
		println("Halted")
	}
}

func q11(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q11")
	println("Finished, result =", strings.Count(s, "0"))
}