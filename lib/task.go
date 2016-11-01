package lib

import "os"

type Task struct {
	Command    *string
	Interval   *int
	Repeat     *int
	OutputFile **os.File
}

func NewTask(addCommand *string, addInterval *int, addRepeat *int, addOutfile **os.File) *Task {
	command := *addCommand
	P("Command:", command)
	if command == "" {
		panic("Uh... you gonna give me a command?")
	}

	interval := *addInterval
	if interval != 0 {
		P("Interval:", interval)
	}
	if interval < 0 {
		panic("Interval must be positive.")
	}

	repeat := *addRepeat
	if repeat != 0 {
		P("Repeat:", repeat)
	}
	if repeat < 0 {
		P("Infinitely Repeating.")
	}

	var outf *os.File
	if *addOutfile != nil {
		P("Outfile:", *addOutfile)
		outf = *addOutfile
	}
	defer outf.Close()
}
