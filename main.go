package main

import "fmt"

func main() {
	s := "i64\nasd <<= *\n  123 123.3  true @@    asd \nkek"

	l := New(s)
	for {
		// fmt.Println(l.lastT)
		t := l.NxToken()
		// l.lastT = t

		if t.cat == EOF {
			break
		}
		// fmt.Println(fmt.Sprintf("cat:%d	 ln:%d 	ival:%d	fval:%f 	sval:%s	bval:%t", t.cat, t.ln, t.ival, t.fval, t.sval, t.bval))
		fmt.Println(fmt.Sprintf("cat:%d	ln:%d col:%d	val: %s", t.cat, t.ln, t.col, t.str))
	}
}
