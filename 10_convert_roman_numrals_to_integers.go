// 20220224  10.Convert Roman numerals to integers

// 罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
// 字符          数值
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// 例如， 罗马数字 2 写做 II ，即为两个并列的 1 。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

// 通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，
// 例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

// I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
// X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
// C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
// 给定一个罗马数字，将其转换成整数。

// 示例 1:
// 输入: s = "III"
// 输出: 3

// 示例 2:
// 输入: s = "IV"
// 输出: 4

// 示例 3:
// 输入: s = "IX"
// 输出: 9

// 示例 4:
// 输入: s = "LVIII"
// 输出: 58
// 解释: L = 50, V= 5, III = 3.

// 示例 5:
// 输入: s = "MCMXCIV"
// 输出: 1994
// 解释: M = 1000, CM = 900, XC = 90, IV = 4.
// https://leetcode-cn.com/problems/integer-to-roman/solution/zheng-shu-zhuan-luo-ma-shu-zi-by-leetcod-75rs/
package main

var valueSymbols = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	roman := []byte{}
	for _, vs := range valueSymbols {
		for num >= vs.value {
			num -= vs.value
			roman = append(roman, vs.symbol...)
		}
		if num == 0 {
			break
		}
	}
	return string(roman)
}

// 方法一：模拟
// 思路
// 根据罗马数字的唯一表示法，为了表示一个给定的整数 \textit{num}num，我们寻找不超过 \textit{num}num 的最大符号值，
// 将 \textit{num}num 减去该符号值，然后继续寻找不超过 \textit{num}num 的最大符号值，将该符号拼接在上一个找到的符号之后，
// 循环直至 \textit{num}num 为 00。最后得到的字符串即为 \textit{num}num 的罗马数字表示。
// 编程时，可以建立一个数值-符号对的列表 \textit{valueSymbols}valueSymbols，按数值从大到小排列。
// 遍历 \textit{valueSymbols}valueSymbols 中的每个数值-符号对，若当前数值 \textit{value}value 不超过 \textit{num}num，
// 则从 \textit{num}num 中不断减去 \textit{value}value，直至 \textit{num}num 小于 \textit{value}value，然后遍历下一个数值-符号对。
// 若遍历中 \textit{num}num 为 00 则跳出循环。

// 方法二：硬编码数字
// 回顾前言中列出的这 1313 个符号，可以发现：
// 千位数字只能由 \texttt{M}M 表示；
// 百位数字只能由 \texttt{C}C，\texttt{CD}CD，\texttt{D}D 和 \texttt{CM}CM 表示；
// 十位数字只能由 \texttt{X}X，\texttt{XL}XL，\texttt{L}L 和 \texttt{XC}XC 表示；
// 个位数字只能由 \texttt{I}I，\texttt{IV}IV，\texttt{V}V 和 \texttt{IX}IX 表示。
// 这恰好把这 1313 个符号分为四组，且组与组之间没有公共的符号。因此，整数 \textit{num}num 的十进制表示中的每一个数字都是可以单独处理的。
// 进一步地，我们可以计算出每个数字在每个位上的表示形式，整理成一张硬编码表。
// 利用模运算和除法运算，我们可以得到 \textit{num}num 每个位上的数字：
// thousands_digit = num / 1000
// hundreds_digit = (num % 1000) / 100
// tens_digit = (num % 100) / 10
// ones_digit = num % 10
// 最后，根据 \textit{num}num 每个位上的数字，在硬编码表中查找对应的罗马字符，并将结果拼接在一起，即为 \textit{num}num 对应的罗马数字。

var (
	thousands = []string{"", "M", "MM", "MMM"}
	hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones      = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

func intToRoman(num int) string {
	return thousands[num/1000] + hundreds[num%1000/100] + tens[num%100/10] + ones[num%10]
}
