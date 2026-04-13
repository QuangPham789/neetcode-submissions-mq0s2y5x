func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    original_color := image[sr][sc]
    if original_color == color{
        return image
    }
    ROWS , COLS := len(image), len(image[0])
    


    var dfs  func(r,c int)

    dfs = func(r, c int){
        if r < 0 || c < 0 || r >= ROWS || c >= COLS || image[r][c] != original_color {
            return
        }

        
        image[r][c] = color
        dfs(r+1, c)
        dfs(r-1, c)
        dfs(r, c+1)
        dfs(r, c-1)

    } 

    dfs(sr,sc)
    return image
}
