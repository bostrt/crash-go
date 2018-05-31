package main

type thing struct {
	value string
}

func main() {
	x := thing{value: "some value"}
	doBadThings(&x)
}

func doBadThings(x *thing) thing {
	x = nil
	return *x // Crash
}
