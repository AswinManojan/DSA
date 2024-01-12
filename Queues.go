package main

import "fmt"

type Node struct {
   data interface{}
   next *Node
}

type Queue struct {
   front *Node
   rear  *Node
}

func (q *Queue) Enqueue(value interface{}) {
   newNode := &Node{data: value}
   if q.front == nil {
       q.front = newNode
       q.rear = newNode
   } else {
       q.rear.next = newNode
       q.rear = newNode
   }
}

func (q *Queue) Dequeue() {
   if q.front == nil {
       fmt.Println("empty queue")
   } else {
       q.front = q.front.next
   }
}

func (q *Queue) display() {
   if q.front == nil {
       fmt.Println("empty queue")
   } else {
       current := q.front
       for current != nil {
           fmt.Println(current.data)
           current = current.next
       }
   }
}

func main() {
   q := &Queue{}
   q.Enqueue(5)
   q.Enqueue(4)
   q.Enqueue(3)
   q.display()
   q.Dequeue()
   q.display()
   q.Dequeue()
   q.display()
   q.Dequeue()
   q.display()
   q.Dequeue()
}



225. Implement Stack using Queues

type MyStack struct {
   data []int
}


func Constructor() MyStack {
   return MyStack{}
}


func (this *MyStack) Push(x int)  {
   tmp := make([]int, len(this.data)+1)
   tmp[0] = x
   for i, v := range this.data {
       tmp[i+1] = v
   }
   this.data = tmp 
}


func (this *MyStack) Pop() int {
   x := this.data[0]
   this.data = this.data[1:]
   return x
}


func (this *MyStack) Top() int {
   return this.data[0]
}


func (this *MyStack) Empty() bool {
   return len(this.data) == 0
}



933. Number of recent calls

type RecentCounter struct {
   requests []int
}


func Constructor() RecentCounter {
   return RecentCounter{}
}


func (this *RecentCounter) Ping(t int) int {
   this.requests = append(this.requests, t)
   for this.requests[0] < t-3000 {
       this.requests = this.requests[1:]
   }
   return len(this.requests)
}

