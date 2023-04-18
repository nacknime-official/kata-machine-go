package gotest

import "path/filepath"

func PrepareArgs(args []string, dayPath string) []string {
	// if the first argument is a flag or there's no arguments at all, then
	// we assume that the user wants to test the all tests for the current day
	// otherwise, we assume that the user wants to test the specific test
	// and we must add the "./" prefix to those tests
	// but the user can add additional flags that will be passed to the go test command
	if len(args) == 0 || args[0][0] == '-' {
		args = append([]string{"./" + filepath.Join(dayPath, "...")}, args...)
	} else {
		for i, arg := range args {
			if arg[0] == '-' {
				break
			}
			args[i] = "./" + filepath.Join(dayPath, arg)
		}
	}

	return args
}
