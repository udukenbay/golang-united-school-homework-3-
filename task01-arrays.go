package homework

import (
	"fmt"
)

func main() {
	a := [15]float32{1, 2, 3, 4, 5, 6}
	fmt.Println(average(a))
}

func average(input [15]float32) (result float32) {
	for i := 0; i < len(input); i++ {
		result += input[i]
	}

	return result / float32(len(input))
}
