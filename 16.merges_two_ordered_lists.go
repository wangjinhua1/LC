// 16. 合并两个有序链表 Merges two ordered lists
// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

// 递归
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

// 迭代
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				cur.Next = l1
				l1 = l1.Next
			} else {
				cur.Next = l2
				l2 = l2.Next
			}
			cur = cur.Next
		} else if l1 != nil {
			cur.Next = l1
			break
		} else {
			cur.Next = l2
			break
		}
	}
	return head.Next
}
