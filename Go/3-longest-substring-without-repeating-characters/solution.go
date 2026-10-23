
func lengthOfLongestSubstring(s string) int {
    seen := make(map[byte]int)
    left := 0
    maxLen := 0

    for right := 0; right < len(s); right++ {
        if index, exists := seen[s[right]]; exists && index >= left {
            left = index + 1
        }

        seen[s[right]] = right

        length := right - left + 1
        if length > maxLen {
            maxLen = length
        }
    }

    return maxLen
}
