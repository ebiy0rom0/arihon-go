package lake_counting

var width int
var height int
var field [][]string

func Calc(n, m int, f [][]string) int {
	field = f
	width = m
	height = n
	cnt := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if field[y][x] == "w" {
				dfs(x, y)
				cnt++
			}
		}
	}
	return cnt
}

func dfs(x, y int) {
	field[y][x] = "."

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			nx := x + i
			ny := y + j

			if nx >= 0 && ny >= 0 && nx < width && ny < height && field[ny][nx] == "w" {
				dfs(nx, ny)
			}
		}
	}
}
