package main

import (
	"fmt"
	"strconv"
	"strings"
)

type equation struct {
	id      int
	answer  int
	accum   int
	unbound []int
	eqn     string
	first   bool
}

var gt int = 0

func (e equation) solve() int {
	if len(e.unbound) == 0 {
		if e.accum == e.answer {
			// fmt.Printf("%d: %s\n", e.id, e.eqn)
			gt = gt + e.answer
			return e.answer
		}
		return 0
	}

	plus_ans := equation{
		e.id,
		e.answer,
		e.accum + e.unbound[0],
		e.unbound[1:],
		e.eqn + fmt.Sprintf(" + %d", e.unbound[0]),
		false,
	}.solve()
	if plus_ans == e.answer {
		return e.answer
	}

	if e.first {
		return 0
	}

	mult_ans := equation{
		e.id,
		e.answer,
		e.accum * e.unbound[0],
		e.unbound[1:],
		e.eqn + fmt.Sprintf(" * %d", e.unbound[0]),
		false,
	}.solve()
	return mult_ans
}

func (e equation) solveTwo() int {
	if len(e.unbound) == 0 {
		if e.accum == e.answer {
			// fmt.Printf("SolveTwo: %s\n", e.eqn)
			return e.answer
		}
		return 0
	}

	plus_ans := equation{
		e.id,
		e.answer,
		e.accum + e.unbound[0],
		e.unbound[1:],
		e.eqn + fmt.Sprintf(" + %d", e.unbound[0]),
		false,
	}.solveTwo()
	if plus_ans == e.answer {
		return e.answer
	}

	cat_number := Must(strconv.Atoi(fmt.Sprintf("%d%d", e.accum, e.unbound[0])))
	cat_ans := equation{
		e.id,
		e.answer,
		cat_number,
		e.unbound[1:],
		e.eqn + fmt.Sprintf(" || %d", e.unbound[0]),
		false,
	}.solveTwo()

	if cat_ans == e.answer {
		return e.answer
	}

	if e.first {
		return 0
	}

	mult_ans := equation{
		e.id,
		e.answer,
		e.accum * e.unbound[0],
		e.unbound[1:],
		e.eqn + fmt.Sprintf(" * %d", e.unbound[0]),
		false,
	}.solveTwo()
	return mult_ans
}

var equation_no int = 0
var seen map[int]bool

func partOneSolver(s string) int {
	if s == "" {
		return 0
	}

	s_parts := strings.Split(s, ":")
	answer := Must(strconv.Atoi(s_parts[0]))
	numbers := []int{}

	numbers_str := strings.Split(s_parts[1], " ")
	for _, e := range numbers_str {
		if e == "" {
			continue
		}
		numbers = append(numbers, Must(strconv.Atoi(e)))
	}

	equation_no = equation_no + 1

	return equation{
		id:      equation_no,
		answer:  answer,
		accum:   0,
		unbound: numbers,
		eqn:     fmt.Sprintf("%d = ", answer),
		first:   true,
	}.solve()
}

func partTwoSolver(s string) int {
	if s == "" {
		return 0
	}

	s_parts := strings.Split(s, ":")
	answer := Must(strconv.Atoi(s_parts[0]))
	numbers := []int{}

	numbers_str := strings.Split(s_parts[1], " ")
	for _, e := range numbers_str {
		if e == "" {
			continue
		}
		numbers = append(numbers, Must(strconv.Atoi(e)))
	}

	equation_no = equation_no + 1

	return equation{
		id:      equation_no,
		answer:  answer,
		accum:   0,
		unbound: numbers,
		eqn:     fmt.Sprintf("%d = ", answer),
		first:   true,
	}.solveTwo()
}

func partOne() int {
	total := 0
	for a := range processInputSimple(partOneSolver) {
		total += a
	}
	return total
}

func partTwo() int {
	total := 0
	for a := range processInputSimple(partTwoSolver) {
		total += a
	}
	return total
}

func main() {
	p1Total := 0 // partOne()
	// fmt.Printf("Part One: %d\n", p1Total)
	fmt.Printf("Part Two: %d\n", p1Total+partTwo())
}
