package main

func singleNumbers(nums []int) []int {

	a, b := 0, 0
	x := 0
	for _, num := range nums{
		x ^= num
	}

	e := 1
	for;x&e == 0; {
		e <<=1
	}

	for _, num := range nums {
		if num & e == 0 {
			a ^= num
		}else {
			b ^= num
		}
	}
	return []int{a, b}
}


func singleNumber(nums []int) int {

	res := 0
	for i := 0; i<32; i++ {
		count:=0
		bit:=1<<i
		for _,num:=range nums{
			if num&bit!=0{
				count++
			}
		}
		if count%3!=0{
			res|=bit
		}
	}
	return res

}
