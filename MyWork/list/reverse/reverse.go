package reverse

type ListNode struct {
	Value int
	Next  *ListNode
}

func Reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode

	for head != nil {
		temp := head.Next
		head.Next = pre
		pre = head
		head = temp
	}

	return pre
}

func InitHeadList(data []int) *ListNode {
	var head = new(ListNode)
	head.Value = 0
	var tail *ListNode
	tail = head //tail用于记录头结点的地址，刚开始tail的的指针指向头结点
	for i := 1; i < 10; i++ {
		var node = ListNode{Value: i}
		node.Next = tail //将新插入的node的next指向头结点
		tail = &node     //重新赋值头结点
	}
	return tail
}

func InitTailList(data []int) *ListNode {
	var head = new(ListNode)
	head.Value = 0
	var tail *ListNode
	tail = head //tail用于记录最末尾的结点的地址，刚开始tail的的指针指向头结点
	for i := 1; i < 10; i++ {
		var node = ListNode{Value: i}
		(*tail).Next = &node
		tail = &node
	}
	return head
}
