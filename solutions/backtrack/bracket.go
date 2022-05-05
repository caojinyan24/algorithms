//https://leetcode-cn.com/problems/generate-parentheses/
//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

package backtrack

// import "fmt"

// import "fmt"
func RunGenerateParenthesis() []string {
	return generateParenthesis(3)
}
func generateParenthesis(n int) []string {
	//左括号的个数和右括号的个数,右括号的个数需要小于等于左括号
	// 通过回溯,每一次的选择包括左括号和右括号,结束条件是左括号等于n
	//每一个位置有两种选项

	type matrixItem []string
	var matrix = make([][]matrixItem, n+1)
	for i := 0; i < n+1; i++ {
		var item = make([]matrixItem, n+1)
		matrix[i] = item
	}
	matrix[0][0] = []string{""}
	// matrix[1][0] = []string{"("}
	// matrix[1][1] = []string{"()"}
	var appendStr = func(list []string, s string) []string {
		var result = make([]string, len(list))
		for index, item := range list {
			result[index] = item + s
		}
		return result
	}
	for i := 0; i <= n; i++ {
		for j := 0; j <= i; j++ {
			// fmt.Println("debug", i, j, matrix[i][j])
			if i >= 1 {
				matrix[i][j] = append(matrix[i][j], appendStr(matrix[i-1][j], "(")...)
			}
			if j >= 1 {
				matrix[i][j] = append(matrix[i][j], appendStr(matrix[i][j-1], ")")...)
			}
			// fmt.Println("matrix", i, j, matrix[i][j])
		}
	}
	return matrix[n][n]
	// var resultItem = ""
	// var result = []string{}
	// var add func(int, int)
	// add = func(left int, right int) {
	// 	fmt.Println("current item=", resultItem, "left=", left, "right=", right)
	// 	if right == 3 {
	// 		result = append(result, resultItem)
	// 	}
	// 	if left < 3 {
	// 		resultItem = resultItem + "("
	// 		add(left+1, right)
	// 		resultItem = resultItem[0 : len(resultItem)-1]

	// 	}

	// 	if left > 0 && right < left {
	// 		resultItem = resultItem + ")"
	// 		add(left, right+1)
	// 		resultItem = resultItem[0 : len(resultItem)-1]

	// 	}
	// }
	// add(0, 0)
	// return result

}

/**

//current item=  left= 0 right= 0
current item= ( left= 1 right= 0
	current item= (( left= 2 right= 0
	current item= ((( left= 3 right= 0
	current item= ((() left= 3 right= 1
	current item= ((()) left= 3 right= 2
	current item= ((())) left= 3 right= 3
	current item= ((()))) left= 2 right= 1
	current item= ((())))( left= 3 right= 1
	current item= ((())))() left= 3 right= 2
	current item= ((())))()) left= 3 right= 3
	current item= ((())))())) left= 2 right= 2
	current item= ((())))()))( left= 3 right= 2
	current item= ((())))()))() left= 3 right= 3
	current item= ((())))()))()) left= 1 right= 1
	current item= ((())))()))())( left= 2 right= 1
	current item= ((())))()))())(( left= 3 right= 1
	current item= ((())))()))())(() left= 3 right= 2
	current item= ((())))()))())(()) left= 3 right= 3
	current item= ((())))()))())(())) left= 2 right= 2
	current item= ((())))()))())(()))( left= 3 right= 2
	current item= ((())))()))())(()))() left= 3 right= 3
*/
