package main


func getLeastNumbers(arr []int, k int) []int {

	if len(arr) <= k {
		return arr
	}
	quicksort(arr, 0, len(arr)-1, k)
	return arr[:k]
}

func quicksort(arr []int, lo, hi, k int) {
	mid := partion(arr, lo, hi)
	if mid == k-1 || mid == k {
		return
	}

	if mid > k-1 || mid-1 > k-1 {
		quicksort(arr, lo, mid-1, k)
	}

	if mid < k-1 {
		quicksort(arr, mid+1, hi, k)
	}
}

// [lo, hi]
func partion(arr []int, lo, hi int) int  {
	if lo >= hi {
		return lo
	}

	pivot := lo
	index := pivot+1

	for i := index; i<=hi; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	// 此时 [pivot+1, index) 都小于 arr[pivot], [index, hi]大于arr[pivot]
	arr[pivot], arr[index-1] = arr[index-1], arr[pivot]
	return index-1
}
