package main

import (
	"flag"
	"fmt"
	"github.com/nacknime-official/kata-machine-go/internal/config"
	"github.com/nacknime-official/kata-machine-go/internal/generate"
	"os"
)

func main() {
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	generateConfigPath := generateCmd.String("configPath", config.DefaultConfigPath, "Config path")
	// TODO: test command

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
	default:
		// TODO: print usage
		fmt.Fprintf(os.Stderr, "Unknown command %s.\nAvailable commands: generate\n", cmd)
		os.Exit(1)
	}
}
