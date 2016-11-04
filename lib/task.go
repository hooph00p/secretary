package lib

import (
	"os"
	"os/exec"
	"strconv"

	"github.com/hooph00p/secretary/lib/util"
)

type Task struct {
	i int
}

func (t *Task) loc() string {
	return "./test/" + strconv.Itoa(t.i) + ".log"
}

func (t *Task) Run() (string, error) {
	file, err := os.Create(t.loc())
	defer file.Close()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("ping", "google.com")
	cmd.Stdout = file
	cmd.Stderr = file
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	util.PL(t.i, "Finishing")

	return "Finished task " + strconv.Itoa(t.i), nil
}
