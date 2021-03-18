package GoLife

import (
  "fmt"
  "math/rand"
)

type Cell struct {
  x          int
  y          int
  name       string
  status     bool
  degrading  bool
  opImmGiven int
  upImmGiven int
  opImmunity int
  upImmunity int
  age        int
  world      *World
}


func NewCell(w *World, x, y int) Cell {
  var c Cell
  c.x = x
  c.y = y
  c.world = w
  c.ToDead()
  return c
}

func (c *Cell) X() int {
  return c.x
}

func (c *Cell) Y() int {
  return c.y
}

func (c *Cell) Neighbors() []*Cell {
  var x1,x2,y1,y2 int
  var n []*Cell
  w := c.world
  if c.x == 0 {
    x1 = w.X() - 1
    x2 = c.x + 1
  } else {
    x1 = c.x - 1
    x2 = c.x + 1
    if x2 == w.X() {
      x2 = 0
    }
  }
  if c.y == 0 {
    y1 = w.Y() - 1
    y2 = c.y + 1
  } else {
    y1 = c.y - 1
    y2 = c.y + 1
    if y2 == w.Y() {
      y2 = 0
    }
  }
  n = append(n, w.GetCell(x1,c.y))
  n = append(n, w.GetCell(x1,y2))
  n = append(n, w.GetCell(c.x,y2))
  n = append(n, w.GetCell(x2,y2))
  n = append(n, w.GetCell(x2,c.y))
  n = append(n, w.GetCell(x2,y1))
  n = append(n, w.GetCell(c.x,y1))
  n = append(n, w.GetCell(x1,y1))
  return n
}

func (c *Cell) SetName(name string)  {
  c.name = name
}

func (c *Cell) Name() string {
  return c.name
}

func (c *Cell) Age() int {
  return c.age
}

func (c *Cell) Alive() bool {
  return c.status
}


func (c *Cell) ToDead() bool {
  c.age         = 0
  c.status      = false
  c.opImmunity  = 0
  c.upImmunity  = 0
  c.opImmGiven  = 0
  c.upImmGiven  = 0
  c.degrading   = false
  // Tell in notificationCh that I am DEAD
  c.world.NotificationCh <- *c
  return true
}

func (c *Cell) ToLife() bool {
  w             := c.world
  c.age         = 0
  c.status      = true
  c.opImmunity  = rand.Intn(w.MaxImmunity())
  c.upImmunity  = rand.Intn(w.MaxImmunity())
  c.opImmGiven  = c.opImmunity
  c.upImmGiven  = c.upImmunity
  c.degrading   = false
  // Tell in notificationCh that I am LIVE
  c.world.NotificationCh <- *c
  return true
}

func (c *Cell) Step() bool {
  c.age++
  n := c.Neighbors()
  aliveNeighbors := 0
  for _, nc := range n {
    if nc.Alive() == true {
      aliveNeighbors += 1
    }
  }
  if c.status == true {
    if c.world.ageMax < c.age {
      // Shall we die of the old age ? Absolutely !
      c.ToDead()
      return c.status
    }
    if aliveNeighbors == 3 {
      c.degrading = false
    }
    if aliveNeighbors < 3 {
      // Underpopulation !!!
      c.upImmunity -= 1
      if c.upImmunity <= 0 {
        // and immunity is off
        c.ToDead()
        return c.status
      }
      c.degrading = true
    }
    if aliveNeighbors > 3 {
      // Overpopulation !!!
      c.opImmunity -= 1
      if c.opImmunity <= 0 {
        // and immunity is off
        c.ToDead()
        return c.status
      }
      c.degrading = true
    }
    // See if we can restore immunity if condition is favorable
    if c.upImmGiven > c.upImmunity && c.world.immunityRestored == true {
      c.upImmunity += 1
    }
    if c.opImmGiven > c.opImmunity && c.world.immunityRestored == true {
      c.opImmunity += 1
    }
  } else {
    // Cell is dead
    if c.world.ageMax < c.age {
      // Shall we procreate life ?
      if c.world.procreateFunc() == true {
        c.ToLife()
        return c.status
      }
    }
    if aliveNeighbors == 3 {
      // 3 neighboring Live cells make a Dead cell Live
      c.ToLife()
      return c.status
    }
  }
  return c.status
}

func (c *Cell) Degrading() bool {
  return c.degrading
}

func (c *Cell) String() string {
  var status string
  var deg string
  if c.Alive() {
    status = "LIVE"
  } else {
    status = "DEAD"
  }
  if c.Degrading() {
    deg = "*"
  } else {
    deg = " "
  }
  return fmt.Sprintf("[%v%v] Name=%s; X=%v; Y=%v; Age=%v", status, deg, c.name, c.x, c.y, c.age)
}
