package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/sebastianwebber/wait4/lib"
)

const currentVersion = "0.2.0"

var (
	portNumber  *int
	interval    *time.Duration
	command     *string
	showVersion *bool
)

func processArgs() {
	flag.Usage = func() {
		fmt.Printf("Usage: wait4 --port=[port number] [options] \n\n")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	portNumber = flag.Int("port", 0, "port name to check")
	command = flag.String("command", "", "port name to check")
	interval = flag.Duration("interval", 1*time.Second, "time for each check")
	showVersion = flag.Bool("version", false, "Print version information and quit")
	flag.Parse()

	if *showVersion {
		fmt.Printf("wait4 v%s\n", currentVersion)
		os.Exit(0)
	}
}

func main() {

	processArgs()

	c := time.Tick(*interval)
	for range c {
		check, _ := lib.PortCheck(*portNumber, 5*time.Second)
		fmt.Printf("Waiting for %d port...\n", *portNumber)
		if check {
			fmt.Printf("%d port is listenning!\n", *portNumber)
			break
		}
	}

}
