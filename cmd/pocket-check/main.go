package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/awilmoth/pocket-check/pkg/checkRam"
	"github.com/awilmoth/pocket-check/pkg/checkCpu"
	"github.com/awilmoth/pocket-check/pkg/checkHd"
	"github.com/awilmoth/pocket-check/pkg/checkPorts"
	"github.com/awilmoth/pocket-check/pkg/checkNginx"
	"github.com/awilmoth/pocket-check/pkg/checkRunning"
	"os"
)

func main() {
	var (
		rootFlagSet = flag.NewFlagSet("pokt-check", flag.ExitOnError)
	)

	f := func(context.Context, []string) error {
		checkRam()
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