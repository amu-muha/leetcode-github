func myAtoi(s string) int {
    i := 0
    n := len(s)

    // 1. Skip leading whitespace
    for i < n && s[i] == ' ' {
        i++
    }

    // 2. Determine the sign
    sign := 1
    if i < n && s[i] == '-' {
        sign = -1
        i++
    } else if i < n && s[i] == '+' {
        i++
    }

    // 3. Convert digits and check overflow
    num := 0
    maxInt := 1<<31 - 1
    minInt := -(1 << 31)

    for i < n && s[i] >= '0' && s[i] <= '9' {
        digit := int(s[i] - '0')

        if num > maxInt/10 ||
            (num == maxInt/10 && digit > maxInt%10) {
            if sign == 1 {
                return maxInt
            }
            return minInt
        }

        num = num*10 + digit
        i++
    }

    return num * sign
}
