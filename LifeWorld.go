package GoLife

type World struct {
  name        string
  xMax        int
  yMax        int
  ageMax      int
  immunityMax int
}

func (w *World) MaxImmunity() int {
  return w.immunityMax
}
