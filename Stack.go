package main

import "fmt"

type Node struct {
   data interface{}
   next *Node
}
type Stack struct {
   top *Node
}

func (s *Stack) push(value interface{}) {
   newNode := &Node{data: value}
   if s.top == nil {
       s.top = newNode
   } else {
       newNode.next = s.top
       s.top = newNode
   }
}

func (s *Stack) pop() {
   if s.top == nil {
       fmt.Println("Stack underflow")
   } else {
       s.top = s.top.next
   }
}
func (s *Stack) display() {
   if s.top == nil {
       fmt.Println("Stack underflow")
   } else {
       current := s.top
       for current != nil {
           fmt.Println(current.data)
           current = current.next
       }
   }
}

func main() {
   stk := &Stack{}
   stk.push(1)
   stk.push(22)
   stk.push(21)
   stk.display()
   stk.pop()
   stk.display()
   stk.pop()
   stk.display()
   stk.pop()
   stk.display()
}



682. Baseball Game

func calPoints(operations []string) int {
   stack:=&Stack{}
   for _,v:= range(operations){
       switch v{
           case "+":
               stack.Addtwo()
           case "C":
               stack.Pop()
           case "D":
               stack.Double()
           default:
               val,_:=strconv.Atoi(v)
               stack.Push(val)
       }
   }
   return stack.FullSum()
}

type Node struct{
   value int
   next *Node
}

type Stack struct{
   top *Node
}

func (s *Stack) Pop(){
   s.top=s.top.next
}
func (s *Stack) Push(data int){
   newNode:=&Node{
       value:data,
   }
   if s.top==nil{
       s.top=newNode
   }else{
       newNode.next=s.top
       s.top=newNode
   }
}
func (s *Stack) Double(){
   data:= s.top.value
   newNode:=&Node{
       value:data*2,
   }
   if s.top==nil{
       s.top=newNode
   }else{
       newNode.next=s.top
       s.top=newNode
   }
}
func (s *Stack) Addtwo(){
   val:= s.top.value + s.top.next.value
   newNode:=&Node{
       value:val,
   }
   if s.top==nil{
       s.top=newNode
   }else{
       newNode.next=s.top
       s.top=newNode
   }
}
func (s *Stack) FullSum() int{
   current:=s.top
   sum:=0
   for current!=nil{
       sum+=current.value
       current=current.next
   }
   return sum
}


20. Valid Parenthesis

func isValid(s string) bool {
   var st Stack
   for _,b := range (s){
       v:=byte(b)
       if v=='[' ||v=='(' ||v=='{' {
           st.Push(v)
       }else{
           if st.top==nil{
               return false
           }else if v==']' && st.top.data=='['{
               st.Pop()
           }else if v==')' && st.top.data=='('{
               st.Pop()
           }else if v=='}' && st.top.data=='{'{
               st.Pop()
           }else{
               return false
           }
       }
   }
  return st.top==nil
}

type Node struct{
   data byte
   next *Node
}

type Stack struct{
   top *Node
}

func(s *Stack) Push(val byte){
   newNode:= Node{data: val,next: s.top}
   s.top=&newNode
}

func (s *Stack) Pop(){
   s.top=s.top.next
}

