package main

import (
  "fmt"
  "math/rand"
  "github.com/vulogov/GoLife"
)

func main() {
  world := GoLife.NewWorld("world", 5, 5, 120, 10)
  world.Procreate(func() bool {
    if rand.Intn(10) == 0 {
      return true
    }
    return false
  })
  world.ToLife(4,0)
  world.ToLife(4,1)
  world.ToLife(4,2)
  world.ToLife(10,10)
  for {
    world.Print()
    world.Step()
    if len(world.NotificationCh) > 0 {
      fmt.Printf("World is changing on step [%v]\n", world.Age())
      for len(world.NotificationCh) > 0 {
        cell := <- world.NotificationCh
        fmt.Println(cell.String())
      }
    }
  }
}
