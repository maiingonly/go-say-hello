package go_say_hello

func SayHello() string {
	return "hello maiing"
}

type SumCalc interface {
	multication() int
}

type Value struct {
	A, B, C int
}

func (n Value) Multiplication() int {
	return n.A * n.B * n.C
}
