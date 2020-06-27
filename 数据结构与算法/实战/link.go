package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func showNode(head *Node) {
	for head != nil {
		fmt.Println(head) // 代表的是链表中的node
		head = head.next  // 移动指针
	}
}

func reverseKGroup(head *Node, k int) *Node {
	// 借助数组把链表先进行分组
	groupList, group := make([][]*Node, 0), make([]*Node, 0)
	count := 0
	tmphead := head
	for tmphead != nil {
		// 给每组创建一个数组
		if count%k == 0 { // k = 2, count = 0,2,4,....
			group = make([]*Node, 0)
		}
		group = append(group, tmphead)

		// 当每个组达到了k个或者下一个为nil了，放到groupList里面去
		if len(group) == k {
			// 翻转一下数组
			newgroup := make([]*Node, 0)
			for i := k - 1; i >= 0; i-- {
				newgroup = append(newgroup, group[i])
			}
			groupList = append(groupList, newgroup)
			fmt.Printf("%+v", groupList)
		} else if tmphead.next == nil {
			// 不需要翻转
			groupList = append(groupList, group)
			fmt.Println("=========")
		}

		// 移动指针 计数
		tmphead = tmphead.next
		count++
	}
	fmt.Printf("%+v", groupList)

	// 把分组后的链表串起来即可(上面已经排序好了)
	rtn := &Node{data: 0, next: head}
	pre := rtn                           // 这个用来改变指针指向
	for _, onegroup := range groupList { // n/k * O(n)
		total := len(onegroup)
		for i := 0; i < total; i++ {
			pre.next = onegroup[i]
			pre = pre.next // 移动
		}
	}
	return rtn.next
}

func main() {
	var head = new(Node)
	head.data = 1
	var node1 = new(Node)
	node1.data = 2
	head.next = node1

	showNode(head)

	reverseKGroup(head, 2)

}
