package lib

import (
	"os"
	"os/exec"
	"strconv"

	"github.com/hooph00p/secretary/lib/util"
)

type Task struct {
	i   int
	out *os.File
}

func (t *Task) Out() *os.File {
	if o == "" {
		t.out = os.Stdout
	} else {
		os.Create(o)
		t.out, _ = os.Open(o)
	}
}

func (t *Task) Run(args []string) (string, error) {

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = t.out
	cmd.Stderr = t.out
	defer t.out.Close()

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	util.PL(t.i, "Finishing")

	return "Finished task " + strconv.Itoa(t.i), nil
}
