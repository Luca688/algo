package main


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



//用两个栈来实现一个队列，完成队列的Push和Pop操作。 队列中的元素为int类型。

func TwoStackImplementQue(){

}
