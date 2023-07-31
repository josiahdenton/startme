package nvim

import (
	"os"
	"os/exec"
)

type Nvim struct {
	cmd *exec.Cmd
}

func New() *Nvim {
	cmd := exec.Command("nvim")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return &Nvim{
		cmd: cmd,
	}
}

func (n *Nvim) Open(path string) error {
	n.cmd.Args = append(n.cmd.Args, path)
	err := n.cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (n *Nvim) Wait() error {
	return n.cmd.Wait()
}
