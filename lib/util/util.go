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
	/**
	 * Cited: https://github.com/laurent22/massren/blob/ae4c57da1e09a95d9383f7eb645a9f69790dec6c/main.go#L172
	 */
	const (
		START  string = "START"
		QUOTES string = "QUOTES"
		ARG    string = "ARG"
	)

	var args []string
	var current string

	state := START
	quote := "\""

	/**
	* Iterate the String, character by character.
	 */
parsing:
	for i := 0; i < len(text); i++ {
		c := text[i]

		switch {

		case state == QUOTES: // Step through Quoted String
			if string(c) != quote {
				current += string(c)
			} else {
				args = append(args, current)
				current = ""
				state = START
			}
			continue parsing

		case c == '"' || c == '\'': // Initiate Quoted String Step-Through
			state = QUOTES
			quote = string(c)
			continue parsing

		case state == ARG: // Step through regular space/tab-delimited argument
			if c == ' ' || c == '\t' {
				args = append(args, current)
				current = ""
				state = START
			} else {
				current += string(c)
			}
			continue parsing

		case c != ' ' && c != '\t': // Start Argument State
			state = ARG
			current += string(c)

		}
	}

	if state == QUOTES {
		return args, errors.New(fmt.Sprintf("Unclosed quote in command line: %s", text))
	}

	if current != "" {
		args = append(args, current)
	}

	return args, nil
}
