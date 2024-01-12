package main

import (
"errors"
"fmt"
)

type Node struct {
Value int
Key int
Next *Node
}

type HashMap struct {
Table []*Node
Size int
}

func NewHashMap(size int) *HashMap {
return &HashMap{
Table: make([]*Node, size),
Size: size,
}
}

func (hm *HashMap) hashFunction(key int) int {
return key % hm.Size
}

func (hm *HashMap) Put(key, val int) {
index := hm.hashFunction(key)
newNode := &Node{Value: val,
Key: key}
if hm.Table[index] == nil {
hm.Table[index] = newNode
} else {
newNode.Next = hm.Table[index]
hm.Table[index] = newNode
}
}

func (hm *HashMap) Get(key int) (int, error) {
index := hm.hashFunction(key)
current := hm.Table[index]
for current != nil {
if current.Key == key {
return current.Value, nil
}
current = current.Next
}
return 0, errors.New("Key not found.")
}

func (hm *HashMap) Remove(key int) {
index := hm.hashFunction(key)
current := hm.Table[index]
var prev *Node
for current != nil {
if current.Key == key {
if prev == nil {
hm.Table[index] = current.Next
} else {
prev.Next = current.Next
}
return
}
prev = current
current = current.Next
}
}

func main() {
hm := NewHashMap(5)
hm.Put(1, 2)
hm.Put(3, 8)
hm.Put(5, 0)

val1, err1 := hm.Get(1)
val2, err2 := hm.Get(3)
val3, err3 := hm.Get(6)
val4, err4 := hm.Get(5)

fmt.Println("Value for key 1:", val1, "Error:", err1)
fmt.Println("Value for key 3:", val2, "Error:", err2)
fmt.Println("Value for key 6:", val3, "Error:", err3)
fmt.Println("Value for key 5:", val4, "Error:", err4)

hm.Remove(3)
val5, err5 := hm.Get(3)
fmt.Println("Value for key 3:", val5, "Error:", err5)
}





242. Valid Anagram

func isAnagram(s string, t string) bool {
   mp:=make(map[rune]int)
   for _,v:= range(s){
       mp[v]++
   }
   for _,v:=range(t){
       mp[v]--
   }
   for _,v:=range(mp){
       if v!=0{
           return false
       }
   }
   return true
}

409. Longest Palindrome

func longestPalindrome(s string) int {
   mp:=make(map [rune]int)
   for _,v:= range(s){
       mp[v]++
   }
   result:=0
   flag:=false
   for _,v:=range mp{
       if v%2==0{
           result+=v
       }else{
           result+=(v-1)
           flag=true
       }
   }
   if flag==false{
       return result
   }
   return result+1
}


1796. Second Largest Digit in a String

func secondHighest(s string) int {
   var temp []int
   d:=0
   myMap:=make(map[int]bool)
   for i:=0;i<len(s);i++{
       if s[i]>=48 && s[i]<=57{
           d=int(s[i]-48)
           if myMap[d]==false{
               temp=append(temp,d)
               myMap[d]=true
           }
       }
   }
   sort.Ints(temp)
   if len(temp)<2{
       return -1
   }
   return temp[len(temp)-2]
}


2032. Two out of Three

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
   var result []int
   myMap:=make(map[int]bool)
   for i:=0;i<len(nums1);i++{
       for j:=0;j<len(nums2);j++{
           if nums1[i]==nums2[j]{
               myMap[nums1[i]]=true
           }
           for k:=0;k<len(nums3);k++{
               if nums1[i]==nums3[k]||nums2[j]==nums3[k]{
                   myMap[nums3[k]]=true
               }
           }
       }
   }
   for key,value:=range myMap{
       if value{
           result=append(result,key)
       }
   }
   return result
}


2085. Count Common words with One Occurance

func countWords(words1 []string, words2 []string) int {
   m:= make(map[string]int)
   m2:=make(map[string]int)
   for _,v:=range(words1){
       m[v]++
   }
   for _,v:=range(words2){
       m2[v]++
   }
   var res int
   for k,v := range m{
       if v==1 && m2[k]==1{
           res++
       }
   }
   return res
}
