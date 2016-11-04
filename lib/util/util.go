package util

import (
	"errors"
	"fmt"
	"os"
)

const (
	outfile string = "log/out.log"
	errfile string = "log/err.log"
)

func P(z ...interface{}) {
	fmt.Print(format(z...)...)
}

func PL(z ...interface{}) {
	fmt.Println(format(z...)...)
}

func PN() {
	fmt.Println("")
}

func LOG(z ...interface{}) {
	f, err := os.Open(outfile)
	if err != nil {
		f = os.Create(outfile)
	}
}

func format(z ...interface{}) (x []interface{}) {
	x = append(x, "[ Secretary ] ")
	for i := range z {
		x = append(x, z[i])
	}
	return
}

func Parse(text string) ([]string, error) {
	var args []string
	state := "start"
	current := ""
	quote := "\""
	for i := 0; i < len(text); i++ {
		c := text[i]

		if state == "quotes" {
			if string(c) != quote {
				current += string(c)
			} else {
				args = append(args, current)
				current = ""
				state = "start"
			}
			continue
		}

		if c == '"' || c == '\'' {
			state = "quotes"
			quote = string(c)
			continue
		}

		if state == "arg" {
			if c == ' ' || c == '\t' {
				args = append(args, current)
				current = ""
				state = "start"
			} else {
				current += string(c)
			}
			continue
		}

		if c != ' ' && c != '\t' {
			state = "arg"
			current += string(c)
		}
	}

	if state == "quotes" {
		return []string{}, errors.New(fmt.Sprintf("Unclosed quote in command line: %s", text))
	}

	if current != "" {
		args = append(args, current)
	}

	if len(args) <= 0 {
		return []string{}, errors.New("Empty command line")
	}

	if len(args) == 1 {
		return args, nil
	}

	return args, nil
}
