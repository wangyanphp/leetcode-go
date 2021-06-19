package main

func permutation(s string) []string {
	res = make([]string, 0)
	b := []byte(s)
	rec(b, make([]byte, 0))
	return res
}

var res []string
func rec(s []byte, pre []byte) {

	if len(s) == 0 {
		res = append(res, string(pre))
		return
	}

	set := make(map[uint8]struct{})

	for i := 0; i < len(s); i++ {

		if _, ok := set[s[i]]; ok {
			continue
		}
		s[0], s[i] = s[i], s[0]
		rec(s[1:], append(pre, s[0]))
		s[0], s[i] = s[i], s[0]
		set[s[i]] = struct{}{}

	}
}

func main()  {

	permutation("abc")

}
