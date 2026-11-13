func longestMonotonicSubarray(nums []int) int {
    if len(nums) == 0{
        return 0
    }

    inc,dec,maxlen:= 1,1,1

    for i:= 1;i<len(nums);i++{
        if nums[i]> nums[i-1]{
            inc++
            dec=1
        }else if nums[i]<nums[i-1]{
            dec++
            inc=1
        }else{
            inc = 1
            dec = 1
        }
        maxlen=max(maxlen,max(inc,dec))
    }
    return maxlen
}