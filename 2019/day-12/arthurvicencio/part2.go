package main

import (
   "fmt"
   // "math"
   "strconv"
   "strings"
)

var input string = `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`

type Position struct {
   X int
   Y int
   Z int
}

type Velocity struct {
   X int
   Y int
   Z int
}

type Moon struct {
   Pos *Position
   NewPos *Position
   Vel *Velocity
}

func main() {
   var parsedInput []string = strings.Split(input, "\n")
   var moonsX []*Moon = make([]*Moon, 0)
   var moonsY []*Moon = make([]*Moon, 0)
   var moonsZ []*Moon = make([]*Moon, 0)

   for _, v := range parsedInput {
      raw := strings.Split(v, ",")
      rawX := strings.Split(raw[0], "=")
      rawY := strings.Split(raw[1], "=")
      rawZ := strings.Split(raw[2], "=")

      x, _ := strconv.Atoi(rawX[1])
      y, _ := strconv.Atoi(rawY[1])
      z, _ := strconv.Atoi(strings.Trim(rawZ[1], ">"))

      pos := &Position{
         X: x,
         Y: y,
         Z: z,
      }

      npos := &Position{
         X: x,
         Y: y,
         Z: z,
      }
         
      moon := &Moon{
         Pos: pos,
         NewPos: npos,
         Vel: &Velocity{},
      }

      moonsX = append(moonsX, moon)
      moonsY = append(moonsY, moon)
      moonsZ = append(moonsZ, moon)
   }

   x := findFirstRepeating(moonsX, "X")
   y := findFirstRepeating(moonsY, "Y")
   z := findFirstRepeating(moonsZ, "Z")

   fmt.Println(LCM(x, y, z))
}

func findFirstRepeating(moon []*Moon, axis string) int64 {
   var seen map[string]bool = make(map[string]bool)
   var steps int64 = 0
   for ;  ; steps++ {

      for _, cMoon := range moon {
         for _, otherMoon := range moon {
            if cMoon == otherMoon {
               continue
            }
            if cMoon.Pos.X < otherMoon.Pos.X {
               cMoon.Vel.X++
            }
            if cMoon.Pos.X > otherMoon.Pos.X {
               cMoon.Vel.X--  
            }
            if cMoon.Pos.Y < otherMoon.Pos.Y {
               cMoon.Vel.Y++
            }
            if cMoon.Pos.Y > otherMoon.Pos.Y {
               cMoon.Vel.Y--  
            }
            if cMoon.Pos.Z < otherMoon.Pos.Z {
               cMoon.Vel.Z++
            }
            if cMoon.Pos.Z > otherMoon.Pos.Z {
               cMoon.Vel.Z--  
            }
         }
         cMoon.NewPos.X = cMoon.Pos.X + cMoon.Vel.X
         cMoon.NewPos.Y = cMoon.Pos.Y + cMoon.Vel.Y
         cMoon.NewPos.Z = cMoon.Pos.Z + cMoon.Vel.Z
      }

      stepKey := ""
      for _, cMoon := range moon {
         cMoon.Pos.X = cMoon.NewPos.X
         cMoon.Pos.Y = cMoon.NewPos.Y
         cMoon.Pos.Z = cMoon.NewPos.Z

         if axis == "X" {
            px := strconv.Itoa(cMoon.Pos.X)
            cx := strconv.Itoa(cMoon.Vel.X)
            stepKey = stepKey + px + "," + cx + ","
         }

         if axis == "Y" {
            py := strconv.Itoa(cMoon.Pos.Y)
            cy := strconv.Itoa(cMoon.Vel.Y)
            stepKey = stepKey + py + "," + cy + ","
         }

         if axis == "Z" {
            pz := strconv.Itoa(cMoon.Pos.Z)
            cz := strconv.Itoa(cMoon.Vel.Z)
            stepKey = stepKey + pz + "," + cz + ","
         }
      }

      if _, exists := seen[stepKey]; exists {
         break
      }

      seen[stepKey] = true
   }

   return steps
}

func GCD(a, b int64) int64 {
   for b != 0 {
      t := b
      b = a % b
      a = t
   }
   return a
}

func LCM(a, b int64, integers ...int64) int64 {
   result := a * b / GCD(a, b)

   for i := 0; i < len(integers); i++ {
      result = LCM(result, integers[i])
   }

   return result
}
