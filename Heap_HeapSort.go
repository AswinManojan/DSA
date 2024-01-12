package main

import "fmt"

type MaxHeap struct {
	arr []int
}

func (m *MaxHeap) Insert(data int) {
	m.arr = append(m.arr, data)
	current := len(m.arr) - 1
	m.ShiftUp(current)
}

func (m *MaxHeap) ShiftUp(current int) {
	parent := m.Parent(current)
	for current > 0 && m.arr[current] > m.arr[parent] {
		m.arr[current], m.arr[parent] = m.arr[parent], m.arr[current]
		current = parent
		parent = m.Parent(current)
	}
}

func (m *MaxHeap) Remove() {
	m.arr[0], m.arr[len(m.arr)-1] = m.arr[len(m.arr)-1], m.arr[0]
	m.arr = m.arr[:len(m.arr)-1]
	m.ShiftDown(0)
}

func (m *MaxHeap) ShiftDown(current int) {
	for {
		largest := current
		left := m.LeftChild(largest)
		right := m.RightChild(largest)
		if len(m.arr) > left && m.arr[largest] < m.arr[left] {
			largest = left
		}
		if len(m.arr) > right && m.arr[largest] < m.arr[right] {
			largest = right
		}
		if largest == current {
			return
		}
		m.arr[largest], m.arr[current] = m.arr[current], m.arr[largest]
		current = largest
	}
}
func (m *MaxHeap) LeftChild(i int) int {
	return (i * 2) + 1
}
func (m *MaxHeap) RightChild(i int) int {
	return (i * 2) + 2
}
func (m *MaxHeap) Parent(i int) int {
	return (i - 1) / 2
}
func (m *MaxHeap) Heapify(arr []int) []int {
	for _, x := range arr {
		m.Insert(x)
	}
	return m.arr
}
func (m *MaxHeap) Display() {
	for _, x := range m.arr {
		fmt.Printf("%d ", x)
	}
}

func (m *MaxHeap) HeapSort(arr []int) {
	var sortedArr []int
	m.Heapify(arr)
	for len(m.arr) > 0 {
		sortedArr = append([]int{m.arr[0]}, sortedArr...)
		m.Remove()
	}
	fmt.Println("Sorted array:", sortedArr)
}
func prepend(arr []int, elem int) []int {
	return append([]int{elem}, arr...)
}
func main() {
	h := &MaxHeap{}
	// h.Insert(5)
	// h.Insert(3)
	// h.Insert(9)
	// h.Insert(6)
	// h.Insert(2)
	// h.Insert(1)
	// h.Insert(7)
	// h.Insert(8)
	// h.Insert(10)
	// h.Display()

	// h.Remove()
	// fmt.Printf("\n")
	// h.Display()

	arr := []int{5, 3, 8, 9, 12, 1, 3, 6, 12, 2}
	// h.Heapify(arr)
	h.HeapSort(arr)
	// h.Display()
	// fmt.Println(arr)

}
