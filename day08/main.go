package main

import (
	"fmt"
)

type day8Solver struct {
	part_one_anode map[Vector]bool
	part_two_anode map[Vector]bool
	only_tx        byte
	grid           map[Vector]Cell
	dimX           int
	dimY           int
	cur_line       int
	broadcast      map[byte][]Vector
}

var _ LineProcessor[int] = &day8Solver{}

func (solver *day8Solver) InRange(x, y int) bool {
	return solver.InRangeVec(vec2(x, y))
}

func (solver *day8Solver) InRangeVec(v Vector) bool {
	return v.x >= 0 && v.x < solver.dimX && v.y >= 0 && v.y < solver.dimY
}

func (solver *day8Solver) At(x, y int) *Cell {
	return solver.AtVec(vec2(x, y))
}

func (solver *day8Solver) AtVec(pos Vector) *Cell {
	if !solver.InRange(pos.x, pos.y) {
		return nil
	}
	c := solver.grid[pos]
	return &c
}

func (solver *day8Solver) Do(s string) int {
	for x, b := range s {
		new_cell := *NewCell(x, solver.cur_line, byte(b))
		if new_cell.HasTx() {
			if solver.only_tx != 0 && new_cell.tx != solver.only_tx {
				new_cell.tx = '.'
			} else {
				if solver.broadcast[new_cell.tx] == nil {
					solver.broadcast[new_cell.tx] = []Vector{}
				}
				solver.broadcast[new_cell.tx] = append(solver.broadcast[new_cell.tx], new_cell.pos)
			}
		}
		solver.grid[new_cell.pos] = new_cell
	}
	if solver.dimX < len(s) {
		solver.dimX = len(s)
	}
	solver.cur_line += 1
	solver.dimY = solver.cur_line
	return 0
}

func (solver *day8Solver) markDistancesPairwise() {
	for _, posns := range solver.broadcast {
		for idx_a, pos_a := range posns {
			for idx_b := idx_a + 1; idx_b < len(posns); idx_b++ {
				pos_b := posns[idx_b]
				solver.markDistancePair(solver.AtVec(pos_a), solver.AtVec(pos_b))
			}
		}
	}
}

func (solver *day8Solver) markDistancePair(fixed_cell *Cell, cell_b *Cell) {
	delta := cell_b.pos.Sub(fixed_cell.pos)

	fixed_pos := fixed_cell.pos.Sub(delta)
	pos_2 := cell_b.pos.Add(delta)

	if solver.InRangeVec(fixed_pos) {
		solver.part_one_anode[fixed_pos] = true
	}
	if solver.InRangeVec(pos_2) {
		solver.part_one_anode[pos_2] = true
	}

	aoeu := map[Vector]bool{}

	for _, dir := range []Vector{delta, delta.Mult(vec_n1n1)} {
		cur_pos := fixed_cell.pos
		for solver.InRangeVec(cur_pos) {
			aoeu[cur_pos] = true
			cur_pos = cur_pos.Add(dir)
		}
	}

	for k := range aoeu {
		solver.part_two_anode[k] = true
	}
}

func vec_dist(va, vb Vector) int {
	dx := va.x - vb.x
	dy := va.y - vb.y
	return AbsInt(dx) + AbsInt(dy)
}

func (s *day8Solver) PrintGrid() {
	for y := 0; y < s.dimY; y++ {
		for x := 0; x < s.dimX; x++ {
			c := s.At(x, y)
			if c.HasTx() {
				fmt.Printf("%c", c.tx)
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}

}

func (s *day8Solver) PrintAntiNodes() {
	for y := 0; y < s.dimY; y++ {
		for x := 0; x < s.dimX; x++ {
			if s.part_two_anode[vec2(x, y)] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func partOne() {
	solver := day8Solver{
		part_one_anode: map[Vector]bool{},
		part_two_anode: map[Vector]bool{},
		// only_tx:   'j',
		grid:      map[Vector]Cell{},
		broadcast: map[byte][]Vector{},
	}
	for range processInput(&solver, false) {
	}

	solver.markDistancesPairwise()

	// solver.findAntiNodes()
	fmt.Printf("Part One_b: %d\n", len(solver.part_one_anode))

	p2 := 0

	for y := 0; y < solver.dimY; y++ {
		for x := 0; x < solver.dimX; x++ {
			c := solver.At(x, y)
			for _, dists := range c.rx_distances {
				if len(dists) > 1 {
					p2 += 1
				}
			}
		}
	}

	fmt.Printf("Part Two: %d\n", len(solver.part_two_anode))
	solver.PrintGrid()
	fmt.Println()
	solver.PrintAntiNodes()
}

func main() {
	partOne()
}
