package lib

import (
	"fmt"
	"net"
	"time"
)

// PortCheck check if a local port is open
// timeout are provide in duration.
func PortCheck(port int, timeout time.Duration) (listen bool, err error) {

	conn, err := net.DialTimeout("tcp", fmt.Sprintf(":%d", port), timeout)
	if conn != nil {
		defer conn.Close()
	}

	return err == nil, err
}
