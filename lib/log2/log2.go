package log2

import (
	"time"
	"strings"
)

func Log2(n int) {
	var track1, track2 string = "B", "B"
	for i := 0; i < n; i++ {
		track1 += "0"
		track2 += "B"
	}
	track1 += "C00B"
	track2 += "BBBB"

	q0(track1, track2, 1, 1)
}

func q0(t1 string, t2 string, pos1 int, pos2 int) {
	println(t1[:pos1], ">", t1[pos1:], "\tstate = q0")
	println(t2[:pos2], ">", t2[pos2:])
	println()
	time.Sleep(time.Second)
	if(t1[pos1] == '0' && t2[pos2] == 'B') {
		q0(t1, t2, pos1+1, pos2)
	} else if(t1[pos1] == 'C' && t2[pos2] == 'B') {
		q1(t1, t2, pos1+1, pos2)
	} else {
		println("Halted")
	}
}

func q1(t1 string, t2 string, pos1 int, pos2 int) {
	println(t1[:pos1], ">", t1[pos1:], "\tstate = q1")
	println(t2[:pos2], ">", t2[pos2:])
	println()
	time.Sleep(time.Second)
	if(t1[pos1] == '0' && t2[pos2] == 'B') {
		q1(t1, t2, pos1+1, pos2)
	} else if(t1[pos1] == 'B' && t2[pos2] == 'B') {
		t1 = t1[:pos1]+"C"+t1[pos1+1:]
		q2(t1, t2, pos1-1, pos2)
	} else {
		println("Halted")
	}
}

func q2(t1 string, t2 string, pos1 int, pos2 int) {
	println(t1[:pos1], ">", t1[pos1:], "\tstate = q2")
	println(t2[:pos2], ">", t2[pos2:])
	println()
	time.Sleep(time.Second)
	if(t1[pos1] == 'C' && t2[pos2] == 'B') {
		q7(t1, t2, pos1+1, pos2)
	} else if(t1[pos1] == '0' && t2[pos2] == 'B') {
		t1 = t1[:pos1]+"X"+t1[pos1+1:]
		q3(t1, t2, pos1-1, pos2)
	} else {
		println("Halted")
	}
}

func q3(t1 string, t2 string, pos1 int, pos2 int) {
	println(t1[:pos1], ">", t1[pos1:], "\tstate = q3")
	println(t2[:pos2], ">", t2[pos2:])
	println()
	time.Sleep(time.Second)
	if(t1[pos1] == '0' && t2[pos2] == 'B') {
		q3(t1, t2, pos1-1, pos2)
	} else if(t1[pos1] == 'C' && t2[pos2] == 'B') {
		q4(t1, t2, pos1-1, pos2)
	} else {
		println("Halted")
	}
}

func q4(t1 string, t2 string, pos1 int, pos2 int) {
	println(t1[:pos1], ">", t1[pos1:], "\tstate = q4")
	println(t2[:pos2], ">", t2[pos2:])
	println()
	time.Sleep(time.Second)
	if(t1[pos1] == 'Y' && t2[pos2] == 'B') {
		q4(t1, t2, pos1-1, pos2)
	} else if(t1[pos1] == '0' && t2[pos2] == 'B') {
		t1 = t1[:pos1]+"Y"+t1[pos1+1:]
		q5(t1, t2, pos1+1, pos2)
	} else if(t1[pos1] == 'B' && t2[pos2] == 'B') {
		q8(t1, t2, pos1+1, pos2)
	} else {
		println("Halted")
	}
}

func q5(t1 string, t2 string, pos1 int, pos2 int) {
	println(t1[:pos1], ">", t1[pos1:], "\tstate = q5")
	println(t2[:pos2], ">", t2[pos2:])
	println()
	time.Sleep(time.Second)
	if(t1[pos1] == 'Y' && t2[pos2] == 'B') {
		q5(t1, t2, pos1+1, pos2)
	} else if(t1[pos1] == 'C' && t2[pos2] == 'B') {
		q6(t1, t2, pos1+1, pos2)
	} else {
		println("Halted")
	}
}

func q6(t1 string, t2 string, pos1 int, pos2 int) {
	println(t1[:pos1], ">", t1[pos1:], "\tstate = q6")
	println(t2[:pos2], ">", t2[pos2:])
	println()
	time.Sleep(time.Second)
	if(t1[pos1] == '0' && t2[pos2] == 'B') {
		q6(t1, t2, pos1+1, pos2)
	} else if(t1[pos1] == 'X' && t2[pos2] == 'B') {
		q2(t1, t2, pos1-1, pos2)
	} else {
		println("Halted")
	}
}

func q7(t1 string, t2 string, pos1 int, pos2 int) {
	println(t1[:pos1], ">", t1[pos1:], "\tstate = q7")
	println(t2[:pos2], ">", t2[pos2:])
	println()
	time.Sleep(time.Second)
	if(t1[pos1] == 'X' && t2[pos2] == 'B') {
		t1 = t1[:pos1]+"0"+t1[pos1+1:]
		q7(t1, t2, pos1+1, pos2)
	} else if(t1[pos1] == 'C' && t2[pos2] == 'B') {
		t2 = t2[:pos2]+"0"+t2[pos2+1:]
		q2(t1, t2, pos1-1, pos2+1)
	} else {
		println("Halted")
	}
}

