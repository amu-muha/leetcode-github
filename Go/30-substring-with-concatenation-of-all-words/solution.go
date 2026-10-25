func findSubstring(s string, words []string) []int {
    result := []int{}

    if len(s) == 0 || len(words) == 0 {
        return result
    }

    wordLen := len(words[0])
    wordCount := len(words)
    totalLen := wordLen * wordCount

    if len(s) < totalLen {
        return result
    }

    // Count how many times each word appears in words.
    wordFreq := make(map[string]int)
    for _, word := range words {
        wordFreq[word]++
    }

    // Try each possible starting offset within a word.
    for offset := 0; offset < wordLen; offset++ {
        left := offset
        seen := make(map[string]int)
        count := 0

        for right := offset; right+wordLen <= len(s); right += wordLen {
            word := s[right : right+wordLen]

            // If this word isn't required, reset the window.
            if _, exists := wordFreq[word]; !exists {
                seen = make(map[string]int)
                count = 0
                left = right + wordLen
                continue
            }

            seen[word]++
            count++

            // If this word occurs too many times, shrink the window.
            for seen[word] > wordFreq[word] {
                leftWord := s[left : left+wordLen]
                seen[leftWord]--
                count--
                left += wordLen
            }

            // All words are present in the correct quantities.
            if count == wordCount {
                result = append(result, left)

                // Move left forward to continue searching.
                leftWord := s[left : left+wordLen]
                seen[leftWord]--
                count--
                left += wordLen
            }
        }
    }

    return result
}