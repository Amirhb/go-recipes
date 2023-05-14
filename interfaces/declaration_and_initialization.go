package main

type printer interface {
	Print()
}

func main() {
	var printable = printer{} // error! invalid composite literal type printer
}
