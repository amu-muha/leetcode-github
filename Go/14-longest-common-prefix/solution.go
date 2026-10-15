func longestCommonPrefix(strs []string) string { 
   pre:= ""
   for i:= 0; i< len(strs[0]);i++{
    for j:= 1; j< len(strs);j++{
        if i >= len(strs[j]) || strs[0][i] != strs[j][i]{
            return pre
        }
    }
    pre += string(strs[0][i])
   }
   return pre

}