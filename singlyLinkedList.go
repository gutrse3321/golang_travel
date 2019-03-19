// Singly Linked List 单链表
package main

import (
    "fmt"
)

type Object interface{}

// 节点 data值 next链节
type Node struct {
    data Object
    next *Node
}

// 链表 size长度 head首个 tail尾部
type List struct {
    size int
    head *Node
    tail *Node
}

// 初始化链表
func (list *List) Init() {
    list.size = 0
    list.head = nil
    list.tail = nil
}

// 添加元素方法 添加到最后节点最里部
func (list *List) Append(node *Node) bool {
    if node == nil {
        return false
    }

    node.next = nil

    if list.size == 0 {
        list.head = node
    } else {
        oldTail := list.tail
        oldTail.next = node
    }

    list.tail = node
    list.size++

    return true
}

// 插入到链表中(插入到尾部的话可以直接用Append，故这里不做处理)
func (list *List) Insert(i int, node *Node) bool {
    // 空间点、超出索引大小和空链表无法插入
    if node == nil || i > list.size || list.size == 0 {
        return false
    }

    if i == 0 {
        node.next = list.head
        list.head = node
    } else {
        // 上一个、前一个节点
        prev := list.head
        for j := 1; j < i; j++ {
            prev = prev.next
        }
        node.next = prev.next
        prev.next = node
    }
    list.size++

    return true
}

// 从链表中删除
func (list *List) Delete(i int) bool {
    if i > list.size || list.size == 0 {
        return false
    }
    
    if i == 0 {
        list.head = list.head.next
    } else {
        prev := list.head
        for j := 1; j < i; j++ {
            prev = prev.next
        }
        next := prev.next.next
        prev.next = next
        
    }
    list.size--
    
    return true
}

func main() {
    list := &List{}
    list.Init()
    node := &Node{
        data: 1,
        next: nil,
    }
    list.Append(node)
    fmt.Println(list)
    node2 := &Node{
        data: 2,
        next: nil,
    }
    list.Append(node2)
    fmt.Println(list)
    
    node3 := &Node{
        data: 3,
        next: nil,
    }
    list.Append(node3)
    
    node4 := &Node{
        data: 4,
        next: nil,
    }
    list.Insert(1, node4)
    
    list.Delete(1)
    
    fmt.Println(list.head.data)
    fmt.Println(list.head.next.data)
    fmt.Println(list.head.next.next.data)
    // fmt.Println(list.head.next.next.next.data)
    // fmt.Println(list.tail.data)
    fmt.Println(list.tail.next)
}
