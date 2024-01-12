package main

import "fmt"

type ListNode struct {
	value interface{}
	next  *ListNode
}

type LinkedList struct {
	head *ListNode
	size int
}

func (ll *LinkedList) append(data interface{}) {
	newNode := &ListNode{value: data}
	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	ll.size++
}
func (ll *LinkedList) prepend(data interface{}) {
	newNode := &ListNode{value: data}
	if ll.head == nil {
		ll.head = newNode
	} else {
		newNode.next = ll.head
		ll.head = newNode
	}
	ll.size++
}
func arrayToLL(arr []int) *LinkedList {
	llist := &LinkedList{}
	for _, v := range arr {
		llist.append(v)
	}
	return llist
}
func (ll *LinkedList) display() {
	if ll.head == nil {
		fmt.Println("No element in the list.")
	} else {
		current := ll.head
		for current != nil {
			fmt.Printf("%v ->", current.value)
			current = current.next
		}
	}
}

func (ll *LinkedList) deletebeforeX(x interface{}) {
	if ll.head == nil {
		fmt.Println("No element in the list.")
	} else {
		current := ll.head
		flag := false
		for current.next.next != nil {
			if current.next.next.value == x {
				flag = true
				break
			}
			current = current.next
		}
		if flag {
			current.next = current.next.next
			ll.size--
		} else {
			fmt.Printf("\nValue not present.")
		}
	}
}

func (ll *LinkedList) deleteAfterX(x interface{}) {
	if ll.head == nil {
		fmt.Println("No element in the list.")
	} else {
		current := ll.head
		flag := false
		for current.next != nil {
			if current.value == x {
				flag = true
				break
			}
			current = current.next
		}
		if flag {
			current.next = current.next.next
			ll.size--
		} else {
			fmt.Printf("\nValue not present or in last position")
		}
	}
}

func (ll *LinkedList) insertAfterX(x interface{}, data interface{}) {
	if ll.head == nil {
		fmt.Println("No element in the list.")
	} else {
		newNode := &ListNode{value: data}
		current := ll.head
		flag := false
		for current.next != nil {
			if current.value == x {
				flag = true
				break
			}
			current = current.next
		}
		if flag {
			newNode.next = current.next
			current.next = newNode
			ll.size++
		} else {
			fmt.Printf("\nValue not present or in last position")
		}
	}
}

func (ll *LinkedList) insertBeforeX(x interface{}, data interface{}) {
	if ll.head == nil {
		fmt.Println("No element in the list.")
	} else {
		newNode := &ListNode{value: data}
		current := ll.head
		flag := false
		for current.next != nil {
			if current.next.value == x {
				flag = true
				break
			}
			current = current.next
		}
		if flag {
			newNode.next = current.next
			current.next = newNode
			ll.size++
		} else {
			fmt.Printf("\nValue not present or in last position")
		}
	}
}

func (ll *LinkedList) reverseLinkedList() {
	if ll.head == nil {
		fmt.Println("No element in the list.")
	} else {
		current := ll.head
		var prev *ListNode
		for current != nil {
			nxt := current.next
			current.next = prev
			prev = current
			current = nxt
		}
		ll.head = prev
	}
}

func main() {
	arr := []int{1, 2, 4, 5, 6}
	ll := arrayToLL(arr)
	fmt.Printf("Size is %v\n", ll.size)
	ll.display()
	ll.prepend(99)
	fmt.Printf("\nSize is %v\n", ll.size)
	ll.display()
	ll.deletebeforeX(4)
	fmt.Printf("\nSize is %v\n", ll.size)
	ll.display()
	ll.deleteAfterX(6)
	fmt.Printf("\nSize is %v\n", ll.size)
	ll.display()
	ll.insertAfterX(4, 44)
	fmt.Printf("\nSize is %v\n", ll.size)
	ll.display()
	ll.insertBeforeX(6, 66)
	fmt.Printf("\nSize is %v\n", ll.size)
	ll.display()
	ll.reverseLinkedList()
	fmt.Printf("\nSize is %v\n", ll.size)
	ll.display()
}
