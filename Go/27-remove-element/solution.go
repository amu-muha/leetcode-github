func removeElement(nums []int, val int) int {
    k:= 0
    for _,num := range nums {
        if num != val {
           nums[k] = num
           k++
        }
    }
    return k
    
}