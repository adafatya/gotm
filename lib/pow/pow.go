package pow

import (
	"time"
	"strings"
)

func Pow(n int, m int) {
	var s string = "B"
	if(n > 0) {
		s += "X"
	} else if(n < 0) {
		s += "Y"
		n *= -1
	}
	for i := 0; i < n; i++ {
		s += "0"
	}
	s += "1"
	if(m > 0) {
		s += "X"
	} else if(m < 0) {
		s += "Y"
		m *= -1
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
		s = s[:pos]+"B"+s[pos+1:]
		q1(s, pos+1)
	} else if(s[pos] == 'Y') {
		s = s[:pos]+"B"+s[pos+1:]
		q36(s, pos+1)
	} else if(s[pos] == '1') {
		s = s[:pos]+"B"+s[pos+1:]
		q37(s, pos+1)
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
	if(s[pos] == 'X') {
		s = s[:pos]+"B"+s[pos+1:]
		q3(s, pos+1)
	} else if(s[pos] == 'Y') {
		s = s[:pos]+"B"+s[pos+1:]
		q41(s, pos-1)
	} else if(s[pos] == 'B') {
		q39(s, pos-1)
	} else {
		println("Halted")
	}
}

func q3(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q3")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q3(s, pos+1)
	} else if(s[pos] == 'B') {
		s = s[:pos]+"XB"
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
		q5(s, pos-1)
	} else {
		println("Halted")
	}
}

func q5(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q5")
	time.Sleep(time.Second)
	if(s[pos] == '1') {
		q6(s, pos-1)
	} else {
		println("Halted")
	}
}

func q6(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q6")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q6(s, pos-1)
	} else if(s[pos] == 'B') {
		q7(s, pos+1)
	} else {
		println("Halted")
	}
}

func q7(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q7")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q7(s, pos+1)
	} else if(s[pos] == '1') {
		q8(s, pos+1)
	} else {
		println("Halted")
	}
}

func q8(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q8")
	time.Sleep(time.Second)
	if(s[pos] == 'B') {
		q8(s, pos+1)
	} else if(s[pos] == '0') {
		s = s[:pos]+"B"+s[pos+1:]
		q9(s, pos-1)
	} else {
		println("Halted")
	}
}

func q9(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q9")
	time.Sleep(time.Second)
	if(s[pos] == 'B') {
		q9(s, pos-1)
	} else if(s[pos] == '1') {
		q10(s, pos-1)
	} else {
		println("Halted")
	}
}

func q10(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q10")
	time.Sleep(time.Second)
	if(s[pos] == 'Z') {
		q10(s, pos-1)
	} else if(s[pos] == '0') {
		s = s[:pos]+"Z"+s[pos+1:]
		q11(s, pos+1)
	} else if(s[pos] == 'B') {
		q16(s, pos+1)
	} else {
		println("Halted")
	}
}

func q11(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q11")
	time.Sleep(time.Second)
	if(s[pos] == 'Z') {
		q11(s, pos+1)
	} else if(s[pos] == '1') {
		q12(s, pos+1)
	} else {
		println("Halted")
	}
}

func q12(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q12")
	time.Sleep(time.Second)
	if(s[pos] == '0' || s[pos] == 'B') {
		q12(s, pos+1)
	} else if(s[pos] == 'X') {
		q13(s, pos+1)
	} else {
		println("Halted")
	}
}

func q13(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q13")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q13(s, pos+1)
	} else if(s[pos] == 'B') {
		s = s[:pos]+"0B"
		q14(s, pos-1)
	} else {
		println("Halted")
	}
}

func q14(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q14")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q14(s, pos-1)
	} else if(s[pos] == 'X') {
		q15(s, pos-1)
	} else {
		println("Halted")
	}
}

func q15(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q15")
	time.Sleep(time.Second)
	if(s[pos] == '0' || s[pos] == 'B') {
		q15(s, pos-1)
	} else if(s[pos] == '1') {
		q10(s, pos-1)
	} else {
		println("Halted")
	}
}

