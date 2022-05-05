package common

import "fmt"

//https://leetcode-cn.com/problems/subsets/

func RunSubSets() [][]int {
	return subsets([]int{1, 2, 3})
}
func subsets(nums []int) (ans [][]int) { //基本思想,每个元素在每个子集中存在两种状态,存在or不存在
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		fmt.Println("func ", cur)
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		fmt.Println("set contain", set)
		dfs(cur + 1)
		set = set[:len(set)-1] //把已经放入的cur取出来,继续遍历
		fmt.Println("set not", set)
		dfs(cur + 1)
	}
	dfs(0)
	return
}

//dfs()表示的是一个动作,是找到包含和不包含当前位置的所有子集
