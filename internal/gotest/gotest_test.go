package gotest

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testPkgPath(testPath, testName string) string {
	return "./" + filepath.Join(testPath, testName)
}

func TestPrepareArgs(t *testing.T) {
	dayPath := filepath.Join("src", "day1")

	tests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "no args",
			args: []string{},
			want: []string{testPkgPath(dayPath, "...")},
		},
		{
			name: "1 specific test",
			args: []string{"Queue"},
			want: []string{testPkgPath(dayPath, "Queue")},
		},
		{
			name: "2 specific tests",
			args: []string{"Queue", "Stack"},
			want: []string{testPkgPath(dayPath, "Queue"), testPkgPath(dayPath, "Stack")},
		},
		{
			name: "1 flag",
			args: []string{"-v"},
			want: []string{testPkgPath(dayPath, "..."), "-v"},
		},
		{
			name: "2 flags",
			args: []string{"-v", "-coverprofile=coverage.out"},
			want: []string{testPkgPath(dayPath, "..."), "-v", "-coverprofile=coverage.out"},
		},
		{
			name: "tests and flags",
			args: []string{"Queue", "Stack", "-v"},
			want: []string{testPkgPath(dayPath, "Queue"), testPkgPath(dayPath, "Stack"), "-v"},
		},
		{
			name: `flags and tests. should not add "./" to tests`,
			args: []string{"-v", "Queue", "Stack"},
			want: []string{testPkgPath(dayPath, "..."), "-v", "Queue", "Stack"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, PrepareArgs(tt.args, dayPath))
		})
	}
}
