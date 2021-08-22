package checkUlimit

import (
	"github.com/fatih/color"
)

func CheckUlimit() {

	data := getUlimit()
	ulimitCompare(data)
}

func getUlimit() int {

	return 1
}

func ulimitCompare(ulimit int) {
	if ulimit >= 16384 {
		color.Green("SUCCESS: ulimit is at least 16384")
	} else {
		color.Red("ERROR: ulimit is under 16384")
	}
}
