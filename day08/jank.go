package main

import "fmt"

func (solver *day8Solver) markDistances_LinearJank() {
	for station := range solver.broadcast {
		for _, tx_pos := range solver.broadcast[station] {
			tx_cell := solver.AtVec(tx_pos)
			station := tx_cell.tx
			fmt.Printf("Marking %s\n", tx_cell.ToString())

			for _, dir := range []Vector{vec2(-1, -1), vec2(-1, 1), vec2(1, 1), vec2(1, -1)} {
				next_pos := tx_pos
				for solver.InRangeVec(next_pos) {
					if next_pos == tx_pos {
						next_pos = next_pos.Add(dir)
						continue
					}

					cell := solver.AtVec(next_pos)
					dx := AbsInt(tx_cell.pos.x - cell.pos.x)
					dy := AbsInt(tx_cell.pos.y - cell.pos.y)
					dist := dx + dy

					if cell.rx_distances[station] == nil {
						cell.rx_distances[station] = IntSet{}
					}
					cell.rx_distances[tx_cell.tx].Add(dist)

					next_pos = next_pos.Add(dir)
				}
			}
		}
	}
}

func (solver *day8Solver) markDistance_XYJank(tx_cell *Cell, x, y int) {
	if tx_cell.pos == vec2(x, y) {
		return
	}

	tgt_pos := vec2(x, y)
	if false && tgt_pos == vec2(0, 3) {
		fmt.Printf("%c -- %s vs %s\n", tx_cell.tx, tx_cell.pos.ToString(), tgt_pos.ToString())
	}

	dx := AbsInt(tx_cell.pos.x - x)
	dy := AbsInt(tx_cell.pos.y - y)
	dist := dx + dy
	if tgt_pos == vec2(0, 3) {
		fmt.Printf("  dx:   %d\n  dy:   %d\n  dist: %d\n", dx, dy, dist)
	}

	tgt_cell := solver.At(x, y)
	station := tx_cell.tx

	if tgt_cell.rx_distances[station] == nil {
		tgt_cell.rx_distances[station] = IntSet{}
	}
	tgt_cell.rx_distances[station].Add(dist)
}

func (solver *day8Solver) findAntiNodes() int {
	count := 0

	for x := 0; x < solver.dimX; x++ {
		for y := 0; y < solver.dimY; y++ {
			if solver.findAntiNodesVec(vec2(x, y)) {
				// fmt.Printf("%s\n", vec2(x, y).ToString())
				count += 1
			}
		}
	}

	return count
}

func (solver *day8Solver) findAntiNodesVec(pos Vector) bool {
	cell := solver.AtVec(pos)
	cell.Print()

	match := false

	for signal, dist_set := range cell.rx_distances {
		if false {
			fmt.Printf("Checking %c\n", signal)
		}
		for dist := range dist_set {
			if dist_set.Has(2 * dist) {
				// fmt.Printf("  %d & %d\n", dist, 2*dist)
				match = true
				return true
			}
		}
	}

	return match
}
