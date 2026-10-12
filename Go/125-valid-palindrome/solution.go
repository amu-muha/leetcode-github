func isPalindrome(s string) bool {
   
    re := regexp.MustCompile("[^a-zA-Z0-9]")
    regex:= re.ReplaceAllString(s,"")

    result:= strings.ToLower(regex)
    

   for i,j := 0,len(result)-1;i<j;i,j=i+1,j-1{
    if result[i]!= result[j]{
        return false
    }
   }

    return true
}