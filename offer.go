package algo

import (
	"errors"
	"math"
	"strings"
)

/**
二叉搜索树的后序遍历序列 todo
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。如果是则返回true,否则返回false。假设输入的数组的任意两个数字都互不相同。（ps：我们约定空树不是二叉搜素树）
[4,8,6,12,16,14,10] true
**/
func VerifySquenceOfBST(sequence []int) bool {

	length := len(sequence)
	if length == 1 {

		return true
	}

	if length == 2 {

		if sequence[1] < sequence[0] {

			return true
		} else {
			return false
		}
	}

	outLast := sequence[:length-1]
	last := sequence[length-1]

	var left, right []int
	for key, val := range outLast {

		if val < last {

			left = append(left, val)
		} else {
			right = outLast[key:]
			break
		}
	}
	leftRes := true
	rightRes := true
	if len(left) > 0 {

		for _, val := range left {

			if val > last {

				return false
			}
		}

		leftRes = VerifySquenceOfBST(left)
	}
	if len(right) > 0 {

		for _, val := range right {

			if val < last {
				return false
			}

		}

		rightRes = VerifySquenceOfBST(right)
	}

	return leftRes && rightRes

}

/**
从上往下打印二叉树
层序遍历
**/
func PrintFromTopToBottom(root *BinaryTreeNode) []int {
	return root.LevelOrder()
}

/**
栈的压入、弹出序列 todo
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否可能为该栈的弹出顺序。
假设压入栈的所有数字均不相等。例如序列1,2,3,4,5是某栈的压入顺序，序列4,5,3,2,1是该压栈序列对应的一个弹出序列，
但4,3,5,1,2就不可能是该压栈序列的弹出序列。（注意：这两个序列的长度是相等的）
[1,2,3,4,5],[4,3,5,1,2]false
**/
func IsPopOrder(pushV []int, popV []int) bool {

	if len(pushV) != len(popV) {

		return len(pushV) == len(popV)
	}

	var stack []int
	stack = append([]int{pushV[0]}, stack...)
	pushV = pushV[1:]

	index := 0
	node := popV[0]

	for {
		if node == stack[0] {
			stack = stack[1:]
			if index < len(popV)-1 {
				index++
				node = popV[index]
			} else {
				break
			}

		} else {
			if len(pushV) > 0 {
				stack = append([]int{pushV[0]}, stack...)
				pushV = pushV[1:]
			} else {
				return false
			}

		}
	}

	return true
}

//todo 包含min函数的栈

//顺时针打印矩阵 输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字
func PrintMatrix(matrix [][]int) []int {

	var res []int

	top := 0
	left := 0
	right := len(matrix[0]) - 1
	bottom := len(matrix) - 1

	x := top
	y := left

	for len(matrix) > 0 {
		res = append(res, matrix[x][y])

		//紧剩下最后一个元素 奇数+奇数
		if left == right && top == bottom {
			break
		}

		//上
		if x == top && y < right {
			y++
			continue
		}

		//剩余一行 偶+奇
		if top == bottom {
			break
		}

		//右
		if y == right && x < bottom {
			x++
			continue
		}

		//剩余一列 奇+偶
		if right == left {
			break
		}

		//下
		if x == bottom && y <= right && y > left {
			y--
			continue

		}
		//左
		if y == left && x <= bottom && x > top {
			x--
			//到达最后一个元素
			if x == top {

				//当前是最后一圈 偶数+偶数
				if left+1 == right && top+1 == bottom {
					break
				}

				left++
				right--
				top++
				bottom--
				x = top
				y = left
				continue

			}
			continue
		}

	}
	// fmt.Println(res)
	return res
}

//二叉树的镜像
//操作给定的二叉树，将其变换为源二叉树的镜像。
//{8,6,10,5,7,9,11} {8,10,6,11,9,7,5}

