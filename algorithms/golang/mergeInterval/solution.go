package main

import "fmt"

func main() {
	testCase1 := [][]int{}
	testCase2 := [][]int{{1, 2}, {3, 4}, {5, 6}}
	testCase3 := [][]int{{1, 3}, {2, 4}, {5, 6}}
	testCase4 := [][]int{{1, 6}, {2, 4}, {5, 6}}

	fmt.Printf("input: %+v, ans: %+v\n", testCase1, merge(testCase1))
	fmt.Printf("input: %+v, ans: %+v\n", testCase2, merge(testCase2))
	fmt.Printf("input: %+v, ans: %+v\n", testCase3, merge(testCase3))
	fmt.Printf("input: %+v, ans: %+v\n", testCase4, merge(testCase4))

	fmt.Printf("output: %+v\n", getMatrix(5, 2, 2, 3))
	fmt.Printf("output: %+v\n", getMatrix(2, 2, 1, 1))
	fmt.Printf("output: %+v\n", getMatrix(1, 1, 1, 1))
}

// 合并区间
func merge(arr [][]int) [][]int {
	var ans [][]int
	for idx := 0; idx < len(arr); idx++ {
		curInterval := make([]int, 2)
		curInterval[0] = arr[idx][0]
		curInterval[1] = arr[idx][1]
		for idx < len(arr) - 1 && curInterval[1] >= arr[idx+1][0] {
			idx++
			curInterval[1] = max(curInterval[1], arr[idx][1])
		}
		ans = append(ans, curInterval)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 回型矩阵求子矩阵
func getMatrix(n, m, x, y int) [][]int {
	x, y = x - 1, y - 1	// 将目标坐标转化为矩阵下标
	// check invalid
	if n <= 0 || m <= 0 || x < 0 || y < 0 || x + m > n || y + m > n{
		panic("invalid input")
	}

	// get matrix
	matrix := make([][]int, n)
	for idx := 0; idx < n; idx++ {
		matrix[idx] = make([]int, n)
	}

	left, right, top, bottom := 0, n - 1, 0, n - 1
	cnt := 1
	total := n * n
	for cnt <= total {
		for idx := top; idx <= bottom; idx++ {
			matrix[idx][left] = cnt
			cnt++
		}
		left++
		for idx := left; idx <= right; idx++ {
			matrix[bottom][idx] = cnt
			cnt++
		}
		bottom--
		for idx := bottom; idx >= top; idx-- {
			matrix[idx][right] = cnt
			cnt++
		}
		right--
		for idx := right; idx >= left; idx-- {
			matrix[top][idx] = cnt
			cnt++
		}
		top++
	}

	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		cur := make([]int, m)
		for j := 0; j < m; j++ {
			cur[j] = matrix[x+i][y+j]
		}
		ans[i] = cur
	}
	return ans
}