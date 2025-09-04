package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	slice := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	ans := make(map[int][]float64)
	sort.Float64s(slice)

	cur := -1

	for _, el := range slice {
		if cur == -1 || (math.Abs(float64(cur))-math.Abs(el)) <= 10 {
			cur = (int(el) / 10) * 10
		}
		ans[cur] = append(ans[cur], el)
	}

	for key, array := range ans {
		fmt.Printf("%v: %v\n", key, array)
	}
}
