func matrixReshape(mat [][]int, r int, c int) [][]int {
    n, m := len(mat), len(mat[0])
    if n * m != r * c {
        return mat
    }
    
    // build 2D array
    arr := make([][]int, r)
    for i := range arr {
        arr[i] = make([]int, c)
    }
    
    var x, y int = 0, 0
    
    for i := 0; i < r; i++ {
        for j := 0; j < c; j++ {
            arr[i][j] = mat[x][y]
            y += 1
            if y == m {
                y = 0
                x += 1
            }
        }
    }
    return arr
}