func Mirror(root *BinaryTreeNode) *BinaryTreeNode {
	if root == nil {
		return nil
	}

	if root.LeftChild == nil && root.RightChild == nil {
		return root
	}

	tmp := root.RightChild
	root.RightChild = root.LeftChild
	root.LeftChild = tmp

	if root.LeftChild != nil {
		Mirror(root.LeftChild)
	}

	if root.RightChild != nil {
		Mirror(root.RightChild)
	}

	return root
}

//树的子结构
//输入两棵二叉树A，B，判断B是不是A的子结构。（ps：我们约定空树不是任意一个树的子结构）
//{8,8,#,9,#,2,#,5},{8,9,#,2} true
//{8,8,7,,9,2,#,#,#,#,4,7},{8,9,2}true
func HasSubtree(pRoot *BinaryTreeNode, cRoot *BinaryTreeNode) bool {
	if pRoot == nil {

		return false
	}

	if cRoot == nil {

		return true
	}

	_, pNodeList := pRoot.PrevOrder()

	for _, val := range pNodeList {
		if val.Val == cRoot.Val {
			issub := isSub(pRoot, cRoot)
			if issub {
				return true
			} else {
				continue
			}
		}
	}

	return false
}

func isSub(pRoot *BinaryTreeNode, cRoot *BinaryTreeNode) bool {

	if pRoot.Val != cRoot.Val {
		return false
	}

	if cRoot.LeftChild == nil && cRoot.RightChild == nil && pRoot.LeftChild == nil && pRoot.RightChild == nil && pRoot.Val == cRoot.Val {
		return true
	}

	if cRoot.LeftChild != nil {
		if pRoot.LeftChild == nil {
			return false
		} else {
			return isSub(pRoot.LeftChild, cRoot.LeftChild)
		}
	}

	if cRoot.RightChild != nil {
		if pRoot.RightChild == nil {
			return false
		} else {
			return isSub(pRoot.RightChild, cRoot.RightChild)
		}
	}

	return false
}

//调整数组顺序使奇数位于偶数前面
//输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有的奇数位于数组的前半部分，所有的偶数位于数组的后半部分，
//并保证奇数和奇数，偶数和偶数之间的相对位置不变。
//输入：[1,2,3,4]复制返回值：[1,3,2,4]
//输入：[2,4,6,5,7]复制返回值：[5,7,2,4,6]
func ReOrderArray(array []int) []int {

	var even, odd, res []int

	for _, v := range array {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}

	res = append(odd, even...)
	return res
}

// func ReOrderArrayRecursive(index int, array *[]int) {

// 	nextIndex := index + 1
// 	if nextIndex >= len(*array) {
// 		return
// 	}

// 	target := (*array)[index]
// 	isEven := true
// 	if target%2 != 0 {
// 		isEven = false
// 	}

// 	nextValue := (*array)[index+1]
// 	nextIsEven := true
// 	if nextValue%2 != 0 {
// 		nextIsEven = false
// 	}

// 	if nextIsEven == isEven {
// 		ReOrderArrayRecursive(nextIndex, array)
// 	} else {
// 		firstIndex := 0
// 		for i := nextIndex + 1; i < len(*array); i++ {
// 			firstValue := (*array)[i]
// 			firstIsEven := true
// 			if firstValue%2 != 0 {
// 				firstIsEven = false
// 			}
// 			//第一个同为奇数或者偶数的数
// 			if firstIsEven == isEven {
// 				firstIndex = i
// 				break
// 			}
// 		}
// 		if firstIndex > 0 {
// 			tmp := nextValue
// 			(*array)[nextIndex] = (*array)[firstIndex]
// 			(*array)[firstIndex] = tmp
// 			ReOrderArrayRecursive(nextIndex, array)
// 		} else {
// 			return
// 		}
// 	}
// }

