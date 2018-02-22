package lib

import (
	"strings"

	"github.com/sebastianwebber/cmdr"
)

// CmdCheck runs a command to check if the service is running
// the main idea its run command into a running server and
// validate if it is running
func CmdCheck(cmd string) (bool, []byte, error) {
	cmdParts := strings.Split(cmd, " ")
	output, err := cmdr.New(false, cmdParts[0], cmdParts[1:len(cmdParts)]...).Run()

	return err == nil, output, err
}
