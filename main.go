package main

func main() {
	s := "<= = . 8"
	l := New(s)
	for {
		t := l.NextToken()
		if t.cat == EOF {

			break
		}

	}
}
