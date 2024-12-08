package main

import "fmt"

type Cell struct {
	pos          Vector
	rx_distances map[byte]IntSet
	tx           byte
}

func NewCell(x, y int, contents byte) *Cell {
	return &Cell{
		pos:          vec2(x, y),
		rx_distances: map[byte]IntSet{},
		tx:           contents,
	}
}

func (c *Cell) HasTx() bool {
	return c.tx != '.'
}

func (c *Cell) ToString() string {
	tx_str := byte('_')
	if c.HasTx() {
		tx_str = c.tx
	}
	return fmt.Sprintf("Cell%s tx: %c", c.pos.ToString(), tx_str)
}

func (c *Cell) Print() {
	fmt.Printf("Cell%s\n", c.pos.ToString())
	fmt.Printf("  Txn: %c\n", c.tx)
	fmt.Printf("  Distances:\n")
	for k, v := range c.rx_distances {
		fmt.Printf("    - %c -> %v\n", k, v.Keys())
	}
}