func q16(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q16")
	time.Sleep(time.Second)
	if(s[pos] == 'Z') {
		s = s[:pos]+"0"+s[pos+1:]
		q16(s, pos+1)
	} else if(s[pos] == '1') {
		q17(s, pos-1)
	} else {
		println("Halted")
	}
}

func q17(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q17")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		s = s[:pos]+"Z"+s[pos+1:]
		q18(s, pos+1)
	} else {
		println("Halted")
	}
}

func q18(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q18")
	time.Sleep(time.Second)
	if(s[pos] == '1') {
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
	} else if(s[pos] == '0') {
		s = s[:pos]+"B"+s[pos+1:]
		q20(s, pos-1)
	} else if(s[pos] == 'X') {
		q32(s, pos-1)
	} else {
		println("Halted")
	}
}

func q20(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q20")
	time.Sleep(time.Second)
	if(s[pos] == '0' || s[pos] == 'B') {
		q20(s, pos-1)
	} else if(s[pos] == '1') {
		q21(s, pos-1)
	} else {
		println("Halted")
	}
}

func q21(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q21")
	time.Sleep(time.Second)
	if(s[pos] == 'Z') {
		q21(s, pos-1)
	} else if(s[pos] == '0') {
		s = s[:pos]+"Z"+s[pos+1:]
		q22(s, pos+1)
	} else if(s[pos] == 'B') {
		q28(s, pos+1)
	} else {
		println("Halted")
	}
}

func q22(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q22")
	time.Sleep(time.Second)
	if(s[pos] == 'Z') {
		q22(s, pos+1)
	} else if(s[pos] == '1') {
		q23(s, pos+1)
	} else {
		println("Halted")
	}
}

func q23(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q23")
	time.Sleep(time.Second)
	if(s[pos] == '0' || s[pos] == 'B') {
		q23(s, pos+1)
	} else if(s[pos] == 'X') {
		q24(s, pos+1)
	} else {
		println("Halted")
	}
}

func q24(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q24")
	time.Sleep(time.Second)
	if(s[pos] == '0' || s[pos] == 'Z') {
		q24(s, pos+1)
	} else if(s[pos] == 'B') {
		q25(s, pos-1)
	} else {
		println("Halted")
	}
}

func q25(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q25")
	time.Sleep(time.Second)
	if(s[pos] == 'Z') {
		q25(s, pos-1)
	} else if(s[pos] == '0') {
		s = s[:pos]+"Z"+s[pos+1:]
		q26(s, pos+1)
	} else if(s[pos] == 'X') {
		q20(s, pos-1)
	} else {
		println("Halted")
	}
}

func q26(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q26")
	time.Sleep(time.Second)
	if(s[pos] == '0' || s[pos] == 'Z') {
		q26(s, pos+1)
	} else if(s[pos] == 'B') {
		s = s[:pos]+"0B"
		q27(s, pos-1)
	} else {
		println("Halted")
	}
}

func q27(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q27")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q27(s, pos-1)
	} else if(s[pos] == 'Z') {
		q25(s, pos-1)
	} else {
		println("Halted")
	}
}

func q28(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q28")
	time.Sleep(time.Second)
	if(s[pos] == 'Z') {
		s = s[:pos]+"0"+s[pos+1:]
		q28(s, pos+1)
	} else if(s[pos] == '1') {
		q29(s, pos+1)
	} else {
		println("Halted")
	}
}

func q29(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q29")
	time.Sleep(time.Second)
	if(s[pos] == '0' || s[pos] == 'B') {
		q29(s, pos+1)
	} else if(s[pos] == 'X') {
		q30(s, pos+1)
	} else {
		println("Halted")
	}
}

func q30(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q30")
	time.Sleep(time.Second)
	if(s[pos] == 'Z') {
		s = s[:pos]+"0"+s[pos+1:]
		q30(s, pos+1)
	} else if(s[pos] == '0') {
		q30(s, pos-1)
	} else if(s[pos] == 'X') {
		q31(s, pos-1)
	} else {
		println("Halted")
	}
}

func q31(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q31")
	time.Sleep(time.Second)
	if(s[pos] == '0' || s[pos] == 'B') {
		q31(s, pos-1)
	} else if(s[pos] == '1') {
		q17(s, pos-1)
	} else {
		println("Halted")
	}
}

