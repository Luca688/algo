package structure

func GeneralStack(slice []int) *Node {

	var head *Node

	for _, val := range slice {
		node := GeneralNode(val)
		node.Next = head
		head = node
	}

	return head
}

//push

func PushStack(head *Node, val int) *Node {
	node := GeneralNode(val)
	node.Next = head
	return  node
}

//pop
func PopStack(heaad *Node) (newHead *Node, val int){

	if heaad.IsEmpty() {
		return heaad, -999
	}

	return heaad.Next, heaad.Val
}

//isEmpty

func (head *Node)IsEmpty() bool {
	if head == nil {
		return  true
	}else {
		return  false
	}
}
