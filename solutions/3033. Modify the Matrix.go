func modifiedMatrix(matrix [][]int) [][]int {
    cols := make([]int, len(matrix[0]))
    for _, r := range matrix {
        for j, c := range r {
            if c > cols[j] {
                cols[j] = c 
            }
        }
    }
    for i, r := range matrix {
        for j, c := range r{
            if c == -1 {
                matrix[i][j] = cols[j]
            }
        }
    }
    return matrix
}