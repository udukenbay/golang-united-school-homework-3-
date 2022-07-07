package homework

import (
	"fmt"
	"sort"
)

func main_map() {
	a := map[int]string{2: "a", 0: "b", 1: "c"}
	fmt.Println(sortMapValues(a))

	b := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	fmt.Println(sortMapValues(b))
}

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0, len(input))

	for k := range input {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		result = append(result, input[k])
	}
	return result
}
