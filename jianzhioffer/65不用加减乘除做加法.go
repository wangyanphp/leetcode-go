package main


func add(a int, b int) int {

	var c int
	for b != 0 {
		c = (a&b) << 1
		a ^= b
		b = c
	}

	return a

}
