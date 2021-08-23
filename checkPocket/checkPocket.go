package checkPocket

import (
	"log"
	"os"
	"os/exec"
)

func CheckPocket() {
	pocket := exec.Command("pocket", "util", "print-configs")

	pocket.Stdout = os.Stdout
	must(pocket.Run())

}

func must(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
