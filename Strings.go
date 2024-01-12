// function to replace each alphabet in the given string with another alphabet occurring 
// at the n-th position from each of them.

func Encode(s string,n int) string{
   Result := ""
   Change := n % 26
   for i := 0 ; i < len(s) ; i++ {
       if s[i] <= 122-change{
           result += string(s[i]+byte(change))
       }else{
		result += string(s[i]-26+byte(change))
}
   }
   return result
}





// Sample workouts

    1. Reveresing a string
func main(){
   s:="hello"
   reversed:=ReverseString(s)
   fmt.Println(reversed)
}


func ReverseString(s string) string {
   result:=""
   for i:=len(s)-1;i>=0;i--{
       result+=string(s[i])
   }
   return result
}

    // 2. Check whether string is palindrome
			
		 
func main(){
   s:="malayalam"
   if IsPallindrome(s){
       fmt.Println("it is pallindrome")
   }else{
       fmt.Println("Not pallindrome")
   }
}


func IsPallindrome(s string) bool{
   for i:=0;i<(len(s)+1)/2;i++{
       if s[i]!=s[len(s)-i-1]{
           return false
       }
   }
   return true
}

		
    // 3. Count vowels in a string
func main(){
   s:="hei"
   count:=CountVowels(s)
   fmt.Println(count)

}
func CountVowels(s string) int {
   count:=0
   for i:=0;i<len(s);i++{
       if IsVowel(s[i]){
           count++
       }
   }
   return count
}

func IsVowel(s byte) bool{
   return s=='a' || s=='e' || s== 'i' || s=='o' || s=='u'
}