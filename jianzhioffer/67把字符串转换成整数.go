package main

import "fmt"

const (
	INT_MAX = 1 << 31 -1
	INT_MIN = - (1<<31)
)
func strToInt(str string) int {


	bn := 214748364

	start := 0

	negFlag := false
	for i := 0; i< len(str); i++ {
		if str[i] == ' ' {
			start++
			continue
		}
		break
	}

	if len(str) <= start {
		return 0
	}

	if str[start] == '-' {
		negFlag = true
	}

	if str[start] == '-' || str[start] == '+' {
		start++
	}

	ret := 0
	for ;start < len(str); start++ {

		if str[start] < '0' || str[start] > '9' {
			break
		}
		tmp := int(str[start] - '0')

		if ret > bn || (ret == bn && tmp > 7) {
			if negFlag {
				return INT_MIN
			}else {
				return INT_MAX
			}
		}

		ret = ret*10 + tmp
	}

	if negFlag {
		ret = -ret
	}

	return ret

}


func main() {

	fmt.Println(strToInt("2147483646"))
}
