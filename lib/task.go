package lib

import "os"

type Task struct {
	Command    *string
	Interval   *int
	Repeat     *int
	OutputFile **os.File
}
