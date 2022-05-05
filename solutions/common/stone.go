//https://leetcode-cn.com/problems/stone-game-vi/
//学到slice的拼接的知识点
/**
注意这里不用copy的话会改变底层数组的值,从而导致a发生改变
func doSlice1(a []int, b []int) {
	fmt.Println("a=", a, "b=", b)

	if len(a) < 2 || len(b) < 2 {
		return
	}
	var aTemp = make([]int, len(a[0:len(a)-2]))
	copy(aTemp, a[0:len(a)-2])
	aTemp = append(aTemp, a[len(a)-1:]...)
	doSlice1(aTemp, b)
}
*/
package common

import "fmt"

func RunStoneGameVI() int {
	return stoneGameVI([]int{1, 3}, []int{2, 1})
}
func stoneGameVI(aliceValues []int, bobValues []int) int {
	//alice先走,所有alice会首先做预判,根据所有的可能性,判断是否能赢,也就是拿到最多的分数,最终根据两者的得分判断输赢
	//所以第一步alice有n种可能性
	a, b := findMaxScore(aliceValues, bobValues, "alice")
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}

}

//输入: 被挑选之后的values,以及谁先手;  输出, 两人的得分,alice在前,bob在后
func findMaxScore(aliceValues []int, bobValues []int, name string) (a int, b int) {
	if len(aliceValues) == 0 && len(bobValues) == 0 {
		return 0, 0
	}
	fmt.Println("findMaxScore", "alice=", aliceValues, "bob=", bobValues, "name=", name)
	defer func() {
		fmt.Println("a=", a, "b=", b)
	}()
	if name == "bob" { //找到让bob最大的
		var bobMaxValue = 0
		var aliceMaxValue = 0
		for i := 0; i < len(bobValues); i++ {
			var bobTempValue = make([]int, len(bobValues[0:i]))
			copy(bobTempValue, bobValues[0:i])
			bobTempValue = append(bobTempValue, bobValues[i+1:]...)
			var aliceTempValue = make([]int, len(aliceValues[0:i]))
			copy(aliceTempValue, aliceValues[0:i])
			aliceTempValue = append(aliceTempValue, aliceValues[i+1:]...)
			aliceMax, bobMax := findMaxScore(aliceTempValue, bobTempValue, "alice")
			if bobValues[i]+bobMax > bobMaxValue {
				bobMaxValue = bobValues[i] + bobMax
				aliceMaxValue = aliceMax
			}
		}
		return aliceMaxValue, bobMaxValue
	} else if name == "alice" {
		var aliceMaxValue = 0
		var bobMaxValue = 0
		for i := 0; i < len(aliceValues); i++ {
			var bobTempValue = make([]int, len(bobValues[0:i]))
			copy(bobTempValue, bobValues[0:i])
			bobTempValue = append(bobTempValue, bobValues[i+1:]...)
			var aliceTempValue = make([]int, len(aliceValues[0:i]))
			copy(aliceTempValue, aliceValues[0:i])
			aliceTempValue = append(aliceTempValue, aliceValues[i+1:]...)
			aliceMax, bobMax := findMaxScore(aliceTempValue, bobTempValue, "bob")
			if aliceValues[i]+aliceMax > aliceMaxValue {
				aliceMaxValue = aliceValues[i] + aliceMax
				bobMaxValue = bobMax
			}
		}
		return aliceMaxValue, bobMaxValue
	}
	return 0, 0
}
