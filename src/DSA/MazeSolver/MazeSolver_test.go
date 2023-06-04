package mazesolver

import (
	"strings"
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
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

	mazeResult := []dsa.Point{
		{X: 10, Y: 0},
		{X: 10, Y: 1},
		{X: 10, Y: 2},
		{X: 10, Y: 3},
		{X: 10, Y: 4},
		{X: 9, Y: 4},
		{X: 8, Y: 4},
		{X: 7, Y: 4},
		{X: 6, Y: 4},
		{X: 5, Y: 4},
		{X: 4, Y: 4},
		{X: 3, Y: 4},
		{X: 2, Y: 4},
		{X: 1, Y: 4},
		{X: 1, Y: 5},
	}

	result := Solve(maze, "x", dsa.Point{X: 10, Y: 0}, dsa.Point{X: 1, Y: 5})

	expected := drawPath(maze, mazeResult)
	got := drawPath(maze, result)

	if !stringSlicesEqual(got, expected) {
		t.Errorf("Expected result: %v, got: %v", expected, got)
	}
}

func drawPath(data []string, path []dsa.Point) []string {
	data2 := make([][]string, len(data))
	for i, row := range data {
		data2[i] = strings.Split(row, "")
	}
	for _, p := range path {
		if p.Y < len(data2) && p.X < len(data2[p.Y]) {
			data2[p.Y][p.X] = "*"
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
