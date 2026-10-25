func solveSudoku(board [][]byte) {
    backtrack(board)
}

func backtrack(board [][]byte) bool {
    for row := 0; row < 9; row++ {
        for col := 0; col < 9; col++ {

            // Skip cells that are already filled.
            if board[row][col] != '.' {
                continue
            }

            // Try digits 1 through 9.
            for digit := byte('1'); digit <= '9'; digit++ {
                if isValid(board, row, col, digit) {
                    board[row][col] = digit

                    // If the rest of the board can be solved,
                    // keep this choice.
                    if backtrack(board) {
                        return true
                    }

                    // Otherwise, undo the choice.
                    board[row][col] = '.'
                }
            }

            // No digit works for this cell.
            return false
        }
    }

    // Every cell is filled successfully.
    return true
}

func isValid(board [][]byte, row, col int, digit byte) bool {
    for i := 0; i < 9; i++ {
        // Check the row.
        if board[row][i] == digit {
            return false
        }

        // Check the column.
        if board[i][col] == digit {
            return false
        }

        // Check the 3x3 box.
        boxRow := (row/3)*3 + i/3
        boxCol := (col/3)*3 + i%3

        if board[boxRow][boxCol] == digit {
            return false
        }
    }

    return true
}