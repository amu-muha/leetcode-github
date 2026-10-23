func convert(s string, numRows int) string {
    if numRows == 1 || numRows >= len(s) {
        return s
    }

    rows := make([]string, numRows)
    currentRow := 0
    goingDown := false

    for i := 0; i < len(s); i++ {
        rows[currentRow] += string(s[i])

        if currentRow == 0 || currentRow == numRows-1 {
            goingDown = !goingDown
        }

        if goingDown {
            currentRow++
        } else {
            currentRow--
        }
    }

    result := ""
    for _, row := range rows {
        result += row
    }

    return result
}
