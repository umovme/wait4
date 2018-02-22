package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sebastianwebber/wait4/lib"
)

const currentVersion = "0.3.0"

var (
	portNumber  *int
	interval    *time.Duration
	timeout     *time.Duration
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
	timeout = flag.Duration("timeout", 5*time.Second, "timeout for a port or command check")
	showVersion = flag.Bool("version", false, "Print version information and quit")
	flag.Parse()

	if *showVersion {
		fmt.Printf("wait4 v%s\n", currentVersion)
		os.Exit(0)
	}

}

func main() {

	processArgs()

	var (
		checkPort    bool
		checkCommand bool
		err          error
	)
	c := time.Tick(*interval)

	if *portNumber > 0 {
		for range c {
			checkPort, _ = lib.PortCheck(*portNumber, *timeout)
			log.Printf("Waiting for %d port...\n", *portNumber)
			if checkPort {
				log.Printf("%d port is listenning!\n", *portNumber)
				break
			}
		}
	}

	if *command != "" {
		var out []byte
		for range c {
			checkCommand, out, err = lib.CmdCheck(*command)
			if err != nil {
				log.Printf("%s returns:\n%s%s\n", *command, string(out), err.Error())
			}
			if checkCommand {
				log.Printf("Command '%s' returns without error.\n", *command)
				break
			}
		}
	}

	log.Println("Done.")
}
