package decodeString

func decodeString(s string) string {
	resStack := make([]string, 0)
	multiStack := make([]int, 0)

	curStr, curMulti := "", 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			curMulti = curMulti * 10 + int(s[i] - '0')
		} else if s[i] == '[' {
			multiStack = append(multiStack, curMulti)
			resStack = append(resStack, curStr)
			curMulti = 0
			curStr = ""
		} else if s[i] == ']' {
			multi := multiStack[len(multiStack) - 1]
			multiStack = multiStack[:len(multiStack) - 1]
			cur := ""
			for j := 0; j < multi; j++ {
				cur += curStr
			}
			preStr := resStack[len(resStack) - 1]
			resStack = resStack[:len(resStack) - 1]
			curStr = preStr + cur
		} else {
			curStr += string(s[i])
		}
	}
	return curStr
}