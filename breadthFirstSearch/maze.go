package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	_, err = fmt.Fscanf(file, "%d %d\n", &row, &col)
	if err != nil {
		panic(err)
	}
	fmt.Println(row, col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			//_, err = fmt.Fscanf(file, "\n")
			_, _ = fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

type point struct {
	i, j int
}

var directions = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func (p point) add(dir point) point {
	return point{p.i + dir.i, p.j + dir.j}
}


func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= (len(grid)) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end point)[][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	queue := []point{start}
	for len(queue)>0{
		cur := queue[0]
		if cur == end{
			break
		}
		queue = queue[1:]
		for _, dir := range directions {
			next := cur.add(dir)
			value, ok := next.at(maze)
			//判断是否为墙或者越界
			if !ok || value == 1 {
				continue
			}
			value, ok = next.at(steps)
			//判断是否已经走过
			if value != 0 {
				continue
			}
			//判断是否走到终点
			if next == start {
				continue
			}
			curStepValue, _ := cur.at(steps)
			steps[next.i][next.j] = curStepValue + 1
			queue = append(queue,next)
		}
	}
	return steps
}

func main() {
	maze := readMaze("maze.txt")
	steps := walk(maze,point{0,0},point{len(maze)-1,len(maze[0])-1})
	for _, row :=range steps{
			for _, value := range row {
				fmt.Printf("%d ", value)
			}
			fmt.Println()
	}
}
