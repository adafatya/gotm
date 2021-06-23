package fac

import (
	"time"
	"strings"
)

func Fac(n int) {
	var s string = "B"
	if(n > 0) {
		s += "X"
	} else if (n < 0) {
		s += "Y"
		n *= -1
	}
	for i := 0; i < n; i++ {
		s += "0"
	}
	s += "1B"
	q0(s, 1)
}

func q0(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q0")
	time.Sleep(time.Second)
	if(s[pos] == 'X') {
		s = s[:pos]+"B"+s[pos+1:]
		q1(s, pos+1)
	} else if(s[pos] == 'Y') {
		s = s[:pos]+"B"+s[pos+1:]
		q30(s, pos+1)
	} else if(s[pos] =='1') {
		s = s[:pos]+"Y"+s[pos+1:]
		q31(s, pos+1)
	} else {
		println("Halted")
	}
}

func q1(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q1")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q1(s, pos+1)
	} else if(s[pos] == '1') {
		q2(s, pos+1)
	} else {
		println("Halted")
	}
}

func q2(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q2")
	time.Sleep(time.Second)
	if(s[pos] == 'B') {
		s = s[:pos]+"XB"
		q3(s, pos-1)
	} else {
		println("Halted")
	}
}

func q3(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q3")
	time.Sleep(time.Second)
	if(s[pos] == '1') {
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
		q5(s, pos+1)
	} else {
		println("Halted")
	}
}

func q5(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q5")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		s = s[:pos]+"Z"+s[pos+1:]
		q6(s, pos+1)
	} else if(s[pos] == '1') {
		q12(s, pos-1)
	} else {
		println("Halted")
	}
}

func q6(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q6")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q6(s, pos+1)
	} else if(s[pos] == '1') {
		q7(s, pos+1)
	} else {
		println("Halted")
	}
}

func q7(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q7")
	time.Sleep(time.Second)
	if(s[pos] == 'X') {
		q8(s, pos+1)
	} else {
		println("Halted")
	}
}

func q8(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q8")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q8(s, pos+1)
	} else if(s[pos] == 'B') {
		s = s[:pos]+"0B"
		q9(s, pos-1)
	} else {
		println("Halted")
	}
}

func q9(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q9")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q9(s, pos-1)
	} else if(s[pos] == 'X') {
		q10(s, pos-1)
	} else {
		println("Halted")
	}
}

func q10(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q10")
	time.Sleep(time.Second)
	if(s[pos] == '1') {
		q11(s, pos-1)
	} else {
		println("Halted")
	}
}

func q11(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q11")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q11(s, pos-1)
	} else if(s[pos] == 'Z') {
		q5(s, pos+1)
	} else {
		println("Halted")
	}
}

func q12(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q12")
	time.Sleep(time.Second)
	if(s[pos] == 'Z') {
		s = s[:pos]+"0"+s[pos+1:]
		q12(s, pos-1)
	} else if(s[pos] == 'B') {
		q13(s, pos+1)
	} else {
		println("Halted")
	}
}

func q13(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q13")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		s = s[:pos]+"B"+s[pos+1:]
		q14(s, pos+1)
	} else {
		println("Halted")
	}
}

func q14(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q14")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		s = s[:pos]+"Z"+s[pos+1:]
		q15(s, pos+1)
	} else if(s[pos] == '1') {
		s = s[:pos]+"B"+s[pos+1:]
		q28(s, pos+1)
	} else {
		println("Halted")
	}
}

func q15(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q15")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		s = s[:pos]+"Z"+s[pos+1:]
		q16(s, pos+1)
	} else if(s[pos] == '1') {
		q24(s, pos+1)
	} else {
		println("Halted")
	}
}

func q16(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q16")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q16(s, pos+1)
	} else if(s[pos] == '1') {
		q17(s, pos+1)
	} else {
		println("Halted")
	}
}

func q17(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q17")
	time.Sleep(time.Second)
	if(s[pos] == 'X') {
		q18(s, pos+1)
	} else {
		println("Halted")
	}
}

func q18(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q0")
	time.Sleep(time.Second)
	if(s[pos] == '0' || s[pos] == 'Z') {
		q18(s, pos+1)
	} else if(s[pos] == 'B') {
		q19(s, pos-1)
	} else {
		println("Halted")
	}
}

func q19(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q19")
	time.Sleep(time.Second)
	if(s[pos] == 'Z') {
		q19(s, pos-1)
	} else if(s[pos] == '0') {
		s = s[:pos]+"Z"+s[pos+1:]
		q20(s, pos+1)
	} else if (s[pos] == 'X') {
		q22(s, pos-1)
	} else {
		println("Halted")
	}
}

func q20(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q20")
	time.Sleep(time.Second)
	if(s[pos] == '0' || s[pos] == 'Z') {
		q20(s, pos+1)
	} else if(s[pos] == 'B') {
		s = s[:pos]+"0B"
		q21(s, pos-1)
	} else {
		println("Halted")
	}
}

func q21(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q21")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q21(s, pos-1)
	} else if(s[pos] == 'Z') {
		q19(s, pos-1)
	} else {
		println("Halted")
	}
}

func q22(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q22")
	time.Sleep(time.Second)
	if(s[pos] == '1') {
		q23(s, pos+1)
	} else {
		println("Halted")
	}
}

func q23(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q23")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q23(s, pos-1)
	} else if(s[pos] == 'Z') {
		q15(s, pos+1)
	} else if(s[pos] == 'X') {
		q24(s, pos+1)
	} else {
		println("Halted")
	}
}

func q24(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q24")
	time.Sleep(time.Second)
	if(s[pos] == 'X') {
		q25(s, pos+1)
	} else {
		println("Halted")
	}
}

func q25(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q25")
	time.Sleep(time.Second)
	if(s[pos] == 'Z') {
		s = s[:pos]+"0"+s[pos+1:]
		q25(s, pos+1)
	} else if(s[pos] == '0') {
		q26(s, pos-1)
	} else {
		println("Halted")
	}
}

func q26(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q26")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q26(s, pos-1)
	} else if(s[pos] == 'X') {
		q27(s, pos-1)
	} else {
		println("Halted")
	}
}

func q27(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q27")
	time.Sleep(time.Second)
	if(s[pos] == '1') {
		q12(s, pos-1)
	} else {
		println("Halted")
	}
}

func q28(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q28")
	time.Sleep(time.Second)
	if(s[pos] == 'X') {
		q29(s, pos+1)
	} else {
		println("Halted")
	}
}

func q29(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q29")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q29(s, pos+1)
	} else if(s[pos] == 'B') {
		q33(s, pos+1)
	} else {
		println("Halted")
	}
}

func q30(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q30")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		s = s[:pos]+"B"+s[pos+1:]
		q30(s, pos+1)
	} else if(s[pos] == '1') {
		q31(s, pos+1)
	} else {
		println("Halted")
	}
}

func q31(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q31")
	time.Sleep(time.Second)
	if(s[pos] == 'B') {
		s = s[:pos]+"XB"
		q32(s, pos+1)
	} else {
		println("Halted")
	}
}

func q32(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q32")
	time.Sleep(time.Second)
	if(s[pos] == 'B') {
		q33(s, pos+1)
	} else {
		println("Halted")
	}
}

func q33(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q33")
	println("Finished, result =", strings.Count(s, "0"))
}