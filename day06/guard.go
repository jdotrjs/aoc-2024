package main

type guard_data struct {
	cur    Vector
	facing Facing
	start  Vector
}

func (g guard_data) next_step() Vector {
	return g.cur.Add(facingToVector(g.facing))
}

func (g *guard_data) TurnRight() {
	g.facing = (g.facing + 1) % 4
}
