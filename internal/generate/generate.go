package generate

import (
	"fmt"
	"github.com/nacknime-official/kata-machine-go/internal/day"
	"github.com/nacknime-official/kata-machine-go/internal/dsa"
)

func Generate(DSAs []string) error {
	nextDay, err := day.GetNext()
	if err != nil {
		return fmt.Errorf("get next day: %w", err)
	}
	nextDayDirPath := day.GetDayDirPath(nextDay)

	for _, d := range DSAs {
		if err := dsa.Copy(d, nextDayDirPath); err != nil {
			return err
		}
	}

	return nil
}
