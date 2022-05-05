/**
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。

candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。

对于给定的输入，保证和为 target 的不同组合数少于 150 个。



示例 1：

输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package backtrack

import (
	"fmt"
	"sort"
)

func RunCombinationSum() [][]int {
	return combinationSum([]int{2, 3, 6, 7}, 7)
}

// func combinationSum(candidates []int, target int) [][]int {
// 	//使用backtrack思想:
// 	//最小单位: find0->find2,find3,find6,find7, resultItem依次加入candidate.
// 	//假想返回的数组,有n个元素,每个元素位置有len个选项,我们需要依次放入这len个选项
// 	//!!!存在重复的选项---加限制条件, 每个元素位置所使用的index需要递增
// 	var array []int
// 	var result [][]int
// 	// sort.Slice(candidates,func(i, j int) bool {
// 	// 	return candidates[i]<candidates[j]
// 	// })
// 	sort.IntSlice(candidates).Sort()
// 	var maxLen = target / candidates[0]
// 	var add func(i int, target int, j int)
// 	add = func(i int, target int, j int) { //表示要放入第i个元素,以及剩余的目标target,j表示上一个元素使用的candidate的index,下一次遍历需要从j开始
// 		fmt.Println("i=", i, "target=", target)
// 		if target == 0 {
// 			var copyItem = make([]int, len(array))
// 			copy(copyItem, array)
// 			result = append(result, copyItem)
// 		}
// 		if i == maxLen {
// 			return
// 		}
// 		// for _, item := range candidates {
// 		// 	fmt.Println("array=", array, "append", item)
// 		// 	array = append(array, item)
// 		// 	fmt.Println("array=", array)
// 		// 	add(i+1, target-item,)
// 		// 	array = array[:len(array)-1]
// 		// }
// 		for ci := j; ci < len(candidates); ci++ {
// 			fmt.Println("array=", array, "append", candidates[ci])
// 			array = append(array, candidates[ci])
// 			fmt.Println("array=", array)
// 			add(i+1, target-candidates[ci], ci)
// 			array = array[:len(array)-1]
// 		}
// 	}
// 	add(0, target, 0)
// 	return result
// }

func combinationSum(candidates []int, target int) [][]int {
	//给定0, 得到0+candidate[i]的组合, 当最后一轮每个都大于target时结束(candidate做下排序,可以只比较最小值)
	//得到一个target之后,return,并将结果放入result
	//不可行,遍历条件会导致有重复的组合,需要按照顺序依次放入

	type resultType []int
	var targetMap = map[int][][]int{}
	var targetIndexMap = map[int][]int{}
	candidates = sort.IntSlice(candidates)
	var times = target / candidates[0]
	for index, item := range candidates {
		targetMap[item] = append([][]int{}, []int{item})
		targetIndexMap[item] = append(targetIndexMap[item], index)
	}
	//记录下targetMap第i个的初始index
	fmt.Println("target", targetMap)
	for index, item := range candidates {
		if times < 0 {
			break
		}
		for key, value := range targetMap {

			fmt.Println("key=", key, "item=", item, "new=", item+key)
			// fmt.Println("targetMap=", targetMap)
			newTarget := item + key
			for valueIndex, valueItem := range value {
				if len(targetIndexMap[key]) > 0 && index < targetIndexMap[key][valueIndex] { //如果小于前一个放入的index,跳过
					continue
				}
				// fmt.Println("valueItem=", valueItem)
				var valueItemCopy = make([]int, len(valueItem))
				copy(valueItemCopy, valueItem)
				valueItemCopy = append(valueItemCopy, item)
				targetMap[newTarget] = append(targetMap[newTarget], valueItemCopy)
				targetIndexMap[newTarget] = append(targetIndexMap[newTarget], index)
			}
		}
		times--
	}

	return targetMap[target]
	// find = func(target int) [][]int {

	// 	//给定target找所有的数组
	// 	var result [][]int
	// 	if target <= 0 {
	// 		return result
	// 	}
	// 	for _, item := range candidates {
	// 		if target == item {
	// 			return [][]int{[]int{item}}
	// 		}
	// 	}
	// 	for _, item := range candidates {
	// 		var itemArray = find(target - item)
	// 		for _, targetArray := range itemArray {
	// 			result = append(result, append(append([]int{}, targetArray...), item))
	// 		}
	// 	}
	// 	fmt.Println("result=", result, "target=", target)
	// 	return result
	// }
	// return find(target)

}

// func combinationSum(candidates []int, target int) [][]int {
// 	//给定的数组是否有序,按照有序来考虑
// 	var sortedArray=make([]int,200)
// 	for _,candidateItem:=range candidates{
// 		sortedArray[candidateItem]=candidateItem
// 	}
// 	 i:=0
// 	for _,item:=range sortedArray{
// 		if item!=0{
// 			candidates[i]=item
// 			i=i+1
// 		}
// 	}
// 	//保存一个map便于做遍历
// 	var candidateMap=make(map[int]bool)

// 		for _,item:=range candidates{
// 			candidateMap[item]=true
// 		}
// 		return find(candidates,target,candidateMap)
// 	}
// 	func find(candidates[]int,target int,candidateMap map[int]bool)[][]int{
// 		var result=make([][]int,0)
// 		if candidates[0]>target{
// 			return result
// 		}
// 		if candidateMap[target]{
// 			result=append(result,[]int{target})
// 		}
// 		//每次取出来一个current,在剩下的数组中寻找target-current
// 		for index,item:=range candidates{
// 			resultItem:= find(candidates[index:],target-item,candidateMap)
// 			if len(resultItem)>0{
// 				for resultItemIndex,_:=range resultItem{
// 					var tempArray=make([]int,len(resultItem[resultItemIndex])+1)
// 					tempArray[0]=item
// 					for i:=1;i<len(tempArray);i=i+1{
// 						tempArray[i]=resultItem[resultItemIndex][i-1]
// 					}
// 					result=append(result,tempArray)
// 				}
// 			}
// 		}
// 		return result
// 	}
