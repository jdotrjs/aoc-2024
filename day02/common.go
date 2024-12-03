package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func processInput[T any](processFunc func(string) T) chan T {
	output_chan := make(chan T)

	go func() {
		reader := bufio.NewReader(os.Stdin)
		var str string
		var e error

		for {
			str, e = reader.ReadString('\n')
			str = strings.TrimSpace(str)
			if e == io.EOF {
				break
			}

			if str == "" {
				continue
			}

			output_chan <- processFunc(str)
		}

		if str != "" {
			last_entry := processFunc(str)
			output_chan <- last_entry
		}
		close(output_chan)
	}()

	return output_chan
}

// TODO: same common lib
func Must[T any](t T, e error) T {
	if e != nil {
		log.Panic(e)
	}

	return t
}
