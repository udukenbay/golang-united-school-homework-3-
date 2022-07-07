package homework

import "fmt"

func main_slice() {
	a := []int64{1, 2, 5, 15}
	fmt.Println(reverse(a))
}

func reverse(input []int64) (result []int64) {
	for i := len(input) - 1; i >= 0; i-- {
		result = append(result, input[i])
	}

	return result
}
