package main

import (
	"fmt"
)

type daySixProcessor struct {
	dimX      int
	dimY      int
	obstacles map[string]bool
	guard     guard_data

	partOnePath []Pair[Vector, Facing]

	loopingNodes map[Pair[Vector, Facing]]bool
	exitingNodes map[Pair[Vector, Facing]]bool

	cur_line int
}

var _ LineProcessor[int] = &daySixProcessor{}

func (self *daySixProcessor) Do(s string) int {
	self.dimX = len(s)
	for x, b := range s {
		if is_guard(byte(b)) {
			posn := vec(x, self.cur_line)
			self.guard = guard_data{
				cur:    posn,
				start:  posn,
				facing: byteToFacing(byte(b)),
			}
		}
		if is_obstacle(byte(b)) {
			self.obstacles[vec(x, self.cur_line).ToString()] = true
		}
	}

	self.cur_line += 1
	self.dimY = self.cur_line
	return 0
}

func (self *daySixProcessor) AtByte(x, y int) byte {
	switch self.At(x, y) {
	case OBSTACLE:
		return 'X'
	case OPEN:
		return '.'
	default:
		panic("nope")
	}
}

func (self *daySixProcessor) At(x, y int) int {
	p := vec(x, y)
	if self.obstacles[p.ToString()] {
		return OBSTACLE
	}
	return OPEN
}

func (self *daySixProcessor) Print() {
	fmt.Printf("Grid:\n")
	for y := 0; y < self.dimY; y++ {
		for x := 0; x < self.dimX; x++ {
			if vec(x, y) == self.guard.cur {
				fmt.Printf("%c", facingToByte(self.guard.facing))
			} else {
				fmt.Printf("%c", self.AtByte(x, y))
			}
		}
		fmt.Println()
	}

	fmt.Printf("\nObstacles:\n")
	for k := range self.obstacles {
		fmt.Printf("  %s\n", k)
	}

	fmt.Printf("\nGuard:\n")
	fmt.Printf("  Start:  %s\n", self.guard.start.ToString())
	fmt.Printf("  Cur:    %s\n", self.guard.cur.ToString())
	fmt.Printf("  Facing: %c\n", facingToByte(self.guard.facing))
}

func (self *daySixProcessor) InRange(loc Vector) bool {
	x := loc.x
	y := loc.y
	return x >= 0 && x < self.dimX && y >= 0 && y < self.dimY
}

func (self *daySixProcessor) partOne() int {
	visited := map[Vector]struct{}{}
	for self.InRange(self.guard.cur) {
		// fmt.Printf("%s %c\n", self.guard.cur.ToString(), facingToByte(self.guard.facing))
		visited[self.guard.cur] = struct{}{}
		next_loc := self.guard.next_step()
		if self.At(next_loc.x, next_loc.y) == OBSTACLE {
			self.guard.TurnRight()
		} else {
			self.guard.cur = next_loc
			self.partOnePath = append(self.partOnePath, pair(self.guard.cur, self.guard.facing))
		}
	}

	return len(visited)
}

func (self *daySixProcessor) exits() bool {
	visited := map[Pair[Vector, Facing]]bool{}

	for self.InRange(self.guard.cur) {
		marker := pair(self.guard.cur, self.guard.facing)
		if visited[marker] {
			return false
		}
		visited[marker] = true

		next_loc := self.guard.next_step()
		if self.At(next_loc.x, next_loc.y) == OBSTACLE {
			self.guard.TurnRight()
		} else {
			self.guard.cur = next_loc
		}
	}
	return true
}

func (self *daySixProcessor) partTwo() int {
	loops_found := map[Vector]struct{}{}
	used_path := map[Vector]bool{}

	for check_idx := 0; check_idx < len(self.partOnePath)-1; check_idx++ {
		guard_location := self.partOnePath[check_idx]
		used_path[guard_location.L] = true
		obs_location := self.partOnePath[check_idx+1].L

		if obs_location == self.guard.start {
			continue
		}
		if !self.InRange(obs_location) {
			continue
		}
		if self.At(obs_location.x, obs_location.y) == OBSTACLE {
			continue
		}
		if self.obstacles[obs_location.ToString()] {
			continue
		}
		if used_path[obs_location] {
			continue
		}

		self.obstacles[obs_location.ToString()] = true
		self.guard.cur = guard_location.L
		self.guard.facing = guard_location.R
		if !self.exits() {
			loops_found[obs_location] = struct{}{}
		}

		delete(self.obstacles, obs_location.ToString())
	}

	return len(loops_found)
}

func main() {
	p := &daySixProcessor{
		obstacles:   map[string]bool{},
		partOnePath: []Pair[Vector, Facing]{},
	}

	for range processInput(p, false) {
	}

	fmt.Printf("Part One: %d\n", p.partOne())
	fmt.Printf("Part Two: %d\n", p.partTwo())
}
