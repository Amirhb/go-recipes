package main

func main() {
	var primeNumbers = []int{2, 3, 5, 7, 11, 13, 17, 19}
	var numbers = [5]int{1, 2, 3, 4, 5}
	var gameList = []string{"soccer", "hockey", "basketball"}

	type employee  struct {
		FirstName string
		LastName  string
	}

	var employees = []employee{
		employee{"Anton", "Pavlov"},
		employee{"Elena", "Kirienko"},
	}
}
