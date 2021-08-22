package checkPorts

import (
	"context"
	"fmt"
	"github.com/PumpkinSeed/cage"
	"github.com/Ullaakut/nmap/v2"
	"github.com/fatih/color"
	"log"
	"time"
)

func CheckPorts() {

	result := portScan()
	portCompare(result)

}

func portScan() []string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	// Equivalent to `/usr/local/bin/nmap -p 8081, 26656 IP localhost`,
	// with a 5 minute timeout.
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(
			"localhost",
		),
		nmap.WithPorts("26656,8081"),
		nmap.WithContext(ctx),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	result, warnings, err := scanner.Run()
	if err != nil {
		log.Fatalf("unable to run nmap scan: %v", err)
	}

	if warnings != nil {
		log.Printf("Warnings: \n %v", warnings)
	}

	capture := cage.Start()

	fmt.Printf("%d %s\n", result.Hosts[0].Ports[0].ID, result.Hosts[0].Ports[0].State)
	fmt.Printf("%d %s\n", result.Hosts[0].Ports[1].ID, result.Hosts[0].Ports[1].State)

	cage.Stop(capture)

	output := capture.Data
	return output
}

func portCompare(nmapResponse []string) {
	if nmapResponse[0] == "8081 open" && nmapResponse[1] == "26656 open" {
		color.Green("SUCCESS: Local ports are open")
	} else {
		color.Red("ERROR: Local ports not open")
	}
}
