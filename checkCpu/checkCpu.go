package checkCpu

import (
	"github.com/fatih/color"
	"runtime"
)

func CheckCpu() {

	cores := runtime.NumCPU()
	cpuCompare(cores)

}

func cpuCompare(cores int) {
	if cores >= 4 {
		color.Green("SUCCESS: 4 or more CPU cores")
	} else {
		color.Red("ERROR: 3 or less CPU cores")
	}
}
