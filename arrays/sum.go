package arrays


// Sum calculates the total from an array of numbers.
func Sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum+=v
	}
	return sum
}