package main

var temp []int


func reversePairs(nums []int) int {
	temp = make([]int, len(nums))
	return mergeSort(nums, 0, len(nums)-1)

}


func mergeSort(arr []int, lo, hi int) int {
	length := hi-lo+1
	if length < 2 {
		return 0
	}
	middle := lo+(length / 2)-1
	left := mergeSort(arr,lo, middle)
	right := mergeSort(arr, middle+1, hi)
	cross := merge(arr, lo, middle, hi)
	return left+right+cross
}

func merge(arr []int, lo, mid, hi int) int {

	res := 0

	i, j := lo, mid+1
	for k := lo; k<=hi; k++ {
		if i > mid {
			temp[k] = arr[j]
			j++
		}else if j >  hi {
			temp[k] = arr[i]
			i++
		}else if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		}else {
			temp[k] = arr[j]
			j++
			res += mid-i+1
		}
	}

	for k := lo; k<= hi; k++ {
		arr[k] = temp[k]
	}
	return res
}