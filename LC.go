import (
	"math"
	"strings"
)

// 1.给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/two-sum

func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}
	for i, n := range nums {
		if j, ok := hashMap[target-n]; ok {
			return []int{j, i}
		}
		hashMap[n] = i
	}
	return nil
}

// 2.给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/add-two-numbers

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		//fmt.Println(a,b, carry)
		c := (a + b + carry) % 10
		//fmt.Println(c)
		carry = (a + b + carry) / 10
		if head == nil {
			head = &ListNode{Val: c}
			tail = head
		} else {
			tail.Next = &ListNode{Val: c}
			tail = tail.Next
		}
	}
	if carry != 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

// 3.给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// 4.给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

// 算法的时间复杂度应该为 O(log (m+n)) 。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	length := m + n
	left, right := 0, 0
	l1, l2 := 0, 0
	// 如果length为奇数，那问题转换为求第k小的数(k=length/2+1)
	// 所以循环k-1次，迭代right即可得到第k小的数
	// 如果length为偶数，那么问题转换为求第k-1小的数和第k小的数两个数的平均值
	for i := 0; i <= length/2; i++ {
		left = right
		if l1 < m && (l2 >= n || nums1[l1] < nums2[l2]) {
			right = nums1[l1]
			l1++
		} else {
			right = nums2[l2]
			l2++
		}
	}
	if length%2 == 1 {
		return float64(right)
	}
	return float64(left+right) / 2.0
}

// 6. 整数反转
// 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
// 如果反转后整数超过 32 位的有符号整数的范围 [−2^31,  2^31 − 1] ，就返回 0。
// 假设环境不允许存储 64 位整数（有符号或无符号）。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/reverse-integer
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func reverse(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return
}

// 7. 字符串转换整数 (atoi)
// 请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。

// 函数 myAtoi(string s) 的算法如下：

// 读入字符串并丢弃无用的前导空格
// 检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
// 读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
// 将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
// 如果整数数超过 32 位有符号整数范围 [−231,  231 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被固定为 −231 ，大于 231 − 1 的整数应该被固定为 231 − 1 。
// 返回整数作为最终结果。
// 注意：

// 本题中的空白字符只包括空格字符 ' ' 。
// 除前导空格或数字后的其余字符串外，请勿忽略 任何其他字符。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/string-to-integer-atoi
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func myAtoi(s string) int {
	var (
		negative  bool // 负数
		i         int32
		signed    bool // 是否出现过'-'，'+'
		startZero bool // 是否开始是0,再出现'-'，'+'，return 0
	)
	for k, v := range s { // 去除前空格
		if v != ' ' {
			s = s[k:]
			break
		}
	}
	for _, v := range s {
		if v == '0' && !signed && i == 0 { // 0000-898
			startZero = true
			continue
		}
		if startZero && (v == '+' || v == '-') {
			break
		}
		if v == '+' && !signed { // +
			signed = true
			continue
		}
		if v == '-' && !signed && i == 0 { // -
			negative = true
			signed = true
			continue
		}
		if v < '0' || v > '9' {
			break
		}
		if negative { //负数
			if -i < (math.MinInt32+(v-'0'))/10 {
				i = math.MinInt32
				break
			}
		} else {
			if i > (math.MaxInt32-(v-'0'))/10 {
				i = math.MaxInt32
				break
			}
		}
		i = v - '0' + i*10
	}
	if negative {
		i = -i
	}
	return int(i)
}

////////////////////////////////////////////////////////////////////
func myAtoi(str string) int {
	return convert(clean(str))
}

func clean(s string) (sign int, abs string) {
	// 先去除首尾空格
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}
	// 判断第一个字符
	switch s[0] {
	// 有效的
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sign, abs = 1, s
	// 有效的，正号
	case '+':
		sign, abs = 1, s[1:]
	// 有效的，负号
	case '-':
		sign, abs = -1, s[1:]
	// 无效的，当空字符处理，并且直接返回
	default:
		abs = ""
		return
	}
	for i, b := range abs {
		// 遍历第一波处理过的字符，如果直到第i个位置有效，那就取s[:i]，从头到这个有效的字符，剩下的就不管了，也就是break掉
		// 比如 s=123abc，那么就取123，也就是s[:3]
		if b < '0' || '9' < b {
			abs = abs[:i]
			// 一定要break，因为后面的就没用了
			break
		}
	}
	return
}

// 接收的输入是已经处理过的纯数字
func convert(sign int, absStr string) int {
	absNum := 0
	for _, b := range absStr {
		// b - '0' ==> 得到这个字符类型的数字的真实数值的绝对值
		absNum = absNum*10 + int(b-'0')
		// 检查溢出
		switch {
		case sign == 1 && absNum > math.MaxInt32:
			return math.MaxInt32
		// 这里和正数不一样的是，必须和负号相乘，也就是变成负数，否则永远走不到里面
		case sign == -1 && absNum*sign < math.MinInt32:
			return math.MinInt32
		}
	}
	return sign * absNum
}
