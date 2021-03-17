package main

import (
  "fmt"
  "github.com/vulogov/GoLife"
)

func main() {
  world := GoLife.NewWorld("world", 10, 10, 120, 10)
  fmt.Println(world)
  cell  := world.GetCell(4,4)
  fmt.Println(cell.String())
  for _, c := range cell.Neighbors() {
    fmt.Println(c.String())
  }
}
