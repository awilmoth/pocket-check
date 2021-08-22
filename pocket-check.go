package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/peterbourgon/ff/v3/ffcli"
	"log"
	"os"
	"os/exec"
)

func main() {
	var (
		rootFlagSet = flag.NewFlagSet("pokt-check", flag.ExitOnError)
	)

	root := &ffcli.Command{
		ShortUsage:  "pokt-check",
		FlagSet:     rootFlagSet,
		Subcommands: []*ffcli.Command{},
		Exec: func(context.Context, []string) error {
			cmd := exec.Command("ls", "-lah")
			out, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatalf("Failed with error %s\n", err)
			}
			fmt.Printf(string(out))
			return nil
		},
	}
	if err := root.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
