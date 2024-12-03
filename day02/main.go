package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Report []int

// TODO: pull into common lib
func processInput[T any](processFunc func(string) T) chan T {
	output_chan := make(chan T)

	go func() {
		reader := bufio.NewReader(os.Stdin)
		var str string
		var e error

		str, e = reader.ReadString('\n')
		str = strings.Clone(str)
		str = strings.TrimSpace(str)
		for e != io.EOF {
			if str == "" {
				continue
			}
			output_chan <- processFunc(str)
			str, e = reader.ReadString('\n')
			str = strings.TrimSpace(str)
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

func stringToReport(s string) Report {
	parts := strings.Split(s, " ")

	r := Report{}
	for _, v := range parts {
		r = append(r, (int)(Must(strconv.ParseInt(v, 10, 32))))
	}

	return r
}

func (r Report) PartOneSafe() bool {
	if len(r) <= 1 {
		return true
	}

	increase := (r[1] - r[0]) > 0
	for i := 1; i < len(r); i++ {
		last := r[i-1]
		cur := r[i]
		delta := cur - last
		if increase {
			if delta < 1 || delta > 3 {
				return false
			}
		} else {
			if delta < -3 || delta > -1 {
				return false
			}
		}
	}

	return true
}

func (r Report) PartTwoSafe() bool {
	if r.PartOneSafe() {
		return true
	}

	// if we failed but we only have two elements removing one makes us trivially
	// pass
	if len(r) == 2 {
		return true
	}

	// at least three elements, we know we've failed for some reason

	deltas := []int{}
	for i := 1; i < len(r); i++ {
		deltas = append(deltas, r[i]-r[i-1])
	}

	// can no longer trust r1 - r0 ot indicate direction
	up := 0
	down := 0
	for _, d := range deltas {
		if d < 0 {
			down += 1
		}
		if d > 0 {
			up += 1
		}
	}

	increase := up > down

	// find the delta that breaks the rules
	check_left := -1
	check_right := -1

	for i := 0; i < len(deltas); i++ {
		delta := deltas[i]
		if increase {
			if delta > 3 || delta < 1 {
				check_left = i
				check_right = i + 1
				break
			}
		} else {
			if delta < -3 || delta > -1 {
				check_left = i
				check_right = i + 1
				break
			}
		}
	}

	if check_left == -1 || check_right == -1 {
		log.Panicf("wat: %v\n%v\n", r, deltas)
	}

	left_report := slices.Clone(r[0:check_left])
	left_report = append(left_report, r[check_left+1:]...)

	right_report := slices.Clone(r[0:check_right])
	right_report = append(right_report, r[check_right+1:]...)

	// if it works with either of the deltas removed it's fine; if not then it doesn't pass muster
	return left_report.PartOneSafe() || right_report.PartOneSafe()
}

func partOne() {
	safeReports := 0
	reportChan := processInput(stringToReport)
	for r := range reportChan {
		if r.PartOneSafe() {
			safeReports += 1
		}
	}
	fmt.Printf("Part 1 Safe Reports: %d\n", safeReports)
}

func partTwo() {
	safeReports := 0
	reportChan := processInput(stringToReport)
	for r := range reportChan {
		if r.PartTwoSafe() {
			safeReports += 1
		}
	}
	fmt.Printf("Part 2 Safe Reports: %d\n", safeReports)
}

func main() {
	partTwo()
	// Report([]int{52, 60, 63, 66, 66, 66}).DayTwoSafe()
	// test := []int{0, 1, 2, 3, 4, 5, 6}
	// fmt.Printf("%v\n%v\n", test, append(test[0:2], test[3:]...))
}
