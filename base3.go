package base3

import (
	"fmt"
	"strconv"
)

func toBase3(n int, digits int) string {
	if n == 0 {
		return fmt.Sprintf("%0*d", digits, 0)
	}
	var result []int
	for n > 0 {
		result = append(result, n%3)
		n = n / 3
	}

	// Convert result slice to string in reverse order and pad with leading zeros
	base3 := ""
	for i := len(result) - 1; i >= 0; i-- {
		base3 += strconv.Itoa(result[i])
	}
	if len(base3) < digits {
		base3 = fmt.Sprintf("%0*s", digits, base3)
	}
	return base3
}
