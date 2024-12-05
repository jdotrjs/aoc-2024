package main

import (
	"fmt"
	"sort"
)

type Grid struct {
	Data [][]byte
	DimX int
	DimY int
}

type Cell struct {
	Value byte
	X     int
	Y     int
	grid  *Grid
}

type Path []Cell

func (p Path) HasElement(c *Cell) bool {
	for _, e := range p {
		if e.X == c.X && e.Y == c.Y {
			return true
		}
	}
	return false
}

func (p Path) ToString() string {
	s := "["
	for i, c := range p {
		s += "{" + c.ToString() + "}"
		if i < len(p)-1 {
			s += ", "
		}
	}
	s += "]"
	return s
}

func stringToBytes(s string) []byte {
	return []byte(s)
}

func inputToGrid() Grid {
	g := Grid{Data: [][]byte{}}
	y := 0
	for bs := range processInput(stringToBytes) {
		gridLine := append([]byte{}, bs...)
		g.Data = append(g.Data, gridLine)
		y += 1
	}

	g.DimY = y
	g.DimX = len(g.Data[0])

	return g
}

func (g *Grid) Print() {
	fmt.Printf("Dim: %d x %d\n", g.DimX, g.DimY)
	fmt.Printf("Data:\n")
	for y := 0; y < len(g.Data); y++ {
		fmt.Printf("  ")
		for x := 0; x < len(g.Data[y]); x++ {
			fmt.Printf("%c ", g.Data[y][x])
		}
		fmt.Println()
	}
}

func (g *Grid) At(x, y int) (Cell, bool) {
	if y >= g.DimY || y < 0 {
		return Cell{}, false
	}
	if x >= g.DimX || x < 0 {
		return Cell{}, false
	}
	return Cell{g.Data[y][x], x, y, g}, true
}

func (c *Cell) ToString() string {
	return fmt.Sprintf("Cell(%d, %d): %c", c.X, c.Y, c.Value)
}

func (c *Cell) FindXmas(path Path) []Path {
	// fmt.Printf("FindXmas - %s - %s\n", c.ToString(), path.ToString())
	if c.Value == 'S' {
		return []Path{append(path, *c)}
	}

	want := c.NextLetterXMas()

	results := []Path{}

	for _, surround := range c.Surrounding() {
		if surround.Value == want {
			results = append(
				results,
				surround.FindXmas(append(path, *c))...,
			)
		}
	}

	return results
}

func (c *Cell) FindXmasVector(path Path, dx, dy int) []Path {
	if c.Value == 'S' {
		return []Path{append(path, *c)}
	}

	want := c.NextLetterXMas()

	results := []Path{}

	if next, ok := c.grid.At(c.X+dx, c.Y+dy); ok && next.Value == want {
		return append(results, next.FindXmasVector(append(path, *c), dx, dy)...)
	}
	return nil
}

func (c *Cell) Surrounding() []Cell {
	opts := []Cell{}
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if x == 0 && y == 0 {
				continue
			}
			if entry, ok := c.grid.At(c.X+x, c.Y+y); ok {
				opts = append(opts, entry)
			}
		}
	}
	return opts
}

func (c *Cell) NextLetterXMas() byte {
	switch c.Value {
	case 'X':
		return 'M'
	case 'M':
		return 'A'
	case 'A':
		return 'S'
	default:
		panic("nope")
	}
}

func partOne() {
	grid := inputToGrid()
	starts := []Cell{}

	for x := 0; x < grid.DimX; x++ {
		for y := 0; y < grid.DimY; y++ {
			v := MustOK(grid.At(x, y))
			if v.Value == 'X' {
				starts = append(starts, v)
			}
		}
	}

	count := 0

	for _, v := range starts {
		fmt.Printf("Start: %s\n", v.ToString())
		for x := -1; x < 2; x++ {
			for y := -1; y < 2; y++ {
				if x == 0 && y == 0 {
					continue
				}
				count += len(v.FindXmasVector(Path{}, x, y))
			}
		}
	}

	fmt.Printf("Part One: %d\n", count)
}

func partTwo() {
	grid := inputToGrid()
	starts := []Cell{}

	for x := 0; x < grid.DimX; x++ {
		for y := 0; y < grid.DimY; y++ {
			v := MustOK(grid.At(x, y))
			if v.Value == 'A' {
				starts = append(starts, v)
			}
		}
	}

	count := 0

	for _, v := range starts {
		top_left, tl_ok := grid.At(v.X-1, v.Y-1)
		top_right, tr_ok := grid.At(v.X+1, v.Y-1)
		bot_left, bl_ok := grid.At(v.X-1, v.Y+1)
		bot_right, br_ok := grid.At(v.X+1, v.Y+1)
		if !(tl_ok && tr_ok && bl_ok && br_ok) {
			continue
		}

		// fucking go ass shit right here
		diag1 := []byte(fmt.Sprintf("%c%c", top_left.Value, bot_right.Value))
		sort.Slice([]byte(diag1), func(i, j int) bool { return diag1[i] < diag1[j] })

		diag2 := []byte(fmt.Sprintf("%c%c", top_right.Value, bot_left.Value))
		sort.Slice(diag2, func(i, j int) bool { return diag2[i] < diag2[j] })

		if string(diag1) == "MS" && string(diag2) == "MS" {
			count += 1
		}
	}

	fmt.Printf("Part Two: %d\n", count)
}

func testSurround() {
	grid := inputToGrid()
	c := MustOK(grid.At(3, 3))
	for _, v := range c.Surrounding() {
		fmt.Println(v)
	}
}

func testSliceAppend() {
	g := inputToGrid()
	p := Path{MustOK(g.At(3, 3))}
	d := append(p, MustOK(g.At(3, 4)))
	p = append(p, MustOK(g.At(3, 5)))

	fmt.Println(p.ToString())
	fmt.Println(d.ToString())
}

func main() {
	// partOne()
	partTwo()
}
