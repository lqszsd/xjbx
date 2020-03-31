package main

import (
	"fmt"
	"sync"
)

type DoubleObject interface {

}
type DoubleNode struct {
	Data DoubleObject
	Prev *DoubleNode
	Next *DoubleNode
}
type DoubleList struct {
	mutex *sync.RWMutex
	Size uint
	Head *DoubleNode
	Tail *DoubleNode
}
func (list *DoubleList)New(){
	list.mutex=new(sync.RWMutex)
	list.Size=0
	list.Head=nil
	list.Tail=nil
}
func (list *DoubleList)Get(index uint)*DoubleNode{
	if list.Size==0||index>list.Size-1{
		return nil
	}
	if index==0{
		return list.Head
	}
	node:=list.Head
	var i uint
	for i=0;i<index;i++{
		node=node.Next
	}
	return node
}
func (list *DoubleList)Append(node *DoubleNode)bool{
	if node==nil{
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if list.Size==0{
		//头节点和尾节点都指向当前节点
		list.Head=node
		list.Tail=node
		node.Next=nil
		node.Prev=nil
	} else{
		//插入节点的前一个节点等于上一个节点
		node.Prev=list.Tail
		//插头结点的下一个和节点 等于开头的nil节点
		node.Next=nil
		list.Tail.Next=node
		list.Tail=node
	}
	list.Size++
	return true
}
func (list *DoubleList)Insert(index uint,node *DoubleNode)bool {
	if index>list.Size||node==nil{
		return false
	}
	if index==list.Size{
		return list.Append(node)
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if index==0{
		//当前节点的下一个节点 等于当前头节点
		node.Next=list.Head
		list.Head=node
		list.Head.Prev=nil
		list.Size++
		return true
	}
	nextNode:=list.Get(index)
	node.Prev=nextNode.Prev
	node.Next=nextNode
	nextNode.Prev.Next=node
	nextNode.Prev=node
	list.Size++
	return true
}
func (list *DoubleList) Delete (index uint) bool {
	if index > list.Size - 1 {
		return false
	}

	list.mutex.Lock()
	defer list.mutex.Unlock()
	if index == 0 {
		if list.Size == 1{
			list.Head = nil
			list.Tail = nil
		} else {
			list.Head.Next.Prev = nil
			list.Head = list.Head.Next
		}
		list.Size--
		return true
	}
	if index == list.Size - 1{
		list.Tail.Prev.Next = nil
		list.Tail = list.Tail.Prev
		list.Size--
		return true
	}

	node := list.Get(index)
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	list.Size--
	return true
}
func (list *DoubleList)Display(){
	if list == nil || list.Size == 0 {
		fmt.Println("this double list is nil or empty")
		return
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	fmt.Printf("this double list size is %d \n", list.Size)
	ptr := list.Head
	for ptr != nil {
		fmt.Printf("data is %v\n", ptr.Data)
		ptr = ptr.Next
	}
}

// Reverse 倒序打印双链表信息
func (list *DoubleList)Reverse(){
	if list == nil || list.Size == 0 {
		fmt.Println("this double list is nil or empty")
		return
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	fmt.Printf("this double list size is %d \n", list.Size)
	ptr := list.Tail
	for ptr != nil {
		fmt.Printf("data is %v\n", ptr.Data)
		ptr = ptr.Prev
	}
}