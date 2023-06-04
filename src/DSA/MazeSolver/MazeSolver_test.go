package mazesolver

import (
	"strings"
	"testing"
)

func TestMazeSolver(t *testing.T) {
	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	mazeResult := []Point{
		{x: 10, y: 0},
		{x: 10, y: 1},
		{x: 10, y: 2},
		{x: 10, y: 3},
		{x: 10, y: 4},
		{x: 9, y: 4},
		{x: 8, y: 4},
		{x: 7, y: 4},
		{x: 6, y: 4},
		{x: 5, y: 4},
		{x: 4, y: 4},
		{x: 3, y: 4},
		{x: 2, y: 4},
		{x: 1, y: 4},
		{x: 1, y: 5},
	}

	result := Solve(maze, "x", Point{x: 10, y: 0}, Point{x: 1, y: 5})

	expected := drawPath(maze, mazeResult)
	got := drawPath(maze, result)

	if !stringSlicesEqual(got, expected) {
		t.Errorf("Expected result: %v, got: %v", expected, got)
	}
}

func drawPath(data []string, path []Point) []string {
	data2 := make([][]string, len(data))
	for i, row := range data {
		data2[i] = strings.Split(row, "")
	}
	for _, p := range path {
		if p.y < len(data2) && p.x < len(data2[p.y]) {
			data2[p.y][p.x] = "*"
		}
	}

	result := make([]string, len(data2))
	for i, d := range data2 {
		result[i] = strings.Join(d, "")
	}
	return result
}

func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
