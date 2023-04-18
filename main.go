package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/nacknime-official/kata-machine-go/internal/config"
	"github.com/nacknime-official/kata-machine-go/internal/day"
	"github.com/nacknime-official/kata-machine-go/internal/generate"
	"github.com/nacknime-official/kata-machine-go/internal/gotest"
)

func main() {
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	generateConfigPath := generateCmd.String("configPath", config.DefaultConfigPath, "Config path")

	testCmd := flag.NewFlagSet("test", flag.ExitOnError)
	testSpecificDay := testCmd.Int("day", 0, "Specific day to test")

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "generate subcommand is required")
		os.Exit(1)
	}

	cmd := os.Args[1]
	switch cmd {
	case "generate":
		generateCmd.Parse(os.Args[2:])
		generateSpecificDSAs := generateCmd.Args()

		cfg, err := config.Load(*generateConfigPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error on config loading: %s\n", err.Error())
			os.Exit(1)
		}

		dsa := cfg.DSA
		if len(generateSpecificDSAs) != 0 {
			dsa = generateSpecificDSAs
		}

		if err := generate.Generate(dsa); err != nil {
			fmt.Fprintf(os.Stderr, "Error on generate: %s\n", err.Error())
			os.Exit(1)
		}

	case "test":
		testCmd.Parse(os.Args[2:])
		args := testCmd.Args()

		specificDay := *testSpecificDay
		if specificDay == 0 {
			var err error
			specificDay, err = day.GetCurrent()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error on getting current day: %s\n", err.Error())
				os.Exit(1)
			}
		}
		specificDayPath := day.GetDayDirPath(specificDay)

		args = gotest.PrepareArgs(args, specificDayPath)
		args = append([]string{"test"}, args...)

		cmd := exec.Command("go", args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error on test: %s\n", err.Error())
			os.Exit(1)
		}

	default:
		// TODO: print usage
		fmt.Fprintf(os.Stderr, "Unknown command %s.\nAvailable commands: generate\n", cmd)
		os.Exit(1)
	}
}
