package go_say_hello

func SayHello() string {
	return "hello maiing"
}

type SumCalc interface {
	multication() int
}

type Value struct {
	a, b, c int
}

func (n Value) Multiplication() int {
	return n.a * n.b * n.c
}
