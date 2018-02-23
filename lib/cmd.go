package lib

import (
	"github.com/sebastianwebber/cmdr"
)

// CmdCheck runs a command to check if the service is running
// the main idea its run command into a running server and
// validate if it is running
func CmdCheck(cmd string) (bool, []byte, error) {
	output, err := cmdr.Parse(cmd).Run()

	return err == nil, output, err
}
