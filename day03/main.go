package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var cmd_extraction *regexp.Regexp = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
var numbers *regexp.Regexp = regexp.MustCompile(`\d+`)

func processLine(s string) []Command {
	r := []Command{}
	parts := cmd_extraction.FindAllStringSubmatch(s, -1)
	for _, v := range parts {
		r = append(r, CommandFromString(v[0]))
	}
	return r
}

type Command interface {
	Type() string
}

type command struct {
	typ string
}

func (c command) Type() string {
	return c.typ
}

func CommandFromString(s string) Command {
	if strings.HasPrefix(s, "mul(") {
		return MulCommandFromString(s)
	}
	return command{typ: s}
}

type MulCommand struct {
	command
	x int
	y int
}

func MulCommandFromString(s string) Command {
	parts := REextract(numbers, s)
	return MulCommand{
		command: command{typ: "mul"},
		x:       Must(strconv.Atoi(parts[0])),
		y:       Must(strconv.Atoi(parts[1])),
	}
}

func partOne() {
	total := 0
	for r := range processInput(processLine) {
		for _, c := range r {
			if c.Type() != "mul" {
				continue
			}
			mc := c.(MulCommand)
			total += mc.x * mc.y
		}
	}
	fmt.Printf("Part One: %d\n", total)
}

func partTwo() {
	total := 0
	working := true
	for r := range processInput(processLine) {
		for _, c := range r {
			if c.Type() == "do()" {
				working = true
			}
			if c.Type() == "don't()" {
				working = false
			}

			if working && c.Type() == "mul" {
				mc := c.(MulCommand)
				total += mc.x * mc.y
			}
		}
	}
	fmt.Printf("Part Two: %d\n", total)
}

func main() {
	// partOne()
	partTwo()
}
