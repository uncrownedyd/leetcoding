package stringCompression

import "strconv"

func compress(chars []byte) int {
	writeIdx := 0
	l := len(chars)
	firstIdx := 0
	for firstIdx < l {
		nextIdx := firstIdx + 1
		dupCnt := 1
		for nextIdx < l && chars[firstIdx] == chars[nextIdx] {
			dupCnt++
			nextIdx++
		}
		chars[writeIdx] = chars[firstIdx]
		writeIdx++
		if dupCnt > 1 {
			writeIdx = helper(chars, writeIdx, dupCnt)
		}
		firstIdx = nextIdx
	}
	return 0
}

func helper(chars []byte, start, num int) int {
	for _, e := range strconv.Itoa(num) {
		chars[start] = byte(e)
		start++
	}
	return start
}
