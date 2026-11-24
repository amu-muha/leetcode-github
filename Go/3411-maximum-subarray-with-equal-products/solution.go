func maxLength(nums []int) int {
    n := len(nums)
    maxLen := 0

    gcd := func(a, b int) int {
        for b != 0 {
            a, b = b, a%b
        }
        return a
    }

    lcm := func(a, b int) int {
        return (a / gcd(a, b)) * b
    }

    for i := 0; i < n; i++ {
        product := 1
        g := nums[i]
        l := nums[i]

        for j := i; j < n; j++ {
            product *= nums[j]
            g = gcd(g, nums[j])
            l = lcm(l, nums[j])

            if product == g*l {
                length := j - i + 1
                if length > maxLen {
                    maxLen = length
                }
            }
        }
    }

    return maxLen
}