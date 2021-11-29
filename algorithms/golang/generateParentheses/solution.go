package generateParentheses

var res []string

func generateParenthesis(n int) []string {
	res = make([]string, 0)
	helper("", n, n)
	return res
}

func helper(prefix string, leftCnt int, rightCnt int) {
	if leftCnt == 0 && rightCnt == 0 {
		res = append(res, prefix)
		return
	}

	if leftCnt == rightCnt {
		helper(prefix+"(", leftCnt - 1, rightCnt)
	} else if leftCnt < rightCnt {
		if leftCnt > 0 {
			helper(prefix+"(", leftCnt - 1, rightCnt)
		}
		helper(prefix+")", leftCnt, rightCnt - 1)
	}
}
