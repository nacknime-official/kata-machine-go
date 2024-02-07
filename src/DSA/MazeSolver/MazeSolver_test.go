package mazesolver

import (
	"testing"

	"github.com/stretchr/testify/assert"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
)

func TestSolve(t *testing.T) {
	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}
	expected := []dsa.Point{
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
	assert.Equal(t, expected, result, "Expected: %v, got: %v", expected, result)
	expectedPath := drawPath(maze, expected)
	actualPath := drawPath(maze, result)
	assert.Equal(t, expectedPath, actualPath, "Expected: %v, got: %v", expectedPath, actualPath)
}

func drawPath(data []string, path []dsa.Point) []string {
	var data2 [][]rune
	for _, row := range data {
		data2 = append(data2, []rune(row))
	}
	for _, p := range path {
		if p.Y < len(data2) && p.X < len(data2[p.Y]) {
			data2[p.Y][p.X] = '*'
		}
	}
	var result []string
	for _, d := range data2 {
		result = append(result, string(d))
	}
	return result
}
