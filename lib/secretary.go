package lib

import (
	"bufio"
	"os"
	"strings"

	"github.com/hooph00p/secretary/lib/util"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

const (
	VERSION = "0.0.1-alpha"
)

var (
	app = kingpin.New("secretary", "A Scheduling Service. Web GUI and Command Line.")

	run = app.Command("run", "Run the Secretary Shell.")
)

type Secretary struct {
	Tasks []*Task
}

func (s *Secretary) Run() {
	reader := bufio.NewReader(os.Stdin)
	messages := make(chan string)

	// gen
	for i := 0; i < 5; i++ {
		t := &Task{i: i}
		s.Tasks = append(s.Tasks, t)
	}

	// Listener
	go func() {
		for {
			msg := <-messages
			util.LOG(msg)
		}
	}()

	for {
		// prompt
		util.P("Enter: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		args, _ := util.Parse(text)

		// parse
		switch kingpin.MustParse(app.Parse(args)) {
		case run.FullCommand():
			for i := range s.Tasks {
				go func(t *Task) {
					msg, err := t.Run()
					util.PL("m,e:", msg, err)
					if err == nil {
						messages <- msg
					} else {
						messages <- err.Error()
					}
				}(s.Tasks[i])
			}
		}
	}
}
