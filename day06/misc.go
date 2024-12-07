package main

import (
	"fmt"
	"regexp"
)

type Pair[T any, U any] struct {
	L T
	R U
}

func pair[T any, U any](l T, r U) Pair[T, U] {
	return Pair[T, U]{L: l, R: r}
}

type Facing int

type Vector struct {
	x int
	y int
}

func (v Vector) ToString() string {
	return fmt.Sprintf("(%d, %d)", v.x, v.y)
}

func (v Vector) Add(other Vector) Vector {
	return vec(v.x+other.x, v.y+other.y)
}

func vec(x, y int) Vector {
	return Vector{x, y}
}

const (
	UP    Facing = 0
	RIGHT Facing = 1
	DOWN  Facing = 2
	LEFT  Facing = 3

	OPEN     = 10
	OBSTACLE = 11
)

func byteToFacing(c byte) Facing {
	switch c {
	case '^':
		return UP
	case '>':
		return RIGHT
	case 'v':
		return DOWN
	case '<':
		return LEFT
	default:
		panic("nope")
	}
}

func facingToByte(i Facing) byte {
	switch i {
	case UP:
		return '^'
	case DOWN:
		return 'v'
	case LEFT:
		return '<'
	case RIGHT:
		return '>'
	default:
		panic("nope")
	}
}

func facingToVector(f Facing) Vector {
	switch f {
	case UP:
		return Vector{0, -1}
	case DOWN:
		return Vector{0, 1}
	case RIGHT:
		return Vector{1, 0}
	case LEFT:
		return Vector{-1, 0}
	default:
		panic("nope")
	}
}

var guard_regexp *regexp.Regexp = regexp.MustCompile(`[\^v<>]`)

func is_guard(b byte) bool {
	return guard_regexp.Match([]byte{b})
}

func is_open(b byte) bool {
	return b == '.'
}

func is_obstacle(b byte) bool {
	return b == '#'
}
