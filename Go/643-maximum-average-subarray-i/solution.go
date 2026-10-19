func findMaxAverage(nums []int, k int) float64 {
    maxSum:= 0
    maxAve:=0

    for i:=0;i<k;i++{
        maxSum += nums[i]
    }
    maxAve = maxSum
    for left,right:= 0,k;right<len(nums);left, right = left+1, right+1{

        maxSum = maxSum - nums[left] + nums[right]
        if maxAve < maxSum{
            maxAve = maxSum
        }
        
    }
    result := float64(maxAve) / float64(k)
    return result

}