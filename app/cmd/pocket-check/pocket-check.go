package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/peterbourgon/ff/v3/ffcli"
	"os"
	"src/checkCpu"
	"src/checkHd"
	"src/checkPocket"
	"src/checkPorts"
	"src/checkRam"
	"src/checkRunning"
	"src/checkUlimit"
)

func main() {
	var (
		rootFlagSet = flag.NewFlagSet("pokt-check", flag.ExitOnError)
	)

	f := func(context.Context, []string) error {
		checkRam.CheckRam()
		checkHd.CheckHd()
		checkCpu.CheckCpu()
		checkPorts.CheckPorts()
		//checkNginx.CheckNginx()
		checkUlimit.CheckUlimit()
		checkRunning.CheckRunning()
		checkPocket.CheckPocket()
		return nil
	}

	root := &ffcli.Command{
		ShortUsage:  "pokt-check",
		FlagSet:     rootFlagSet,
		Subcommands: []*ffcli.Command{},
		Exec:        f,
	}

	if err := root.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		fprintf, err := fmt.Fprintf(os.Stderr, "error: %v\n", err)
		if err != nil {
			fmt.Println(fprintf)
		}
		os.Exit(1)
	}
}
