package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	value int
}

type BST struct {
	root *Node
}

func (b *BST) Insert(data int) {
	newNode := &Node{value: data}
	if b.root == nil {
		b.root = newNode
		return
	} else {
		current := b.root
		for {
			if current.value < data {
				if current.right == nil {
					current.right = newNode
					return
				}
				current = current.right
			} else if current.value > data {
				if current.left == nil {
					current.left = newNode
					return
				}
				current = current.left
			} else {
				return
			}
		}
	}
}
func (b *BST) InOrder() []int {
	arr := []int{}
	b.InOrderHelper(&arr, b.root)
	return arr
}
func (b *BST) InOrderHelper(arr *[]int, node *Node) {
	if node != nil {
		b.InOrderHelper(arr, node.left)
		*arr = append(*arr, node.value)
		b.InOrderHelper(arr, node.right)
	}
}

func (b *BST) PreOrder() []int {
	arr := []int{}
	b.PreOrderHelper(&arr, b.root)
	return arr
}
func (b *BST) PreOrderHelper(arr *[]int, node *Node) {
	if node != nil {
		fmt.Printf("%d ", node.value)
		b.PreOrderHelper(arr, node.left)
		b.PreOrderHelper(arr, node.right)
	}
}

func (b *BST) PostOrder() {
	arr := []int{}
	b.PostOrderHelper(&arr, b.root)
}
func (b *BST) PostOrderHelper(arr *[]int, node *Node) {
	if node != nil {
		b.PostOrderHelper(arr, node.left)
		b.PostOrderHelper(arr, node.right)
		fmt.Printf("%d ", node.value)
	}
}

func (b *BST) Contains(data int) bool {
	if b.root == nil {
		return false
	}
	current := b.root
	for current != nil {
		if data > current.value {
			current = current.right
		} else if data < current.value {
			current = current.left
		} else {
			return true
		}
	}
	return false
}
func (b *BST) GetMin(current *Node) int {
	for current.left != nil {
		current = current.left
	}
	return current.value
}
func (b *BST) Remove(value int) {
	b.RemoveHelper(value, nil, b.root)
}
func (b *BST) RemoveHelper(value int, parent, current *Node) {
	for current != nil {
		if value < current.value {
			parent = current
			current = current.left
		} else if value > current.value {
			parent = current
			current = current.right
		} else {
			if current.left != nil && current.right != nil {
				current.value = b.GetMin(current.right)
				b.RemoveHelper(current.value, current, current.right)
			} else {
				if parent == nil {
					if current.right == nil {
						b.root = current.left
					} else {
						b.root = current.right
					}
				} else {
					if parent.left == current {
						if current.left == nil {
							parent.left = current.right
						} else {
							parent.left = current.left
						}
					} else {
						if current.right == nil {
							parent.right = current.left
						} else {
							parent.right = current.right
						}
					}

				}
			}
			return
		}
	}
	fmt.Println("Not a value in tree")
}

func (b *BST) isValid() bool {
	arr := b.InOrder()
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}
	return true
}

func (b *BST) FindHeight(node *Node) int {
	if node == nil {
		return -1
	}
	left := b.FindHeight(node.left)
	right := b.FindHeight(node.right)
	max := 0
	if left > right {
		max = left
	} else {
		max = right
	}
	return max + 1
}

var flag int

func (b *BST) IsBalanced(node *Node) bool {
	b.IsBalancedHelper(node)
	return flag == 1
}
func (b *BST) IsBalancedHelper(node *Node) int {
	if node == nil {
		return -1
	}
	left := b.IsBalancedHelper(node.left)
	right := b.IsBalancedHelper(node.right)

	if left-right > 1 || right-left > 1 {
		flag = 0
	}
	max := 0
	if left > right {
		max = left
	} else {
		max = right
	}
	return max + 1
}

func main() {
	bst := &BST{}
	bst.Insert(25)
	bst.Insert(10)
	bst.Insert(35)
	bst.Insert(5)
	bst.Insert(20)
	bst.Insert(40)
	bst.Insert(3)
	bst.Insert(6)
	bst.Insert(15)
	bst.Insert(21)
	bst.Insert(16)

	fmt.Println(bst.FindHeight(bst.root))
	arr1 := bst.InOrder()
	fmt.Println(arr1)
	bst.Remove(0)
	fmt.Println(bst.Contains(1))
	arr := bst.InOrder()
	fmt.Println(arr)
	fmt.Println(bst.isValid())
	bst.PreOrder()
	fmt.Printf("\n")
	bst.PostOrder()
}
