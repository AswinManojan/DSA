package main

import "fmt"

func main() {
	arr := []int{11, 2, 33, 4, 55, 6}
	QuickSort(arr)
	fmt.Println(arr)
}

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	return arr
}

func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

func insertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}
func merge(left []int, right []int) []int {
	combined := make([]int, len(left)+len(right))
	index, i, j := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			combined[index] = left[i]
			i++
			index++
		} else {
			combined[index] = right[j]
			j++
			index++
		}
	}
	for i < len(left) {
		combined[index] = left[i]
		index++
		i++
	}
	for j < len(right) {
		combined[index] = right[j]
		index++
		j++
	}
	return combined
}

func Pivot(arr []int, pivotIndex int, endIndex int) int {
	swapIndex := pivotIndex
	for i := pivotIndex + 1; i <= endIndex; i++ {
		if arr[i] < arr[pivotIndex] {
			swapIndex++
			arr[swapIndex], arr[i] = arr[i], arr[swapIndex]
		}
	}
	arr[pivotIndex], arr[swapIndex] = arr[swapIndex], arr[pivotIndex]
	return swapIndex
}
func QuickSortHelper(arr []int, left int, right int) {
	if left < right {
		pivotIndex := Pivot(arr, left, right)
		QuickSortHelper(arr, left, pivotIndex-1)
		QuickSortHelper(arr, pivotIndex+1, right)
	}
}
func QuickSort(arr []int) {
	QuickSortHelper(arr, 0, len(arr)-1)
}
