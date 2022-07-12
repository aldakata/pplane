package main

func Min(is ...int) int {
	min := is[0]
	for _, i := range is[1:] {
		if i < min {
			min = i
		}
	}
	return min
}

func Max(is ...int) int {
	max := is[0]
	for _, i := range is[1:] {
		if i > max {
			max = i
		}
	}
	return max
}
