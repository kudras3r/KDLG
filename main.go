package main

import "fmt"

func main() {
	s := "asd\ni64\n asd 12\n ] "
	l := New(s)
	for {
		// fmt.Println(l.lastT)
		t := l.NextToken()
		// l.lastT = t

		if t.cat == EOF {
			break
		}
		fmt.Println(fmt.Sprintf("cat:%d	 ln:%d 	ival:%d	fval:%f 	sval:%s	bval:%t", t.cat, t.ln, t.ival, t.fval, t.sval, t.bval))
	}
}
