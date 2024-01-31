package linkedList

import "fmt"

// 鏈結練習
// https://ithelp.ithome.com.tw/articles/10241532

type Node struct{
    Val interface{}
    Next *Node
}

type LinkedList struct{
    head *Node    // 指向開始的 Node
    tail *Node    // 指向結束的 Node
}

func New() *LinkedList{
    return new(LinkedList)
}

func NewNode(val interface{}) *Node{
    n := new(Node)
    n.Val = val
    return n
}

func (l *LinkedList) isEmpty() bool{
    return l.head == nil
}

func (l *LinkedList) Prepend(n *Node){
    if l.isEmpty(){
        l.head = n
        l.tail = n
        return
    }
    n.Next = l.head     // 將新節點的下個點指向原串列的頭
    l.head = n          // 將開頭改成新節點
}

func (l *LinkedList) Append(n *Node){
    if l.isEmpty(){
        l.head = n
        l.tail = n
        return
    }
    l.tail.Next = n     // 將結尾的下個點指向新節點
    l.tail = n          // 將結尾改成新節點
}

func (l *LinkedList) Head() *Node{
    return l.head
}

func (l *LinkedList) Tail() *Node{
    return l.tail
}

func (l *LinkedList) PopFirst() (*Node, error){
    // 串列為空 (噴錯)
    if l.isEmpty(){
        return nil, fmt.Errorf("Linked-List is empty")
    }
    // 串列只有一個值
    if l.head == l.tail{
        l.tail = nil     // 恢復成預設值
    }

    tmp := l.head
    l.head = l.head.Next // 將開頭指向第二個節點
    return tmp, nil
}

func RunLL(){
    myLikedList := New()
    myLikedList.Append(NewNode("小櫻"))
    myLikedList.Append(NewNode("小狼"))
    myLikedList.Prepend(NewNode("知世"))
    print(myLikedList)
}

func print(l *LinkedList){
    for current := l.Head(); current != nil; current = current.Next{
        fmt.Printf("%v -> ", current.Val)
    }
    fmt.Println("nil")
}
