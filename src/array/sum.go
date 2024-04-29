package main

func Sum(numbers []int) int {
	ans := 0
	for _, v := range numbers {
		ans += v
	}
	return ans
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return
}
