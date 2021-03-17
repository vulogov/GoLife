package GoLife

import (
  "fmt"
  "errors"
)

type World struct {
  name              string
  xMax              int
  yMax              int
  ageMax            int
  immunityMax       int
  immunityRestored  bool
  procreateFunc     func() bool
  data              [][]Cell
}

func NewWorld(name string, xsize int, ysize int, agemax int, immmax int) *World {
  var world *World
  world = new(World)
  world.name    = name
  world.xMax    = xsize
  world.yMax    = ysize
  world.ageMax  = agemax
  world.immunityMax = immmax
  world.immunityRestored = false
  world.procreateFunc = func () bool { return false }
  world.CreateWorld()
  return world
}

func (w *World) CreateWorld() {
  w.data = make([][]Cell, 0)
  for x := 0; x < w.xMax; x++ {
    tmp := make([]Cell, 0)
    for y := 0; y < w.yMax; y++ {
      tmp = append(tmp, NewCell(w, x, y))
    }
    w.data = append(w.data, tmp)
  }
}

func (w* World) Cell(x int, y int) (*Cell, error) {
  if x > w.xMax && x < 0 {
    return nil, errors.New(fmt.Sprintf("X dimention is out of bounds: %v", x))
  }
  if y > w.yMax && y < 0 {
    return nil, errors.New(fmt.Sprintf("Y dimention is out of bounds: %v", x))
  }
  return &w.data[x][y], nil
}

func (w* World) GetCell(x int, y int) *Cell {
  cell, err := w.Cell(x,y)
  if err != nil {
    return nil
  }
  return cell
}

func (w *World) MaxImmunity() int {
  return w.immunityMax
}

func (w *World) X() int {
  return w.xMax
}

func (w *World) Y() int {
  return w.yMax
}

func (w *World) ImmunityRecovery() {
  w.immunityRestored = true
}

func (w *World) NoImmunityRecovery() {
  w.immunityRestored = false
}
