package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/go-playground/log"
	"github.com/go-playground/log/handlers/console"

	"github.com/umovme/wait4/lib"
)

const currentVersion = "0.5.0"

var (
	portNumber  *int
	interval    *time.Duration
	timeout     *time.Duration
	command     *string
	showVersion *bool
	debugMode   *bool
	logLevels   = []log.Level{
		log.InfoLevel,
		log.NoticeLevel,
		log.WarnLevel,
		log.ErrorLevel,
		log.PanicLevel,
		log.AlertLevel,
		log.FatalLevel,
	}
)

func processArgs() {

	flag.Usage = func() {
		fmt.Printf("Usage: wait4 [options] \n\n")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	portNumber = flag.Int("port", 0, "port name to check")
	command = flag.String("command", "", "port name to check")
	interval = flag.Duration("interval", 1*time.Second, "time for each check")
	timeout = flag.Duration("timeout", 5*time.Second, "timeout for a port or command check")
	showVersion = flag.Bool("version", false, "Print version information and quit")
	debugMode = flag.Bool("debug", false, "show debug messages")
	flag.Parse()

	if *showVersion {
		fmt.Printf("wait4 v%s\n", currentVersion)
		os.Exit(0)
	}

	if *debugMode {
		logLevels = append(logLevels, log.DebugLevel)
	}

	log.Debugf("%#v", *debugMode)

	cLog := console.New(true)
	log.AddHandler(cLog, logLevels...)
}

func main() {
	processArgs()
	defer log.WithTrace().Info("Done.")

	var (
		checkPort    bool
		checkCommand bool
		err          error
	)
	c := time.Tick(*interval)

	if *portNumber > 0 {
		for range c {
			checkPort, _ = lib.PortCheck(*portNumber, *timeout)
			log.Infof("Waiting for %d port...", *portNumber)
			if checkPort {
				log.Infof("%d port is listenning!", *portNumber)
				break
			}
		}
	}

	if *command != "" {
		var out []byte
		for range c {
			checkCommand, out, err = lib.CmdCheck(*command)
			if err != nil {
				log.Infof("%s returns:\n%s%s", *command, string(out), err.Error())
			}
			if checkCommand {
				log.Infof("Command '%s' returns without error.", *command)
				break
			}
		}
	}
}
