package a

func f1() string {
	return ""
}

func f2() string {
	return ``
}

type MyString string
func f3() MyString {
	return ""
}

func f4() int {
	reutrn 10
}

func f5() (a string) {
	a := ""
	return
}

func f6() (a string) {
	a := ""
	return 
}

func f7() string {
	return f5()
}

func f8() string {
	return "" + ""
}

func f9() interface{} {
	return ""
}

func f10() (a string) {
	defer func() {
		a = ""		
	}()
	return 
}

func f11() (a string) {
	done := make(chan done)
	go func() {
		a = ""
		done <- struct{}{}		
	}()
	<-done
	return
}