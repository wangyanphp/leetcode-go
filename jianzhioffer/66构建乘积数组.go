package main

func constructArr(a []int) []int {

	if len(a) == 0 {
		return nil
	}

	b := make([]int, len(a))
	b[0] = 1
	for i := 1; i< len(a); i++ {
		b[i] = b[i-1]*a[i-1]
	}

	var tmp = 1
	for i := len(a)-1; i>=0; i-- {
		b[i] *= tmp
		tmp *= a[i]
	}

	return b


}