func q32(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q32")
	time.Sleep(time.Second)
	if(s[pos] == 'B') {
		q32(s, pos-1)
	} else if(s[pos] == '1') {
		s = s[:pos]+"B"+s[pos+1:]
		q33(s, pos-1)
	} else {
		println("Halted")
	}
}

func q33(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q33")
	time.Sleep(time.Second)
	if(s[pos] == 'Z') {
		s = s[:pos]+"B"+s[pos+1:]
		q34(s, pos-1)
	} else {
		println("Halted")
	}
}

func q34(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q34")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		s = s[:pos]+"B"+s[pos+1:]
		q34(s, pos-1)
	} else if(s[pos] == 'B') {
		q34(s, pos+1)
	} else if(s[pos] == 'X') {
		q35(s, pos+1)
	} else {
		println("Halted")
	}
}

func q35(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q35")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		q35(s, pos+1)
	} else if(s[pos] == 'B') {
		q47(s, pos+1)
	} else {
		println("Halted")
	}
}

func q36(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q36")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		s = s[:pos]+"B"+s[pos+1:]
		q36(s, pos+1)
	} else if(s[pos] == '1') {
		s = s[:pos]+"B"+s[pos+1:]
		q37(s, pos+1)
	} else {
		println("Halted")
	}
}

func q37(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q37")
	time.Sleep(time.Second)
	if(s[pos] == 'X' || s[pos] == 'Y' || s[pos] == 'B') {
		s = s[:pos]+"B"+s[pos+1:]
		q38(s, pos+1)
	} else {
		println("Halted")
	}
}

func q38(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q38")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		s = s[:pos]+"B"+s[pos+1:]
		q38(s, pos+1)
	} else if(s[pos] == 'B') {
		s = s[:pos]+"Y"+s[pos+1:]
		q45(s, pos+1)
	} else {
		println("Halted")
	}
}

func q39(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q39")
	time.Sleep(time.Second)
	if(s[pos] == '1') {
		s = s[:pos]+"B"+s[pos+1:]
		q40(s, pos-1)
	} else {
		println("Halted")
	}
}

func q40(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q40")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		s = s[:pos]+"B"+s[pos+1:]
		q40(s, pos-1)
	} else if(s[pos] == 'B') {
		s = s[:pos]+"Y"+s[pos+1:]
		q45(s, pos+1)
	} else {
		println("Halted")
	}
}

func q41(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q41")
	time.Sleep(time.Second)
	if(s[pos] == '1') {
		s = s[:pos]+"B"+s[pos+1:]
		q42(s, pos-1)
	} else {
		println("Halted")
	}
}

func q42(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q42")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		s = s[:pos]+"B"+s[pos+1:]
		q42(s, pos-1)
	} else if(s[pos] == 'B') {
		q43(s, pos+1)
	} else {
		println("Halted")
	}
}

func q43(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q43")
	time.Sleep(time.Second)
	if(s[pos] == 'B') {
		q43(s, pos+1)
	} else if(s[pos] == '0') {
		s = s[:pos]+"B"+s[pos+1:]
		q44(s, pos+1)
	} else {
		println("Halted")
	}
}

func q44(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q44")
	time.Sleep(time.Second)
	if(s[pos] == '0') {
		s = s[:pos]+"B"+s[pos+1:]
		q44(s, pos+1)
	} else if(s[pos] == 'B') {
		s = s[:pos]+"Y"+s[pos+1:]
		q1(s, pos+1)
	} else {
		println("Halted")
	}
}

func q45(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q45")
	time.Sleep(time.Second)
	if(s[pos] == 'B') {
		s = s[:pos]+"X"+s[pos+1:]
		q46(s, pos+1)
	} else {
		println("Halted")
	}
}

func q46(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q46")
	time.Sleep(time.Second)
	if(s[pos] == 'B') {
		q47(s, pos+1)
	} else {
		println("Halted")
	}
}

func q47(s string, pos int) {
	println(s[:pos], ">", s[pos:], "\tstate = q11")
	println("Finished, result =", strings.Count(s, "0"))
}