func q8(t1 string, t2 string, pos1 int, pos2 int) {
	println(t1[:pos1], ">", t1[pos1:], "\tstate = q8")
	println(t2[:pos2], ">", t2[pos2:])
	println()
	time.Sleep(time.Second)
	if(t1[pos1] == 'C' && t2[pos2] == 'B') {
		q9(t1, t2, pos1, pos2-1)
	} else if(t1[pos1] == 'Y' && t2[pos2] == 'B') {
		t1 = t1[:pos1]+"B"+t1[pos1+1:]
		q8(t1, t2, pos1+1, pos2)
	} else {
		println("Halted")
	}
}

func q9(t1 string, t2 string, pos1 int, pos2 int) {
	println(t1[:pos1], ">", t1[pos1:], "\tstate = q9")
	println(t2[:pos2], ">", t2[pos2:])
	println()
	time.Sleep(time.Second)
	if(t1[pos1] == 'C' && t2[pos2] == 'B') {
		t1 = t1[:pos1]+"B"+t1[pos1+1:]
		q15(t1, t2, pos1+1, pos2+1)
	} else if(t1[pos1] == 'C' && t2[pos2] == '0') {
		q10(t1, t2, pos1+1, pos2+1)
	} else {
		println("Halted")
	}
}

func q10(t1 string, t2 string, pos1 int, pos2 int) {
	println(t1[:pos1], ">", t1[pos1:], "\tstate = q10")
	println(t2[:pos2], ">", t2[pos2:])
	println()
	time.Sleep(time.Second)
	if(t1[pos1] == '0' && t2[pos2] == 'B') {
		q10(t1, t2, pos1+1, pos2)
	} else if(t1[pos1] == 'X' && t2[pos2] == 'B') {
		t1 = t1[:pos1]+"0"+t1[pos1+1:]
		q10(t1, t2, pos1+1, pos2)
	} else if(t1[pos1] == 'C' && t2[pos2] == 'B') {
		t1 = t1[:pos1]+"B"+t1[pos1+1:]+"B"
		if(len(t2) < len(t1)) {
			t2 += "B"
		}
		q11(t1, t2, pos1+1, pos2)
	} else {
		println("Halted")
	}
}

func q11(t1 string, t2 string, pos1 int, pos2 int) {
	println(t1[:pos1], ">", t1[pos1:], "\tstate = q11")
	println(t2[:pos2], ">", t2[pos2:])
	println()
	time.Sleep(time.Second)
	if(t1[pos1] == '0' && t2[pos2] == 'B') {
		q11(t1, t2, pos1+1, pos2)
	} else if(t1[pos1] == 'B' && t2[pos2] == 'B') {
		t1 = t1[:pos1]+"0B"
		if(len(t2) < len(t1)) {
			t2 += "B"
		}
		q12(t1, t2, pos1-1, pos2)
	} else {
		println("Halted")
	}
}

func q12(t1 string, t2 string, pos1 int, pos2 int) {
	println(t1[:pos1], ">", t1[pos1:], "\tstate = q12")
	println(t2[:pos2], ">", t2[pos2:])
	println()
	time.Sleep(time.Second)
	if(t1[pos1] == '0' && t2[pos2] == 'B') {
		q12(t1, t2, pos1-1, pos2)
	} else if(t1[pos1] == 'B' && t2[pos2] == 'B') {
		q13(t1, t2, pos1-1, pos2)
	} else {
		println("Halted")
	}
}

func q13(t1 string, t2 string, pos1 int, pos2 int) {
	println(t1[:pos1], ">", t1[pos1:], "\tstate = q13")
	println(t2[:pos2], ">", t2[pos2:])
	println()
	time.Sleep(time.Second)
	if(t1[pos1] == '0' && t2[pos2] == 'B') {
		q13(t1, t2, pos1-1, pos2)
	} else if(t1[pos1] == 'C' && t2[pos2] == 'B') {
		q14(t1, t2, pos1-1, pos2-1)
	} else {
		println("Halted")
	}
}

func q14(t1 string, t2 string, pos1 int, pos2 int) {
	println(t1[:pos1], ">", t1[pos1:], "\tstate = q14")
	println(t2[:pos2], ">", t2[pos2:])
	println()
	time.Sleep(time.Second)
	if(t1[pos1] == 'B' && t2[pos2] == 'B') {
		q0(t1, t2, pos1+1, pos2+1)
	} else if(t1[pos1] == 'B' && t2[pos2] == '0') {
		t1 = t1[:pos1]+"0"+t1[pos1+1:]
		t2 = t2[:pos2]+"B"+t2[pos2+1:]
		q14(t1, t2, pos1-1, pos2-1)
	} else {
		println("Halted")
	}
}

func q15(t1 string, t2 string, pos1 int, pos2 int) {
	println(t1[:pos1], ">", t1[pos1:], "\tstate = q15")
	println(t2[:pos2], ">", t2[pos2:])
	println()
	time.Sleep(time.Second)
	if(t1[pos1] == '0' && t2[pos2] == 'B') {
		t1 = t1[:pos1]+"B"+t1[pos1+1:]
		q15(t1, t2, pos1+1, pos2)
	} else if(t1[pos1] == 'X' && t2[pos2] == 'B') {
		t1 = t1[:pos1]+"B"+t1[pos1+1:]
		q15(t1, t2, pos1+1, pos2)
	} else if(t1[pos1] == 'C' && t2[pos2] == 'B') {
		t1 = t1[:pos1]+"B"+t1[pos1+1:]
		q16(t1, t2, pos1+1, pos2)
	} else {
		println("Halted")
	}
}

func q16(t1 string, t2 string, pos1 int, pos2 int) {
	println(t1[:pos1], ">", t1[pos1:], "\tstate = q16")
	println(t2[:pos2], ">", t2[pos2:])
	println("Finished, result =", strings.Count(t1, "0"))
}