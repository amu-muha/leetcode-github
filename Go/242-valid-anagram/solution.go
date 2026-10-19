func isAnagram(s string, t string) bool {
    sum:= make(map[rune]int)
    
    for _,char := range s{
        sum[char]++
    }

    for _,char:= range t{
        sum[char]--
    }

    for _,count:= range sum {
        if count != 0 {
            return false
        }

        
    }

    return true
}