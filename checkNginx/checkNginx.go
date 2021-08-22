package checkNginx

import (
	"bufio"
	"fmt"
	"github.com/PumpkinSeed/cage"
	"github.com/fatih/color"
	"log"
	"os"
	"strings"
)

func CheckNginx() {

	proxyConfig := readConfig()
	configCompare(proxyConfig)

}

func readConfig() string {
	file, err := os.Open("/etc/nginx/sites-enabled/pocket-proxy.conf")
	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	capture := cage.Start()

	for _, each_ln := range text {
		fmt.Println(each_ln)
	}

	cage.Stop(capture)

	output := capture.Data
	configString := strings.Join(output, "\n")
	return configString
}

func configCompare(proxyConfig string) {
	if listenSSL(proxyConfig) == true && certificateSSL(proxyConfig) == true && keySSL(proxyConfig) == true {
		color.Green("SUCCESS: Reverse proxy is configured correctly")
	} else {
		color.Red("Error: Reverse proxy is not configured correctly")
	}
}

func listenSSL(proxyConfig string) bool {
	return true
}

func certificateSSL(proxyConfig string) bool {
	return true
}

func keySSL(proxyConfig string) bool {
	return true
}
