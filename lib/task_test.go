package lib_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hooph00p/secretary/lib"
	"github.com/hooph00p/secretary/lib/util"
)

func TestRun(t *testing.T) {
	messages := make(chan string)
	doneChan := make(chan bool)
	count := 0

	go func() {
		for {
			msg := <-messages
			count += 1
			util.PL(msg)
			if count == 5 {
				doneChan <- true
			}
		}
	}()

	for i := 0; i < 5; i++ {
		go func(id int) {
			task := lib.Task{TaskId: id}
			msg, err := task.Run([]string{"ping", "google.com"}, loc(task))
			if err != nil {
				panic(err)
			}
			messages <- msg
		}(i)
	}

	done := <-doneChan
	if done {
		fmt.Println("done")
	}

}

func loc(t lib.Task) string {
	return "./test/" + strconv.Itoa(t.TaskId) + ".log"
}
