package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type ListEntry []int64

func processInput() chan ListEntry {
	output_chan := make(chan ListEntry)

	go func() {
		reader := bufio.NewReader(os.Stdin)
		var str string
		var e error

		str, e = reader.ReadString('\n')
		for e != io.EOF {
			output_chan <- stringToEntry(str)
			str, e = reader.ReadString('\n')
		}
		last_entry := stringToEntry(str)
		output_chan <- last_entry
		close(output_chan)
	}()

	return output_chan
}

func stringToEntry(s string) ListEntry {
	if !numbers.Match([]byte(s)) {
		return nil
	}

	n := numbers.FindStringSubmatch(s)

	return ListEntry([]int64{
		Must(strconv.ParseInt(n[1], 10, 64)),
		Must(strconv.ParseInt(n[2], 10, 64)),
	})
}

func Must[T any](t T, e error) T {
	if e != nil {
		log.Panic(e)
	}

	return t
}

var numbers *regexp.Regexp = Must(regexp.Compile(`(\d+)\s+(\d+)`))

func part1() {
	var leftList []int
	var rightList []int

	for e := range processInput() {
		if e == nil {
			break
		}
		leftList = append(leftList, int(e[0]))
		rightList = append(rightList, int(e[1]))
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	dist := 0

	for i, lv := range leftList {
		rv := rightList[i]
		dist += int(math.Abs((float64)(lv - rv)))
	}

	fmt.Printf("day 1.1: %v\n", dist)
}

func part2() {
	var leftList []int
	rightLookup := map[int64]int{}

	for e := range processInput() {
		if e == nil {
			break
		}
		leftList = append(leftList, int(e[0]))
		v := rightLookup[e[1]]
		rightLookup[e[1]] = v + 1
	}

	score := 0
	for _, v := range leftList {
		score += v * rightLookup[int64(v)]
	}

	fmt.Printf("day 1.2: %v\n", score)
}

func main() {
	part2()
}
