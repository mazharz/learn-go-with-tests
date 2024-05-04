package aas

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(listOfListOfNumbers ...[]int) []int {
	var sums []int

	for _, numbers := range listOfListOfNumbers {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(listOfListOfNumbers ...[]int) []int {
	var sums []int

	for _, numbers := range listOfListOfNumbers {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
