package checkRam

import (
	"fmt"
	"github.com/PumpkinSeed/cage"
	//"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/rwtodd/Go.Sed/sed"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func checkRam() {
	var err error
	free := exec.Command("free")
	grep := exec.Command("grep", "Mem")

	// Set grep's stdin to the output of the free command.
	grep.Stdin, err = free.StdoutPipe()
	if err != nil {
		log.Fatalln(err)
	}

	//start capturing stdout data
	capture := cage.Start()
	// Set grep's stdout to os.Stdout
	grep.Stdout = os.Stdout

	// Start the grep command first. (The order will be last command first)
	must(grep.Start())
	// Run the free command. (Run calls start and also calls wait)
	must(free.Run())
	// Wait for the grep command to finish.
	must(grep.Wait())
	cage.Stop(capture)

	ram := capture.Data[0]

	//removing extra characters from the string
	removeChars, err := sed.New(strings.NewReader(`s|[Mem:,]||g`))
	ram, err = removeChars.RunString(ram)

	//removing extra characters from the string - 16
	removeNumbers, err := sed.New(strings.NewReader(`s/.{55}$//`))
	ram, err = removeNumbers.RunString(ram)

	//removing extra characters from the string
	removeWhitespace, err := sed.New(strings.NewReader(`s/\s+//g`))
	ram, err = removeWhitespace.RunString(ram)

	ram = strings.TrimSuffix(ram, "\n")

	ramInt, err := strconv.Atoi(ram)
	if err == nil {
		fmt.Println(ramInt)
	}
	fmt.Printf("\nValue: %v", ramInt)

}

func must(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
