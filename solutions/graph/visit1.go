package graph

import (
	"encoding/json"
	"fmt"
)

//https://leetcode-cn.com/problems/construct-quad-tree/
//一个很低级的错误
//node.TopLeft.Val==node.TopRight.Val&&node.TopRight.Val==node.BottomRight.Val&&node.BottomRight.Val==node.BottomLeft.Val
//和node.TopLeft.Val==node.TopRight.Val==node.BottomRight.Val==node.BottomLeft.Val
//!!!一定注意==是不能传递的
//这道题做的不太顺,一开始没有把完整的思路想清楚,导致在做的过程中不断改动
//另一个是,在写代码的时候,对于其中的每一个细节在一开始就想清楚写对,否则在后期的调试过程中容易造成反复,引起更大的麻烦

//Definition for a QuadTree node.
type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func RunConstruct() *Node {
	return construct([][]int{{1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}})
}
func construct(grid [][]int) *Node {
	return constructItem(grid, 0, len(grid)-1, 0, len(grid)-1)
}

func constructItem(grid [][]int, i1, i2, j1, j2 int) (node *Node) {
	fmt.Println("i1=", i1, "i2=", i2, "j1=", j1, "j2=", j2)
	defer func() {
		str, _ := json.Marshal(node)
		fmt.Println("node=", string(str))
	}()
	if i2-i1 == 1 {
		isSame := true
		standard := grid[i1][j1]
		for i := i1; i <= i2; i++ {
			for j := j1; j <= j2; j++ {
				if grid[i][j] != standard {
					isSame = false
					break
				}
			}
			if isSame == false {
				break
			}
		}
		if isSame {
			return &Node{
				Val:    grid[i1][j1] == 1,
				IsLeaf: true,
			}
		} else {
			return &Node{
				IsLeaf:      false,
				TopLeft:     &Node{Val: grid[i1][j1] == 1, IsLeaf: true},
				TopRight:    &Node{Val: grid[i1][j2] == 1, IsLeaf: true},
				BottomLeft:  &Node{Val: grid[i2][j1] == 1, IsLeaf: true},
				BottomRight: &Node{Val: grid[i2][j2] == 1, IsLeaf: true},
			}
		}
	} else {
		ii1, ii2, ii3, ii4 := i1, (i2-i1+1)/2+i1-1, (i2-i1+1)/2+i1, i2
		jj1, jj2, jj3, jj4 := j1, (j2-j1+1)/2+j1-1, (j2-j1+1)/2+j1, j2
		// fmt.Println(j1,j2,j3,j4)
		node := &Node{}
		node.TopLeft = constructItem(grid, ii1, ii2, jj1, jj2)
		node.TopRight = constructItem(grid, ii1, ii2, jj3, jj4)
		node.BottomLeft = constructItem(grid, ii3, ii4, jj1, jj2)
		node.BottomRight = constructItem(grid, ii3, ii4, jj3, jj4)
		if node.TopLeft.IsLeaf && node.TopRight.IsLeaf && node.BottomRight.IsLeaf && node.BottomLeft.IsLeaf &&
			(node.TopLeft.Val == node.TopRight.Val && node.TopRight.Val == node.BottomRight.Val && node.BottomRight.Val == node.BottomLeft.Val) {
			node.Val = node.TopLeft.Val
			node.IsLeaf = true
			node.TopLeft = nil
			node.TopRight = nil
			node.BottomLeft = nil
			node.BottomRight = nil
			return node
		} else {
			node.IsLeaf = false
			return node
		}

	}

}
