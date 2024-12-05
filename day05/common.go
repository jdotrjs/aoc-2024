package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

type LineProcessor[T any] interface {
	Do(s string) T
}

type simpleProcessor[T any] struct {
	processFunc func(string) T
}

func (sp simpleProcessor[T]) Do(s string) T {
	return sp.processFunc(s)
}

func processInputSimple[T any](processFunc func(string) T) chan T {
	return processInput(simpleProcessor[T]{processFunc}, false)
}

func processInput[T any](processor LineProcessor[T], pass_newlines bool) chan T {
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

			if str == "" && !pass_newlines {
				continue
			}

			output_chan <- processor.Do(str)
		}

		if str != "" {
			last_entry := processor.Do(str)
			output_chan <- last_entry
		}
		close(output_chan)
	}()

	return output_chan
}

func Must[T any](t T, e error) T {
	if e != nil {
		log.Panic(e)
	}

	return t
}

func MustOK[T any](t T, ok bool) T {
	if !ok {
		log.Panic("Expected ok but got not ok")
	}
	return t
}

func REextract(re *regexp.Regexp, s string) []string {
	r := []string{}
	for _, v := range re.FindAllStringSubmatch(s, -1) {
		if len(v) != 1 {
			log.Panicf("Expected one submatch in %v for %v, found %d", re, s, len(v))
		}
		r = append(r, v[0])
	}
	return r
}
