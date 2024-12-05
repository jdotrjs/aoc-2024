package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Set []int

func (s Set) Contains(i int) bool {
	for _, v := range s {
		if i == v {
			return true
		}
	}
	return false
}

type inputParser[T any] struct {
	beforeRules     map[int]Set // key must come before
	afterRules      map[int]Set // key must come after
	finishedRuleDef bool
	partOne         int
	partTwo         int
	stubValue       T
}

func newInputParser[T any](stub T) inputParser[T] {
	return inputParser[T]{
		beforeRules:     map[int]Set{},
		afterRules:      map[int]Set{},
		finishedRuleDef: false,
		partOne:         0,
		partTwo:         0,
		stubValue:       stub,
	}
}

func (self inputParser[T]) PrintRules() {
	fmt.Printf("Before Rules:\n")
	for k, v := range self.beforeRules {
		v_strs := []string{}
		for _, e := range v {
			v_strs = append(v_strs, fmt.Sprintf("%d", e))
		}
		fmt.Printf("  %d -> %s\n", k, strings.Join(v_strs, ", "))
	}

	fmt.Printf("After Rules:\n")
	for k, v := range self.afterRules {
		v_strs := []string{}
		for _, e := range v {
			v_strs = append(v_strs, fmt.Sprintf("%d", e))
		}
		fmt.Printf("  %d -> %s\n", k, strings.Join(v_strs, ", "))
	}

}

func (self *inputParser[T]) Do(s string) T {
	if s == "" {
		self.finishedRuleDef = true
		// self.PrintRules()
		return self.stubValue
	}
	if self.finishedRuleDef {
		valid := self.lineValid(s)
		if valid {
			parts := strings.Split(s, ",")
			self.partOne += Must(strconv.Atoi(parts[len(parts)/2]))
		} else {
			newLine := self.reorder(s)
			self.partTwo += newLine[len(newLine)/2]
		}

		return self.stubValue
	}

	parts := strings.Split(s, "|")
	p0 := Must(strconv.Atoi(parts[0]))
	p1 := Must(strconv.Atoi(parts[1]))

	self.beforeRules[p0] = append(self.beforeRules[p0], p1)
	self.afterRules[p1] = append(self.afterRules[p1], p0)
	return self.stubValue
}

func (self *inputParser[T]) lineValid(s string) bool {
	values := map[int][]int{}
	for i, part := range strings.Split(s, ",") {
		part_key := Must(strconv.Atoi(part))
		values[part_key] = append(values[part_key], i)
	}

	for key, locations := range values {
		// rules list all numbers that must come after key
		rules := self.beforeRules[key]
		for _, key_loc := range locations {
			// key_loc is a singular locations of key

			for _, check_value := range rules {
				for _, value_location := range values[check_value] {
					if value_location < key_loc {
						return false
					}
				}
			}
		}
	}
	return true
}

func (self *inputParser[T]) reorder(s string) []int {
	result := []int{}
	input := Set{}

	for _, v := range strings.Split(s, ",") {
		input = append(input, Must(strconv.Atoi(v)))
	}

	for len(input) > 0 {
		selected := -1

		for i, check_val := range input {
			input_clear := true

			for _, after_vals := range self.afterRules[check_val] {
				input_clear = input_clear && !input.Contains(after_vals)
			}

			if input_clear {
				selected = i
				break
			}
		}

		if selected == -1 {
			panic("wat")
		}
		result = append(result, input[selected])
		input = append(input[0:selected], input[selected+1:]...)
	}

	return result
}

func main() {
	ip := newInputParser(0)

	for range processInput(&ip, true) {
	}

	fmt.Printf("Part One: %d\n", ip.partOne)
	fmt.Printf("Part Two: %d\n", ip.partTwo)
}
