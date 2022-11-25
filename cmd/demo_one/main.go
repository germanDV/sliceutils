package main

import (
	"fmt"
	"math"

	su "gitlab.com/germandv/sliceutils"
)

func main() {
	input := []int{99, 42, -34, 0, -23, 45, 9, 5, 7}
	display(input, "Input")

	negatives := su.Filter(input, isNegative)
	display(negatives, "Negatives")

	squares := su.Map(input, square)
	display(squares, "Squares")

	nine, _ := su.Find(input, isNine)
	display([]int{nine}, "Find")

	max := su.Reduce(input, findMax, 0)
	display([]int{max}, "Max")

	min := su.Reduce(input, findMin, math.MaxInt)
	display([]int{min}, "Min")

	hasOdd := su.Some(input, isOdd)
	fmt.Printf("Does it include at least one odd number? %t\n", hasOdd)

	allNegative := su.Every(input, isNegative)
	fmt.Printf("Does it consist entirely of negative numbres? %t\n", allNegative)
}

func display(arr []int, title string) {
	fmt.Printf("--- %s\n", title)
	fmt.Printf("%v\n", arr)
	fmt.Println()
}

func isNegative(n int) bool {
	return n < 0
}

func square(n int) int {
	return int(math.Pow(float64(n), 2))
}

func numberMatcher(n int) func(m int) bool {
	return func(m int) bool {
		return m == n
	}
}

var isNine = numberMatcher(9)

func findMax(prev, curr int) int {
	if curr > prev {
		return curr
	}
	return prev
}

func findMin(prev, curr int) int {
	if curr < prev {
		return curr
	}
	return prev
}

func isOdd(n int) bool {
	return n%2 != 0
}
