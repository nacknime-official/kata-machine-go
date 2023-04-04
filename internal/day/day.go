package day

import (
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
)

var DaysPath = filepath.FromSlash("src")

func GetCurrent() (int, error) {
	files, err := filepath.Glob(filepath.Join(DaysPath, "day") + "*")
	if err != nil {
		return 0, fmt.Errorf("glob: %w", err)
	}

	var days []int
	for _, file := range files {
		rawDay := filepath.Base(file)[3:]
		day, err := strconv.Atoi(rawDay)
		if err != nil {
			continue
		}
		days = append(days, day)
	}
	sort.Ints(days)

	if len(days) == 0 {
		return 0, nil
	}
	return days[len(days)-1], nil
}

func GetNext() (int, error) {
	currDay, err := GetCurrent()
	if err != nil {
		return currDay, err
	}
	return currDay + 1, nil
}

func GetDayDirPath(day int) string {
	return filepath.Join(DaysPath, "day"+strconv.Itoa(day))
}
