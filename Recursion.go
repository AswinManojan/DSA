
// Factorial using recursive function

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// Binary search using recursive function

func BinarySearchR(arr []int, target, low, high int) int {
	mid := (low + high) / 2
	if low <= high {
		if arr[mid] == target {
			return mid
		} else if target < arr[mid] {
			BinarySearchR(arr, target, low, mid-1)
		} else {
			BinarySearchR(arr, target, mid+1, high)
		}
	}
	return mid
}

// The nth term of fibonacci series using recursion

func FibonacciSeries(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciSeries(n-1) + FibonacciSeries(n-2)
}

