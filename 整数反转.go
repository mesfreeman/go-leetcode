package main

import (
	"fmt"
	"math"
)

// @solution-sync:begin
func reverse(x int) int {
	res := 0
	for x != 0 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		res = res*10 + digit
	}
	return res
}

// @solution-sync:end

func main() {
	var x = 123

	var result = reverse(x)
	fmt.Printf("%v\n", result)
}
