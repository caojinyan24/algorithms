package main

import (
	"algorithms/solutions/backtrack"
	// "algorithms/solutions/common"
	"fmt"
)

func main() {
	// var result interface{}
	// result := tree.RunVisitByLevel()
	// result = common.RunMemLeak()
	// result = strings.RunAlertNames()

	// doSlice(array, array[1:])
	// fmt.Println("array=", array)

	// result = common.RunStoneGameVI()
	// fmt.Println("result=", result)

	// var array = []int{1, 2, 3, 4, 5, 6}
	// var temp = array[1:3]
	// temp = append(temp, array[4:5]...)
	// fmt.Println("temp=", temp)
	// fmt.Println("array=", array)
	// fmt.Println(common.RunMinimalHeaviestSetA())
	// fmt.Println(common.RunSubSets
	// fmt.Println(backtrack.RunGenerateParenthesis())
	fmt.Println(backtrack.RunCombinationSum())
}

/**
1. 审题,理解清楚题目的意思
2. 对于题目没有明说的条件,做二次确认
3. 在初始的时候,多考虑些边界条件或者case,可以避免在思考的时候思路跑偏
4. 通过改变变量的值完成嵌套操作(使用递归会额外消耗内存)
*/

/**
注意array的输出
var array = []int{1, 2, 3, 4, 5, 6}
var temp = array[1:3]
temp = append(temp, array[4:5]...)
fmt.Println("temp=", temp)
fmt.Println("array=", array)

outPut:
temp= [2 3 5]
array= [1 2 3 5 5 6]
*/

/**
双指针,适用于两边边界动态变换||长度指定的情况(https://leetcode-cn.com/problems/container-with-most-water/)
**/

//每次做完一道题,还是需要好好琢磨下,认真做下总结,否则做完之后就忘掉了,做题的目的是成长,不是看自己还有多少道题没有做
