package main

import "fmt"

// / Sorting an array
func main() {
	arr := []int{5, 1, 2, 7, 9, 3}
	fmt.Println(sorter(arr))
}

func sorter(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}

// Delete an element of an array
func main() {
   arr := []int{5, 1, 2, 7, 9, 3}
   x := 1
   fmt.Println(delete(arr, x))
}

func delete(arr []int, x int) []int {
   for i := 0; i < len(arr); i++ {
       if arr[i] == x {
           for j := i; j < len(arr)-1; j++ {
               arr[j] = arr[j+1]
           }
       }
   }
   return arr
}

// Find the largest element of an array
func main() {
   arr := []int{5, 1, 2, 7, 9, 3}
   fmt.Println(LargestElement(arr))
}

func LargestElement(arr []int) int {
   largest := arr[0]
   for i := 0; i < len(arr); i++ {
       if largest < arr[i] {
           largest = arr[i]
       }
   }
   return largest
}

// LEETCODE-283

func moveZeroes(nums []int)  {
   var ind,i int
   ind=0
   for i=0;i<len(nums);i++{
       if nums[i]!=0{
           nums[i],nums[ind]=nums[ind],nums[i]
           ind++
       }
   }
}

// LEETCODE - 217

func containsDuplicate(nums []int) bool {
   for i:=0;i<len(nums);i++{
       for j:=i+1;j<len(nums);j++{
           if(nums[i]==nums[j]){
               return true
           }
       }
   }
   return false
}

// LEETCODE - 66

func plusOne(digits []int) []int {
   for i:=len(digits)-1;i>=0;i--{
       if digits[i]!=9{
           digits[i]++
           break
       }else if digits[i]==9 && i==0{
           digits[0]=1
           digits=append(digits,0)
       }else if digits[i]==9{
           digits[i]=0
       }
   }
   return digits
}
