package lib

import (
	"fmt"
	"log"
	"net"
	"testing"
	"time"
)

const (
	defaultTimeout = 5 * time.Second
	defaultNetwork = "tcp"
	newPortTest    = 48541
)

func openPort(port, count int) (err error) {

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer l.Close()
		l.Accept()
	}()

	return
}

func TestCheckLocalPort(t *testing.T) {

	tests := []struct {
		name       string
		wantListen bool
		wantErr    bool
		openPort   bool
		port       int
	}{
		// {name: "check postgres on my machine", wantListen: true, wantErr: false, port: 5432},
		{name: "check new port", wantListen: true, wantErr: false, openPort: true, port: newPortTest},
		{name: "check wrong port", wantListen: false, wantErr: true, port: 83463},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.openPort {
				openPort(tt.port, 1)
			}

			gotListen, err := PortCheck(tt.port, defaultTimeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckLocalPort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotListen != tt.wantListen {
				t.Errorf("CheckLocalPort() = %v, want %v", gotListen, tt.wantListen)
			}
		})
	}
}
