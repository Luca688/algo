package main

//删除链表中重复的节点
//在一个排序的链表中，存在重复的结点，请删除该链表中重复的结点，重复的结点不保留，返回链表头指针。
// 例如，链表1->1->1->1->1->1->2->3->3->3->3->4->4->5 处理后为 1->2->5
func SortLinkListRepeat(head *Node) *Node {
	if head == nil {
		return head
	}

	currentNode := head
	var lastNode *Node

	for currentNode != nil{

		if currentNode.Next == nil {
			break
		}else {
			//当前元素值与下一个相等
			if currentNode.Val == currentNode.Next.Val {
				//遍历丢弃中间相同的元素
				for  {
					if currentNode.Next != nil && currentNode.Val == currentNode.Next.Val {
						currentNode.Next = currentNode.Next.Next
					}else {
						break
					}
				}
				//头节点就有相同元素
				if lastNode == nil {
					lastNode = currentNode.Next
					head = lastNode
				}else {
					lastNode.Next = currentNode.Next
				}
				currentNode = currentNode.Next

			}else {
				lastNode = currentNode
				currentNode = currentNode.Next
			}
		}
	}

	return head
}



//输入两个单调递增的链表，输出两个链表合成后的链表，当然我们需要合成后的链表满足单调不减规则。
//{1,3,5},{2,4,6} 返回值 {1,2,3,4,5,6}
//{2,11,77,79,80}{3,4,5,82,83,84} 返回  {2,3,4,5,11,77,79,80,83,84}

func IncreasingLinkMerge(head1 *Node, head2 *Node) *Node{

	var res *Node
	var resTail *Node
	leftCurrent := head1
	rightCurrent := head2

	for {
		if leftCurrent == nil {
			resTail.Next = rightCurrent
			break
		}
		if rightCurrent == nil{
			resTail.Next = leftCurrent
			break
		}

		if leftCurrent.Val <= rightCurrent.Val {

			if res == nil {
				res = leftCurrent
			}else {
				resTail.Next = leftCurrent
			}
			resTail = leftCurrent
			leftCurrent = leftCurrent.Next
		}else {

			if res == nil{
				res = rightCurrent
			}else {
				resTail.Next = rightCurrent
			}
			resTail = rightCurrent
			rightCurrent = rightCurrent.Next
		}
	}

	return res

}
