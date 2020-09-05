package lc

import "strconv"

func calPoints(ops []string) int {
	stack := []int{}
	sum, points := 0, 0
	for _, v := range ops {
		switch v {
		case "C":
			sum -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			continue
		case "D":
			points = stack[len(stack)-1] * 2
			sum += points
		case "+":
			points = stack[len(stack)-1] + stack[len(stack)-2]
			sum += points
		default:
			points, _ = strconv.Atoi(v)
			sum += points
		}
		stack = append(stack, points)
	}

	return sum
}
