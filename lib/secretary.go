package lib

import kingpin "gopkg.in/alecthomas/kingpin.v2"

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
	// reader := bufio.NewReader(os.Stdin)
	// messages := make(chan string)
	//
	// // Listener
	// go func() {
	// 	for {
	// 		msg := <-messages
	// 		util.PL(msg)
	// 	}
	// }()

	// for {
	// prompt
	// util.P("Enter: ")
	// text, _ := reader.ReadString('\n')
	// text = strings.TrimSpace(text)
	// args, _ := util.Parse(text)

	// parse
	// switch kingpin.MustParse(app.Parse(args)) {
	// case run.FullCommand():
	// 	// for i := range s.Tasks {
	// 	// 	go func(t *Task) {
	// 	// 		msg, err := t.Run([]string{"ping", "google.com"})
	// 	// 		util.PL("m,e:", msg, err)
	// 	// 		if err == nil {
	// 	// 			messages <- msg
	// 	// 		} else {
	// 	// 			messages <- err.Error()
	// 	// 		}
	// 	// 	}(s.Tasks[i])
	//
	// }
	// }
}
