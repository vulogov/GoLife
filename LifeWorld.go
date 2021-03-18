package GoLife

import (
  "os"
  "fmt"
  "time"
  "errors"
  "github.com/jedib0t/go-pretty/table"
)

type World struct {
  name              string
  xMax              int
  yMax              int
  ageMax            int
  immunityMax       int
  age               int
  immunityRestored  bool
  procreateFunc     func() bool
  stepFunc          func()
  NotificationCh    chan Cell
  data              [][]Cell
}

const (
  CHSIZE = 1000000
)

func NewWorld(name string, xsize int, ysize int, agemax int, immmax int) *World {
  var world *World
  world = new(World)
  world.name    = name
  world.xMax    = xsize
  world.yMax    = ysize
  world.ageMax  = agemax
  world.immunityMax = immmax
  world.age     = 0
  world.immunityRestored = false
  world.procreateFunc = func () bool { return false }
  world.stepFunc      = func () { time.Sleep(5 * time.Second) }
  world.NotificationCh = make(chan Cell, CHSIZE)
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

func (w *World) Step() {
  w.age++
  for row := 0; row < w.X(); row++  {
    for col := 0; col < w.Y(); col++ {
      w.data[row][col].Step()
    }
  }
  w.stepFunc()
}

func (w *World) Print() {
  var cellRepr string
  var degrading string

  t := table.NewWriter()
  t.SetOutputMirror(os.Stdout)
  for row := 0; row < w.X(); row++  {
    rowData := table.Row{}
    for col := 0; col < w.Y(); col++ {
      cell := w.GetCell(row,col)
      if cell.Degrading() == true {
        degrading = "*"
      } else {
        degrading = " "
      }
      if cell.Alive() {
        cellRepr = fmt.Sprintf("LIVE(%v,%v)[%v]%v",row,col,cell.Age(),degrading)
      } else {
        cellRepr = fmt.Sprintf("DEAD(%v,%v)[%v]%v",row,col,cell.Age(),degrading)
      }
      rowData = append(rowData, cellRepr)
    }
    t.AppendRow(rowData)
  }
  t.Render()
}

func (w *World) ToLife(x int, y int) bool {
  cell, err := w.Cell(x,y)
  if err != nil {
    return false
  }
  return cell.ToLife()
}

func (w *World) ToDead(x int, y int) bool {
  cell, err := w.Cell(x,y)
  if err != nil {
    return false
  }
  return cell.ToDead()
}

func (w *World) Procreate(p func() bool) {
  w.procreateFunc = p
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

func (w *World) Age() int {
  return w.age
}

func (w *World) ImmunityRecovery() {
  w.immunityRestored = true
}

func (w *World) NoImmunityRecovery() {
  w.immunityRestored = false
}
