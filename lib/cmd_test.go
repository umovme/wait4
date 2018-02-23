package lib

import (
	"testing"
)

func TestCmdCheck(t *testing.T) {
	tests := []struct {
		name    string
		cmd     string
		want    bool
		wantErr bool
	}{
		{name: "ok echo", cmd: "echo 'hello world'", want: true, wantErr: false},
		{name: "exit with the error code", cmd: "exit 2", want: false, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := CmdCheck(tt.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("CmdCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CmdCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
