package executer

import (
	"os/exec"

	sdapi "github.com/screwdriver-cd/sd-cmd/screwdriver/api"
)

// Binary is Binary Executer object
type Binary struct {
	Cmd *sdapi.Command
	Arg []string
}

// NewBinary return Binary object
func NewBinary(cmd *sdapi.Command, arg []string) (*Binary, error) {
	binary := &Binary{
		Cmd: cmd,
		Arg: arg,
	}
	return binary, nil
}

// Run exec command and return output
func (b *Binary) Run() ([]byte, error) {
	return exec.Command("ls", b.Arg...).Output()
}
