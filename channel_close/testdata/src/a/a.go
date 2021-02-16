package a

func f() {
	// The pattern can be written in regular expression.
	var gopher int // want "pattern"
	print(gopher)  // want "identifier is gopher"

	ch := make(chan int)
	close(ch)
	close(ch)

	ch2 := make(chan int)
	close(ch2)

	myclose := func(ch chan int) {
		close(ch)
	}
}
