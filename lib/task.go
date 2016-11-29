package lib

import (
	"os"
	"os/exec"
	"strconv"

	"github.com/hooph00p/secretary/lib/util"
)

type Task struct {
	TaskId int
}

func (t *Task) Run(args []string, o string) (string, error) {

	var outStream *os.File
	if o == "" {
		outStream = os.Stdout
	} else {
		os.Create(o)
		outStream, _ = os.Open(o)
	}
	defer outStream.Close()

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = outStream
	cmd.Stdout = outStream

	util.PL(t.TaskId, "Starting")

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	util.PL(t.TaskId, "Finishing")

	return "Finished task " + strconv.Itoa(t.TaskId), nil
}
