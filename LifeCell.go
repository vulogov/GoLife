package GoLife

import (
  "math/rand"
)

type Cell struct {
  x          int
  y          int
  name       string
  status     bool
  opImmunity int
  upImmunity int
  age        int
}


func (c *Cell) X() int {
  return c.x
}

func (c *Cell) Y() int {
  return c.y
}

func (c *Cell) SetName(name string)  {
  return c.name = name
}

func (c *Cell) Name() string {
  return c.name
}

func (c *Cell) Age() int {
  return c.age
}

func (c *Cell) ToDead() bool {
  c.age         = 0
  c.status      = false
  c.opImmunity  = 0
  c.upImmunity  = 0
}

func (c *Cell) ToLife(w *World) int {
  c.age         = 0
  c.status      = true
  c.opImmunity  = rand.Intn(w.MaxImmunity())
  c.upImmunity  = rand.Intn(w.MaxImmunity())
}
