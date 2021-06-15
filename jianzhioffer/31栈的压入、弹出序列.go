package main

// 这里的一个关键点：数字各不相同
// 假设我们用一个栈来模拟出栈入栈，那么最终应该为空
func validateStackSequences(pushed []int, popped []int) bool {

	stack := make([]int, 0)


	for i, j := 0, 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for ;len(stack) > 0 && stack[len(stack)-1] == popped[j]; {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return len(stack) == 0

}


