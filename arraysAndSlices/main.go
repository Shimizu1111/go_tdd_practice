package main

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	// for i := 0; i < 3; i++ {
	// 	sum += numbers[i]
	// }
	// _はインデックス番号
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// fmt.Println("length = ", lengthOfNumbers)
	// sums := make([]int, lengthOfNumbers)
	// fmt.Println("sums = ", sums)
	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// 	fmt.Println("numbers = ", sums)
	// }
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
		fmt.Println("sums = ", sums)
	}
	fmt.Println("sums = ", sums)
	return sums
}
