package checkIP

import (
	"fmt"
	"github.com/fatih/color"
	"io"
	"net/http"
	"os"
)

func CheckIP() {
	fmt.Println()
	color.Blue("Your external IP is the following:")

	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	fmt.Println()
	fmt.Println()
}
