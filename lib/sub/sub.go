package sub

import (
	"time"
	"strings"
)

func Sub(n int, m int) {
	var s string = "B"
	if(n > 0) {
		s += "X"
	}
	for i := 0; i < n; i++ {
		s += "0"
	}
	s += "1"
	if(m < 0) {
		s += "Y"
		m*= -1
	} else if(m > 0) {
		s += "X"
	}
	for i := 0; i < m; i++ {
		s += "0"
	}
	s += "B"
	q0(s, 1)
}

func q0(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q0")
	time.Sleep(time.Second)
	if(s[pos] == 'X') {
		q1(s, pos+1)
	} else if(s[pos] == '1') {
		s = s[:pos]+"B"+s[pos+1:]
		q12(s, pos+1)
	} else {
		println("Halted")
	}
}

func q1(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q1")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		s = s[:pos]+"B"+s[pos+1:]
		q2(s, pos+1)
	} else if(s[pos] == '1') {
		s = s[:pos]+"B"+s[pos+1:]
		q8(s, pos-1)
	} else {
		println("Halted")
	}
}

func q2(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q2")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q2(s, pos+1)
	} else if(s[pos] == '1') {
		q3(s, pos+1)
	} else {
		println("Halted")
	}
}

func q3(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q3")
	time.Sleep(time.Second)
	if(s[pos] == 'Y') {
		q4(s, pos+1)
	} else if(s[pos] == 'X') {
		q9(s, pos+1)
	} else if(s[pos] == 'B') {
		q9(s, pos-1)
	} else {
		println("Halted")
	}
}

func q4(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q4")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q4(s, pos+1)
	} else if(s[pos] == 'B') {
		s = s[:pos]+"0B"
		q5(s, pos-1)
	} else {
		println("Halted")
	}
}

func q5(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q5")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q5(s, pos-1)
	} else if(s[pos] == 'Y') {
		q6(s, pos-1)
	} else {
		println("Halted")
	}
}

func q6(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q6")
	time.Sleep(time.Second)
	if(s[pos] == '1') {
		q7(s, pos-1)
	} else {
		println("Halted")
	}
}

func q7(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q7")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q7(s, pos-1)
	} else if(s[pos] == 'B') {
		q1(s, pos+1)
	} else {
		println("Halted")
	}
}

func q8(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q8")
	time.Sleep(time.Second)
	if(s[pos] == 'B') {
		q8(s, pos-1)
	} else if(s[pos] == 'X') {
		s = s[:pos]+"B"+s[pos+1:]
		q13(s, pos+1)
	} else {
		println("Halted")
	}
}

func q9(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q9")
	time.Sleep(time.Second)
	if(s[pos] == 'B') {
		q10(s, pos-1)
	} else if(s[pos] == '0') {
		q9(s, pos+1)
	} else if(s[pos] == '1') {
		s = s[:pos]+"B"+s[pos+1:]
		q14(s, pos-1)
	} else {
		println("Halted")
	}
}

func q10(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q10")
	time.Sleep(time.Second)
	if(s[pos] == 'B') {
		q11(s, pos-1)
	} else if(s[pos] == 'X') {
		q16(s, pos-1)
	} else {
		println("Halted")
	}
}

func q11(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q11")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q11(s, pos-1)
	} else if(s[pos] == 'X') {
		q6(s, pos-1)
	} else {
		println("Halted")
	}
}

func q12(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q12")
	time.Sleep(time.Second)
	if(s[pos] == 'X') {
		s = s[:pos]+"Y"+s[pos+1:]
		q23(s, pos+1)
	} else {
		println("Halted")
	}
}

func q13(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q13")
	time.Sleep(time.Second)
	if(s[pos] == 'B') {
		q13(s, pos+1)
	} else if(s[pos] == 'X') {
		s = s[:pos]+"Y"+s[pos+1:]
		q22(s, pos+1)
	} else if(s[pos] == 'Y') {
		s = s[:pos]+"X"+s[pos+1:]
		q22(s, pos+1)
	} else {
		println("Halted")
	}
}

func q14(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q14")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q14(s, pos-1)
	} else if(s[pos] == 'B') {
		s = "B"+s[:pos]+"0"+s[pos+1:]
		q15(s, pos+1)
	} else {
		println("Halted")
	}
}

func q15(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q15")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q15(s, pos+1)
	} else if(s[pos] == 'B') {
		q24(s, pos+1)
	} else {
		println("Halted")
	}
}

func q16(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q16")
	time.Sleep(time.Second)
	if(s[pos] == '1') {
		q17(s, pos-1)
	} else {
		println("Halted")
	}
}

func q17(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q17")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q17(s, pos-1)
	} else if(s[pos] == 'X') {
		q18(s, pos+1)
	} else {
		println("Halted")
	}
}

func q18(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q18")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q18(s, pos+1)
	} else if(s[pos] == 'B') {
		s = s[:pos]+"0"+s[pos+1:]
		q19(s, pos+1)
	} else {
		println("Halted")
	}
}

func q19(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q19")
	time.Sleep(time.Second)
	if(s[pos] == 'B') {
		q19(s, pos+1)
	} else if(s[pos] == '1') {
		s = s[:pos]+"B"+s[pos+1:]
		q20(s, pos+1)
	} else if(s[pos] == '0') {
		s = s[:pos]+"B"+s[pos+1:]
		q17(s, pos-1)
	} else {
		println("Halted")
	}
}

func q20(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q20")
	time.Sleep(time.Second)
	if(s[pos] == 'X') {
		s = s[:pos]+"B"+s[pos+1:]
		q21(s, pos+1)
	} else {
		println("Halted")
	}
}

func q21(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q21")
	time.Sleep(time.Second)
	if(s[pos] == 'B') {
		q24(s, pos+1)
	} else {
		println("Halted")
	}
}

func q22(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q22")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q22(s, pos+1)
	} else if(s[pos] == 'B') {
		q24(s, pos+1)
	} else {
		println("Halted")
	}
}


func q23(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q23")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q23(s, pos+1)
	} else if(s[pos] == 'B') {
		q24(s, pos+1)
	} else {
		println("Halted")
	}
}

func q24(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q24")
	println("Finished, result =", strings.Count(s, "0"))
}