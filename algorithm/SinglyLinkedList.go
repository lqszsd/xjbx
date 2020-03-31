package main

import (
	"fmt"
	"sync"
)

type SingleObject interface {

}
type SingleNode struct {
	Data SingleObject
	Next *SingleNode
}
type SingleList struct {
	mutex *sync.RWMutex
	Head *SingleNode
	Tail *SingleNode
	Size uint
}
//链表数据结构初始化
func (list *SingleList)Init()  {
	list.Size=0
	list.Head=nil
	list.Tail=nil
	list.mutex=new(sync.RWMutex)
}
//添加节点
func (list *SingleList)Add(node *SingleNode) bool  {
	if node==nil{
		return false
	}
	//并发锁
	list.mutex.Lock()
	//执行结束后解锁
	defer list.mutex.Unlock()
	if list.Size==0{
		list.Head=node
		list.Tail=node
		list.Size=1
		return true
	}
	tail:=list.Tail
	tail.Next=node
	//当前节点头指向下一个节点
	list.Tail=node
	list.Size+=1
	return true
}
//指定指针 添加节点
func (list *SingleList)Insert(index uint,node *SingleNode) bool{
	if node==nil{
		return false
	}
	//如果索引长度大于记录的链表长度
	if index > list.Size{
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if index==0{
		node.Next=list.Head
		list.Head=node
		list.Size+=1
		return true
	}
	var i uint
	ptr :=list.Head
	for i=1;i<index;i++{
		ptr=ptr.Next
	}
	next:=ptr.Next
	ptr.Next=node
	node.Next=next
	list.Size+=1
	return true
}
//删除节点
func(list *SingleList)Delete(index uint) bool{
	if list==nil||list.Size==0||index>list.Size-1{
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if index==0{
		head:=list.Head.Next
		list.Head=head
		if list.Size==1{
			list.Tail=nil
		}
		list.Size-=1
		return true
	}
	ptr:=list.Head
	var i uint
	for i=1;i<index;i++{
		ptr=ptr.Next
	}
	next:=ptr.Next
	ptr.Next=next.Next
	if index == list.Size - 1 {
		list.Tail = ptr
	}
	list.Size -= 1
	return true
}
func (list *SingleList)Get(index uint)*SingleNode{
	if list == nil || list.Size == 0 || index > list.Size - 1 {
		return nil
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()

	if index == 0{
		return list.Head
	}
	node := list.Head
	var i uint
	for i = 0; i < index; i ++ {
		node = node.Next
	}
	return node
}
// 输出链表
func (list *SingleList)Display(){
	if list == nil {
		fmt.Println("this single list is nil")
		return
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	fmt.Printf("this single list size is %d \n", list.Size)
	ptr := list.Head
	var i uint
	for i = 0; i < list.Size; i++{
		fmt.Printf("No%3d data is %v\n", i + 1, ptr.Data)
		ptr = ptr.Next
	}
}
func main(){
	var s=SingleList{}
	s.Init()
	for i:=0;i<10;i++{
		s.Add(&SingleNode{Data:1})
		s.Insert(1, &SingleNode{Data:2})
	}
	s.Display()
}