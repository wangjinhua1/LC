// 09_盛最多水的容器

// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。
// 说明：你不能倾斜容器。
// eg: 输入：[1,8,6,2,5,4,8,3,7]
//     输出：49

// 1,   8,    6,    2,    5,    4,     8,    3,    7
// 1    2     3     4     5     6      7     8     9

// 双指针 https://leetcode-cn.com/problems/container-with-most-water/solution/sheng-zui-duo-shui-de-rong-qi-by-leetcode-solution/

// 执行用时：68 ms, 在所有 Go 提交中击败了92.04%的用户
// 内存消耗：8.5 MB, 在所有 Go 提交中击败了74.88%的用户

package main

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	area := 0
	for i != j {
		s := (j - i) * min(height[i], height[j])
		if s > area {
			area = s
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return area
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
