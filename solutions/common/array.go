// https://leetcode-cn.com/problems/incremental-memory-leak/submissions/

// 给你两个整数 memory1 和 memory2 分别表示两个内存条剩余可用内存的位数。现在有一个程序每秒递增的速度消耗着内存。

// 在第 i 秒（秒数从 1 开始），有 i 位内存被分配到 剩余内存较多 的内存条（如果两者一样多，则分配到第一个内存条）。如果两者剩余内存都不足 i 位，那么程序将 意外退出 。

// 请你返回一个数组，包含 [crashTime, memory1crash, memory2crash] ，其中 crashTime是程序意外退出的时间（单位为秒）， memory1crash 和 memory2crash 分别是两个内存条最后剩余内存的位数。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/incremental-memory-leak
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//这道题目比较简单,没有看到难点,不多说了
package common

func RunMemLeak() []int {
	return memLeak(8, 10)
}
func memLeak(memory1 int, memory2 int) []int {
	var times = 1
	var left1 = memory1
	var left2 = memory2
	for i := 1; ; i++ {
		if left1 < i && left2 < i {
			break
		}
		if left1 < left2 {
			left2 = left2 - i
		} else {
			left1 = left1 - i
		}
		times++
	}
	return []int{times, left1, left2}
}
