package main

// 这里的问题是 可以采用高端的方法
func hammingWeight(num uint32) int {

	//ret := 0
	//for; num > 0; {
	//	if num & 1 == 1 {
	//		ret += 1
	//	}
	//	num >>= 1
	//}
	//return ret

	ret := 0
	for ; num > 0; {
		ret +=1
		num = num &(num-1)
	}
	return ret
}
