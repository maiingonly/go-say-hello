package go_say_hello

func SayHello() string {
	return "hello maiing"
}

type SumCalc interface {
}

type Value struct {
	a, b, c int
}

func (n Value) multiplication() int {
	return n.a * n.b * n.c
}
