package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {

	// extend Println
	p := func(z ...interface{}) {
		var x []interface{}
		x = append(x, "[ Secretary ]")
		for i := range z {
			x = append(x, z[i])
		}
		fmt.Println(x...)
	}

	// Create pointers to flag memory addresses
	commandPtr := flag.String("command", "echo none", "command string")
	intervalPtr := flag.Int("interval", 1, "interval in seconds. Must be positive.")
	repeatPtr := flag.Int("repeat", 1, "times to repeat. Set to negative number to repeat infinitely.")
	outfilePtr := flag.String("outfile", "", "outfile")

	// Add values into given memory addresses
	flag.Parse()

	// Command Parsing
	command := *commandPtr
	if command == "" {
		panic("Need a command.")
	} else {
		p("Command:", command)
	}

	// Interval Parsing
	interval := *intervalPtr
	if interval != 0 {
		p("Interval:", interval)
	}

	// Interval validation
	if interval < 0 {
		panic("Interval must be positive.")
	}

	// Repeat Parsing
	repeat := *repeatPtr
	if repeat > 0 {
		p("Repeat:", repeat)
	}
	if repeat < 0 {
		p("Infinitely Repeating.")
	}

	// Outfile Parsing
	outfile := *outfilePtr
	var outf *os.File
	if outfile != "" {
		p("Outfile:", outfile)
		outf, _ = os.Create(outfile)
	}
	defer outf.Close()

	// Create Ticker based on Interval
	ticker := time.NewTicker(time.Second * time.Duration(interval))

	// Create channel to monitor the completion of the goroutine.
	doneChan := make(chan bool)

	// Create a variable to store the repetitions
	repeated := 0
	var rstr string
	if repeat > 0 {
		rstr = strconv.Itoa(repeat)
	} else {
		rstr = "[Infinite]"
	}

	// Declare the run function
	run := func(t time.Time) {
		repeated += 1

		// Output to the console
		fmt.Println("")
		p("Iteration:", repeated, "of", rstr)
		p("Time:", t)

		// Determine the command line arguments
		args := strings.Split(command, " ")

		// Execute the command line arguments
		cmd := exec.Command(args[0], args[1:]...)

		// Set default stdout to the console
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// If an output file is defined, set that as the channel for
		// errors and logging.
		if outf != nil {
			cmd.Stdout = outf
			cmd.Stderr = outf
		}

		// If outputting to file, Print some salt to distinguish iterations.
		if outf != nil {
			outf.WriteString("Iteration:\t" + strconv.Itoa(repeated))
			outf.WriteString("\nTime:\t" + t.String())
			outf.WriteString("\n" + strings.Repeat("-", 100) + "\n")
		}

		// Execute the command declared above.
		cmd.Run()

		// If outputting to file, Print the end of an iteration
		if outf != nil {
			outf.WriteString(strings.Repeat("-", 100) + "\n")
		}

		// Check to see if we send a signal to the program-ending channel.
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
			p("Completed task. Ending.")
			return
		}
	}

}
