package go_say_hello

func SayHello(name string) string {
	return "hello " + name
}

type SumCalc interface {
	Multication() int
}

type Value struct {
	A, B, C int
}

func (n Value) Multiplication() int {
	return n.A * n.B * n.C
}
