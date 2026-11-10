func climbStairs(n int, costs []int) int {
    dp := make([]int, n+1)

    for i := 1; i <= n; i++ {
        dp[i] = int(^uint(0) >> 1)

        for j := 1; j <= 3; j++ {
            if i-j >= 0 {
                cost := dp[i-j] + costs[i-1] + j*j

                if cost < dp[i] {
                    dp[i] = cost
                }
            }
        }
    }

    return dp[n]
}