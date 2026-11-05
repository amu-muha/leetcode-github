func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
    count:= 0
    
    for i:= 0; i<len(arr1); i++{
        isValid:= true
        for j:= 0; j<len(arr2);j++{
            dif:= arr1[i] - arr2[j]
            if dif < 0 {
                dif = -dif
            }
            if dif <= d{
                isValid = false
                break
            }
        
            
        }
         if isValid {
            count++
        }
    }
   
       
    

    return count
}