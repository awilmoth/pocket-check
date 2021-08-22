package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/peterbourgon/ff/v3/ffcli"
	"os"
	"src/checkRam"
)

func main() {
	var (
		rootFlagSet = flag.NewFlagSet("pokt-check", flag.ExitOnError)
	)

	f := func(context.Context, []string) error {
		checkRam.CheckRam()
		return nil
	}

	root := &ffcli.Command{
		ShortUsage:  "pokt-check",
		FlagSet:     rootFlagSet,
		Subcommands: []*ffcli.Command{},
		Exec:        f,
	}

	if err := root.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
