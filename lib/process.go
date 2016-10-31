package lib

type Process struct {
	Command string
}

func Args() {
	// 	// Create pointers to flag memory addresses
	//
	// 	// Add values into given memory addresses
	// 	flag.Parse()
	//
	// 	// Command Parsing
	// 	proc.Command = *commandPtr
	// 	if proc.Command == "" {
	// 		panic("Need a command.")
	// 	} else {
	//
	// 		P("Command:", proc.Command)
	// 	}
	//
	// 	// Interval Parsing
	// 	proc.Interval = *intervalPtr
	// 	if proc.Interval != 0 {
	// 		P("Interval:", proc.Interval)
	// 	}
	//
	// 	// Interval validation
	// 	if proc.Interval < 0 {
	// 		panic("Interval must be positive.")
	// 	}
	//
	// 	// Repeat Parsing
	// 	proc.Repeat = *repeatPtr
	// 	if proc.Repeat > 0 {
	// 		P("Repeat:", proc.Repeat)
	// 	}
	// 	if proc.Repeat < 0 {
	// 		P("Infinitely Repeating.")
	// 	}
	//
	// 	// Outfile Parsing
	// 	proc.Outfile = *outfilePtr
	// 	var outf *os.File
	// 	if proc.Outfile != "" {
	// 		P("Outfile:", proc.Outfile)
	// 		outf, _ = os.Create(proc.Outfile)
	// 	}
	// 	defer outf.Close()
	//
	// }
	//
	// func Run() {
	//
	// 	// Create Ticker based on Interval
	// 	ticker := time.NewTicker(time.Second * time.Duration(interval))
	//
	// 	// Create channel to monitor the completion of the goroutine.
	// 	doneChan := make(chan bool)
	//
	// 	// Create a variable to store the repetitions
	// 	repeated := 0
	// 	var rstr string
	// 	if repeat > 0 {
	// 		rstr = strconv.Itoa(repeat)
	// 	} else {
	// 		rstr = "[Infinite]"
	// 	}
	//
	// 	// Declare the run function
	// 	run := func(t time.Time) {
	// 		repeated += 1
	//
	// 		// Output to the console
	// 		fmt.Println("")
	// 		P("Iteration:", repeated, "of", rstr)
	// 		P("Time:", t)
	//
	// 		// Determine the command line arguments
	// 		args := strings.Split(command, " ")
	//
	// 		// Execute the command line arguments
	// 		cmd := exec.Command(args[0], args[1:]...)
	//
	// 		// Set default stdout to the console
	// 		cmd.Stdout = os.Stdout
	// 		cmd.Stderr = os.Stderr
	//
	// 		// If an output file is defined, set that as the channel for
	// 		// errors and logging.
	// 		if outf != nil {
	// 			cmd.Stdout = outf
	// 			cmd.Stderr = outf
	// 		}
	//
	// 		// If outputting to file, Print some salt to distinguish iterations.
	// 		if outf != nil {
	// 			outf.WriteString("Iteration:\t" + strconv.Itoa(repeated))
	// 			outf.WriteString("\nTime:\t" + t.String())
	// 			outf.WriteString("\n" + strings.Repeat("-", 100) + "\n")
	// 		}
	//
	// 		// Execute the command declared above.
	// 		cmd.Run()
	//
	// 		// If outputting to file, Print the end of an iteration
	// 		if outf != nil {
	// 			outf.WriteString(strings.Repeat("-", 100) + "\n")
	// 		}
	//
	// 		// Check to see if we send a signal to the program-ending channel.
	// 		if repeated == repeat {
	// 			doneChan <- true
	// 		}
	// 	}
	//
	// 	// Establish and run the goroutine.
	// 	go func() {
	// 		if repeat != 0 {
	// 			run(time.Now())
	// 		}
	// 		if repeat != 1 {
	// 			for t := range ticker.C {
	// 				run(t)
	// 			}
	// 		}
	// 	}()
	//
	// 	// Infinite loop that listens
	// 	// on the done channel for the
	// 	// completion of the goroutine.
	// 	for {
	// 		select {
	// 		case <-doneChan:
	// 			PN()
	// 			P("Completed task. Ending.")
	// 			return
	// 		}
	// 	}

}
