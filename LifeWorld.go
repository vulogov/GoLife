package GoLife

type World struct {
  name        string
  xMax        int
  yMax        int
  ageMax      int
  immunityMax int
  data        [][]Cell
}

func NewWorld(name string, xsize int, ysize int, agemax int, immmax int) *World {
  var *world = new(World)
  world.name    = name
  world.xMax    = xsize
  world.yMax    = ysize
  world.ageMax  = agemax
  world.immunityMax = immmax
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

func (w *World) MaxImmunity() int {
  return w.immunityMax
}
