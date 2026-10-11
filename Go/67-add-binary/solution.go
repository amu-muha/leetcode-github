func addBinary(a string, b string) string {
      i := len(a) - 1
    j := len(b) - 1

    carry := 0
    result := ""

    for i >= 0 || j >= 0 || carry > 0 {
        sum := carry

        if i >= 0 {
            sum += int(a[i] - '0')
            i--
        }

        if j >= 0 {
            sum += int(b[j] - '0')
            j--
        }

        result = string(rune('0'+sum%2)) + result
        carry = sum / 2
    }

    return result
}