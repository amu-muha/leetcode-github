func reverse(x int) int {
    result := 0

    for x != 0 {
        digit := x % 10
        x /= 10

        if result > 214748364 || (result == 214748364 && digit > 7) {
            return 0
        }

        if result < -214748364 || (result == -214748364 && digit < -8) {
            return 0
        }

        result = result*10 + digit
    }

    return result
}