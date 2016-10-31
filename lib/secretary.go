package lib

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

const (
	VERSION = "0.0.1-alpha"
)

var (
	app           = kingpin.New("secretary", "A Scheduling Service. Web GUI and Command Line.")
	shell         = app.Command("shell", "Run the Secretary Shell.")
	shellCommand  = shell.Flag("command", "Command String.").Required().String()
	shellInterval = shell.Flag("interval", "Interval in Seconds. Must be positive.").Default("1").Int()
	shellRepeat   = shell.Flag("repeat", "Times to repeat. Set negative to repeat infinitely.").Default("1").Int()
	shellOutfile  = shell.Flag("outfile", "Output file.").File()
)

type Secretary struct {
	Tasks []Task
}

func (s *Secretary) Run() {
	app.Version(VERSION)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		args := strings.SplitN(text, " ", 2)
		switch kingpin.MustParse(app.Parse(args)) {
		case shell.FullCommand():
			s.shell()
		}
	}
}

func (s *Secretary) shell() {

	command := *shellCommand
	P("Command:", command)
	if command == "" {
		panic("Uh... you gonna give me a command?")
	}

	interval := *shellInterval
	if interval != 0 {
		P("Interval:", interval)
	}
	if interval < 0 {
		panic("Interval must be positive.")
	}

	repeat := *shellRepeat
	if repeat != 0 {
		P("Repeat:", repeat)
	}
	if repeat < 0 {
		P("Infinitely Repeating.")
	}

	var outf *os.File
	if *shellOutfile != nil {
		P("Outfile:", *shellOutfile)
		outf = *shellOutfile
	}
	defer outf.Close()

	ticker := time.NewTicker(time.Second * time.Duration(interval))
	doneChan := make(chan bool)
	repeated := 0
	run := func(t time.Time) {
		repeated += 1

		fmt.Println("")
		if repeat > 0 {
			P("Iteration:", repeated, "of", strconv.Itoa(repeat))
		} else {
			P("Iteration:", repeated, "of Infinite.")
		}
		P("Time:", t)

		args := strings.Split(command, " ")
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if outf != nil {
			cmd.Stdout = outf
			cmd.Stderr = outf
		}

		if outf != nil {
			outf.WriteString("Iteration:\t" + strconv.Itoa(repeated))
			outf.WriteString("\nTime:\t" + t.String())
			outf.WriteString("\n" + strings.Repeat("-", 100) + "\n")
		}

		cmd.Run()

		if outf != nil {
			outf.WriteString(strings.Repeat("-", 100) + "\n")
		}

		if repeated == repeat {
			doneChan <- true
		}
	}

	// Establish and run the goroutine.
	go func() {
		if repeat != 0 {
			run(time.Now())
		}
		if repeat != 1 {
			for t := range ticker.C {
				run(t)
			}
		}
	}()

	// Infinite loop that listens
	// on the done channel for the
	// completion of the goroutine.
	for {
		select {
		case <-doneChan:
			P("Completed task.")
			return
		}
	}
}
