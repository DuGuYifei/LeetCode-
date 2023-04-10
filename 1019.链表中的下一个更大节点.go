package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=1019 lang=golang
 *
 * [1019] 链表中的下一个更大节点
 *
 * https://leetcode.cn/problems/next-greater-node-in-linked-list/description/
 *
 * algorithms
 * Medium (60.98%)
 * Likes:    254
 * Dislikes: 0
 * Total Accepted:    34K
 * Total Submissions: 54.7K
 * Testcase Example:  '[2,1,5]'
 *
 * 给定一个长度为 n 的链表 head
 *
 * 对于列表中的每个节点，查找下一个 更大节点 的值。也就是说，对于每个节点，找到它旁边的第一个节点的值，这个节点的值 严格大于 它的值。
 *
 * 返回一个整数数组 answer ，其中 answer[i] 是第 i 个节点( 从1开始 )的下一个更大的节点的值。如果第 i
 * 个节点没有下一个更大的节点，设置 answer[i] = 0 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：head = [2,1,5]
 * 输出：[5,5,0]
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：head = [2,7,4,3,5]
 * 输出：[7,0,5,5,0]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点数为 n
 * 1 <= n <= 10^4
 * 1 <= Node.val <= 10^9
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
	res := make([]int, 0)
	q := make([][]int, 0)
	q = append(q, []int{head.Val, 0})
	res = append(res, 0)
	cur := head.Next
	i := 1
	for cur != nil {
		res = append(res, 0)
		for len(q) > 0 && cur.Val > q[len(q)-1][0] {
			res[q[len(q)-1][1]] = cur.Val
			q = q[:len(q)-1]
		}
		q = append(q, []int{cur.Val, i})
		i++
		cur = cur.Next
	}
	return res
}

// @lc code=end
