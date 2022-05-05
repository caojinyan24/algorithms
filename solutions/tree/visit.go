//https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
//给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
package tree

func RunVisitByLevel() [][]int {
	return levelOrder(&TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 9}}})
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func levelOrder(root *TreeNode) [][]int {
// 	if root == nil {
// 		return [][]int{}
// 	}
// 	children := visitChildren([]*TreeNode{root})
// 	var allNodes = [][]*TreeNode{0: {root}}
// 	for {
// 		if len(children) == 0 {
// 			break
// 		}
// 		allNodes = append(allNodes, children)
// 		children = visitChildren(children)
// 	}
// 	var result = make([][]int, len(allNodes))
// 	for index, item := range allNodes {
// 		for _, node := range item {
// 			result[index] = append(result[index], node.Val)
// 		}
// 	}
// 	return result
// }
// func visitChildren(parent []*TreeNode) []*TreeNode { //根据给定的parent列表,返回childer的列表
// 	var result = make([]*TreeNode, 0)
// 	for _, node := range parent {
// 		if node.Left != nil {
// 			result = append(result, node.Left)
// 		}
// 		if node.Right != nil {
// 			result = append(result, node.Right)
// 		}
// 	}
// 	return result
// }

// % 执行用时：
// % 0 ms
// % , 在所有 Go 提交中击败了
// % 100.00%
// % 的用户
// % 内存消耗：
// % 2.7 MB
// % , 在所有 Go 提交中击败了
// % 30.94%
// % 的用户
// % 通过测试用例：
// % 34 / 34

//参考下题解,这里返回结果的二次遍历可以做下简化
//代码修改成如下方式,节省内存消耗

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	children := []*TreeNode{root}
	var result = [][]int{0: []int{root.Val}}
	var resultItem []int
	for {
		children, resultItem = visitChildren(children)
		if len(children) == 0 {
			break
		}
		result = append(result, resultItem)
	}
	return result
}
func visitChildren(parent []*TreeNode) ([]*TreeNode, []int) { //根据给定的parent列表,返回childer的列表
	var result = make([]*TreeNode, 0)
	var resultInt = make([]int, 0)
	for _, node := range parent {
		if node.Left != nil {
			result = append(result, node.Left)
			resultInt = append(resultInt, node.Left.Val)
		}
		if node.Right != nil {
			result = append(result, node.Right)
			resultInt = append(resultInt, node.Right.Val)

		}
	}
	return result, resultInt
}

// % 执行用时：
// % 0 ms
// % , 在所有 Go 提交中击败了
// % 100.00%
// % 的用户
// % 内存消耗：
// % 2.7 MB
// % , 在所有 Go 提交中击败了
// % 35.67%
// % 的用户
// % 通过测试用例：
// % 34 / 34

/**
总结: 这道题是二叉树的广度遍历, 我依旧是使用了递归调用的方式,把整个问题拆解成一套操作的叠加,然后把这套操作抽出来做一个问题解法
标准题解是把这套叠加的操作通过一个数组进行控制,每次对数组进行变换,从而达到嵌套操作的目的
 试运行了下,内存的消耗更少(2.6M),我想是通过公用一个内存,而不是两个方法使用独立的内存栈,达到节省内存的目的
 所以,下次可以尝试下通过控制循环变量的方式,达到操作嵌套的目的
*/
//标准题解
// func levelOrder(root *TreeNode) [][]int {
// 	ret := [][]int{}
//   if root == nil {
// 	  return ret
//   }
//   q := []*TreeNode{root}
//   for i := 0; len(q) > 0; i++ {
// 	  ret = append(ret, []int{})
// 	  p := []*TreeNode{}
// 	  for j := 0; j < len(q); j++ {
// 		  node := q[j]
// 		  ret[i] = append(ret[i], node.Val)
// 		  if node.Left != nil {
// 			  p = append(p, node.Left)
// 		  }
// 		  if node.Right != nil {
// 			  p = append(p, node.Right)
// 		  }
// 	  }
// 	  q = p
//   }
//   return ret
// }

// 执行用时：
// 0 ms
// , 在所有 Go 提交中击败了
// 100.00%
// 的用户
// 内存消耗：
// 2.6 MB
// , 在所有 Go 提交中击败了
// 97.33%
// 的用户
// 通过测试用例：
// 34 / 34