//数值的整数次方
//给定一个double类型的浮点数base和int类型的整数exponent。求base的exponent次方。
//保证base和exponent不同时为0。不得使用库函数，同时不需要考虑大数问题，也不用考虑小数点后面0的位数。
func Power(base float64, exponent int) float64 {
	if exponent == 0 {
		return 1.0
	}

	res := base

	isNegative := false
	if exponent < 0 {
		exponent = exponent * -1
		isNegative = true
	}
	if exponent > 0 {
		for i := 1; i < exponent; i++ {
			res *= base
		}
	}
	if isNegative {
		res = 1 / res
	}

	return res
}

//二进制中1的个数 todo

//跳台阶扩展问题
//一只青蛙一次可以跳上1级台阶，也可以跳上2级……它也可以跳上n级。求该青蛙跳上一个n级的台阶总共有多少种跳法。
func jumpFloorII(number int) int {
	return int(math.Pow(2, float64(number-1)))
}

//跳台阶
//一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法（先后次序不同算不同的结果）。
func JumpFloor(number int) int {
	if number == 1 {
		return 1
	}
	if number == 2 {
		return 2
	}
	return JumpFloor(number-1) + JumpFloor(number-2)
}

//斐波那契数列
//输入一个整数n，请你输出斐波那契数列的第n项（从0开始，第0项为0，第1项是1）。n≤39
//0、1、1、2、3、5、8、13、21、34、55、89、144、233、377、610、987、1597、2584、4181、6765、10946、17711、28657、46368
//第0项是0，第1项是第一个1。此数列从第2项开始，每一项都等于前两项之和。
func Fibonacci(n int) int {

	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

//todo 旋转数组的最小数字

var stack1 []int //push
var stack2 []int //pop

//用两个栈实现队列
//用两个栈来实现一个队列，完成队列的Push和Pop操作。 队列中的元素为int类型。
func Push(node int) {
	stack1 = append([]int{node}, stack1...)
}

func Pop() int {

	for _, v := range stack1 {
		stack2 = append([]int{v}, stack2...)
	}
	stack1 = []int{}
	if len(stack2) > 0 {
		tmp := stack2[0]
		stack2 = stack2[1:]
		return tmp
	} else {
		return -999
	}

}

//输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
//例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。
//[1,2,3,4,5,6,7],[3,2,4,1,6,5,7] {1,2,5,3,4,6,7}
func ReConstructBinaryTree(pre []int, vin []int) *BinaryTreeNode {
	if len(pre) == 0 {
		return nil
	}
	root := NewBinaryTreeNode(pre[0])

	middle := findMiddle(vin, pre[0])
	root.LeftChild = ReConstructBinaryTree(pre[1:middle+1], vin[0:middle])
	root.RightChild = ReConstructBinaryTree(pre[middle+1:len(pre)], vin[middle+1:len(vin)])

	return root
	//todo
}

func findMiddle(arr []int, val int) int {
	index := 0
	for i, v := range arr {
		if v == val {
			index = i
			break
		}
	}

	return index
}

//从尾到头打印链表
//输入一个链表，按链表从尾到头的顺序返回一个ArrayList。
//{67,0,24,58}
//[58,24,0,67]
func PrintListFromTailToHead(head *ListNode) []int {
	var res []int

	if head == nil {
		return res
	}
	currentNode := head
	for currentNode != nil {
		tmp := []int{currentNode.Val}
		res = append(tmp, res...)
		currentNode = currentNode.Next
	}

	return res
}

//替换空格
//请实现一个函数，将一个字符串中的每个空格替换成“%20”。例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。
func ReplaceSpace(s string) string {

	res := strings.ReplaceAll(s, " ", "%20")
	return res
}

//二维数组中的查找
/**
在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
[
  [1,2,8,9],
  [2,4,9,12],
  [4,7,10,13],
  [6,8,11,15]
]
给定 target = 7，返回 true。
给定 target = 3，返回 false。
**/
func Find(target int, array [][]int) bool {

	//空数组
	if len(array) == 0 {
		return false
	}
	//只有一行
	if len(array) == 1 {
		res, _ := BinarySearch(array[0], target)
		return res
	}

	columnLength, columnMax := len(array), len(array)-1

	rowMax := len(array[0]) - 1

	//小于最小或大于最大元素
	if target < array[0][0] || target > array[columnMax][rowMax] {
		return false
	}
	//等于最大最小元素
	if target == array[0][0] || target == array[columnMax][rowMax] {
		return true
	}

	//第一列切片
	var column []int
	for i := 0; i < columnLength; i++ {
		column = append(column, array[i][0])
	}

	//行
	min := 0
	max := columnMax
	var targetRow []int

	for {
		if target == column[min] || target == column[max] {
			return true
		}

		//无法确定最后两行哪个大
		if max-min == 1 {
			for k, v := range array[min] {
				targetRow = append(targetRow, v)
				targetRow = append(targetRow, array[max][k])
			}

			break
		}

		mid := int(math.Floor(float64((max + min) / 2)))
		if target >= column[mid] {
			min = mid
		} else {
			max = mid
		}
	}

	//行查找
	res, _ := BinarySearch(targetRow, target)

	return res
}

//删除链表中重复的节点
//在一个排序的链表中，存在重复的结点，请删除该链表中重复的结点，重复的结点不保留，返回链表头指针。
// 例如，链表1->1->1->1->1->1->2->3->3->3->3->4->4->5 处理后为 1->2->5
func SortLinkListRepeat(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	currentNode := head
	var lastNode *ListNode

	for currentNode != nil {

		if currentNode.Next == nil {
			break
		} else {
			//当前元素值与下一个相等
			if currentNode.Val == currentNode.Next.Val {
				//遍历丢弃中间相同的元素
				for {
					if currentNode.Next != nil && currentNode.Val == currentNode.Next.Val {
						currentNode.Next = currentNode.Next.Next
					} else {
						break
					}
				}
				//头节点就有相同元素
				if lastNode == nil {
					lastNode = currentNode.Next
					head = lastNode
				} else {
					lastNode.Next = currentNode.Next
				}
				currentNode = currentNode.Next

			} else {
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

func IncreasingLinkMerge(head1 *ListNode, head2 *ListNode) *ListNode {

	var res *ListNode
	var resTail *ListNode
	leftCurrent := head1
	rightCurrent := head2

	for {
		if leftCurrent == nil {
			resTail.Next = rightCurrent
			break
		}
		if rightCurrent == nil {
			resTail.Next = leftCurrent
			break
		}

		if leftCurrent.Val <= rightCurrent.Val {

			if res == nil {
				res = leftCurrent
			} else {
				resTail.Next = leftCurrent
			}
			resTail = leftCurrent
			leftCurrent = leftCurrent.Next
		} else {

			if res == nil {
				res = rightCurrent
			} else {
				resTail.Next = rightCurrent
			}
			resTail = rightCurrent
			rightCurrent = rightCurrent.Next
		}
	}

	return res

}

//输入一个链表，输出该链表中倒数第k个结点。
//如果该链表长度小于k，请返回空。
//{1,2,3,4,5},1   {5}

func ReciprocalKNode(head *ListNode, k int) (res int, err error) {

	fast := 1
	fastNode := head
	slow := 1
	slowNode := head

	if head == nil {
		return 0, errors.New("empty link")
	}
	if k <= 0 {
		return 0, errors.New("k less 0")
	}
	//note fast <= k
	for fast = 1; fast < k; fast++ {
		if fastNode == nil {
			break
		} else {
			fastNode = fastNode.Next
		}

	}

	if fast < k {

		return 0, errors.New("length less then k")
	}

	for {
		if fastNode.Next == nil {
			break
		} else {
			slow++
			fast++
			slowNode = slowNode.Next
			fastNode = fastNode.Next
		}
	}

	if k == 1 {
		return fastNode.Val, nil
	} else {
		return slowNode.Val, nil
	}

}

//给一个链表，若其中包含环，请找出该链表的环的入口结点，否则，输出null。
//https://my.oschina.net/songjilong/blog/3159816
func EntryNodeOfLoop(head *ListNode) *ListNode {

	//1.遍历一遍有重复的就是环
	allNode := make(map[*ListNode]*ListNode)
	currentNode := head
	var res *ListNode
	for currentNode != nil {
		_, ok := allNode[currentNode]
		if ok == false {
			allNode[currentNode] = currentNode
			currentNode = currentNode.Next
		} else {
			res = currentNode
			break
		}
	}
	return res

	//2. 双指针
}

//输入两个无环的单链表，找出它们的第一个公共结点。
func FindFirstCommonNode(left *ListNode, right *ListNode) *ListNode {

	var res *ListNode
	allNode := make(map[*ListNode]*ListNode)

	leftNode := left
	rightNode := right

	for leftNode != nil && rightNode != nil {

		_, hasLeft := allNode[leftNode]
		if hasLeft {
			res = leftNode
			break
		} else {
			allNode[leftNode] = leftNode
			leftNode = leftNode.Next
		}
		_, hasRight := allNode[rightNode]
		if hasRight {
			res = rightNode
			break
		} else {
			allNode[rightNode] = rightNode
			rightNode = rightNode.Next
		}
	}

	return res

}

//输入一个链表，反转链表后，输出新链表的表头。
//{1,2,3} {3,2,1}
//{1,2,3,4,5}
//转成栈再转链表
//todo

//输入一个复杂链表（每个节点中有节点值，以及两个指针，一个指向下一个节点，另一个特殊指针random指向一个随机节点），
//请对此链表进行深拷贝，并返回拷贝后的头结点。（注意，输出结果中请不要返回参数中的节点引用，否则判题程序会直接返回空）
//todo

//todo 用两个栈来实现一个队列，完成队列的Push和Pop操作。 队列中的元素为int类型。

//todo 定义栈的数据结构，请在该类型中实现一个能够得到栈中所含最小元素的min函数（时间复杂度应为O（1））。

//todo 输入一棵二叉树，判断该二叉树是否是平衡二叉树。
//具有以下性质：它是一棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是一棵平衡二叉树。

//在一个字符串(0<=字符串长度<=10000，全部由字母组成)中找到第一个只出现一次的字符
//,并返回它的位置, 如果没有则返回 -1（需要区分大小写）.（从0开始计数）

//从上到下按层打印二叉树，同一层结点从左至右输出。每一层输出一行。
//{8,6,10,5,7,9,11}  [[8],[6,10],[5,7,9,11]]
func (root *BinaryTreeNode) Print() [][]int {

	// var res [][]int
	res := make([][]int, 0, 100)

	currentNodeList := []*BinaryTreeNode{root}
	for len(currentNodeList) > 0 {
		var nextNodeList []*BinaryTreeNode
		var currentList []int

		for _, node := range currentNodeList {
			currentList = append(currentList, node.Val)
			if node.LeftChild != nil {
				nextNodeList = append(nextNodeList, node.LeftChild)
			}
			if node.RightChild != nil {
				nextNodeList = append(nextNodeList, node.RightChild)
			}
		}
		res = append(res, currentList)
		currentList = currentList[0:0]

		currentNodeList = nextNodeList
		nextNodeList = nextNodeList[0:0]

	}
	return res
}

//给定一棵二叉搜索树，请找出其中的第k小的TreeNode结点。
//输入：{5,3,7,2,4,6,8},3 返回值：{4}
func KthNode(root *BinaryTreeNode, k int) *BinaryTreeNode {

	_, nodeList := root.InOrderByRecursive()

	if k > len(nodeList) {
		return nil
	}

	return nodeList[k-1]
}
