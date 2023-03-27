package main

func throwPanic(cars []string) {
	if len(cars) == 0 {
		panic("No cars for sale")
	}
}

func main() {
	throwPanic([]string{})
}
