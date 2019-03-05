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
    size uint64
    head *Node
    tail *Node
}

// 初始化链表
func (list *List) Init() {
    list.size = 0
    list.head = nil
    list.tail = nil
}

// 添加元素方法 参数添加的节点即是头部也是尾部
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

func main() {
    fmt.Println("hello")
}
