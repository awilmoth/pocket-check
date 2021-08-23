package checkUlimit

import (
	"github.com/PumpkinSeed/cage"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"strconv"
)

func CheckUlimit() {
	ul := getUlimit()
	ulimitCompare(ul)
}

func getUlimit() int {
	ulimit := exec.Command("/bin/bash", "-c", "ulimit -n")

	capture := cage.Start()
	ulimit.Stdout = os.Stdout

	err := ulimit.Run()
	if err != nil {
		return 0
	}

	cage.Stop(capture)

	strUlimit := capture.Data[0]
	intUlimit, _ := strconv.Atoi(strUlimit)

	return intUlimit
}

func ulimitCompare(ulimit int) {
	if ulimit >= 16384 {
		color.Green("SUCCESS: ulimit is at least 16384")
	} else {
		color.Red("ERROR: ulimit is under 16384")
	}
}
