package lib

import (
	"github.com/go-playground/log"
	"github.com/sebastianwebber/cmdr"
)

// CmdCheck runs a command to check if the service is running
// the main idea its run command into a running server and
// validate if it is running
func CmdCheck(cmd string) (bool, []byte, error) {

	command := cmdr.Parse(cmd)
	log.Debugf("%#v", command)

	output, err := command.Run()
	log.Debugf("%#v", string(output))

	return err == nil, output, err
}
