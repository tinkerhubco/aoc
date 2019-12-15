package main

import (
   "fmt"
   "math"
   "strconv"
   "strings"
)

var input string = `<x=16, y=-11, z=2>
<x=0, y=-4, z=7>
<x=6, y=4, z=-10>
<x=-3, y=-2, z=-4>`

var steps = 1000

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
   Vel *Velocity
}

func main() {
   var parsedInput []string = strings.Split(input, "\n")
   var moons []*Moon = make([]*Moon, 0)
   
   for _, v := range parsedInput {
      raw := strings.Split(v, ",")
      rawX := strings.Split(raw[0], "=")
      rawY := strings.Split(raw[1], "=")
      rawZ := strings.Split(raw[2], "=")

      x, _ := strconv.Atoi(rawX[1])
      y, _ := strconv.Atoi(rawY[1])
      z, _ := strconv.Atoi(strings.Trim(rawZ[1], ">"))

      moon := &Moon{
         Pos: &Position{
            X: x,
            Y: y,
            Z: z,
         },
         Vel: &Velocity{},
      }

      moons = append(moons, moon)
   }

   for _, moon := range moons {
      fmt.Println("After 0 steps:")
      fmt.Printf(
         "pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>\n",
         moon.Pos.X,
         moon.Pos.Y,
         moon.Pos.Z,
         moon.Vel.X,
         moon.Vel.Y,
         moon.Vel.Z,
      )
   }

   for i := 1; i <= steps; i++ {
      fmt.Println()
      positions := make([]*Position, len(moons))
      for moonIndex, cMoon := range moons {
         for _, otherMoon := range moons {
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
         position := &Position{}
         position.X = cMoon.Pos.X + cMoon.Vel.X
         position.Y = cMoon.Pos.Y + cMoon.Vel.Y
         position.Z = cMoon.Pos.Z + cMoon.Vel.Z

         positions[moonIndex] = position
      }

      fmt.Printf("After %d steps:\n", i)
      for positionIndex, position := range positions {
         moons[positionIndex].Pos.X = position.X
         moons[positionIndex].Pos.Y = position.Y
         moons[positionIndex].Pos.Z = position.Z
         fmt.Printf(
            "pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>\n",
            moons[positionIndex].Pos.X,
            moons[positionIndex].Pos.Y,
            moons[positionIndex].Pos.Z,
            moons[positionIndex].Vel.X,
            moons[positionIndex].Vel.Y,
            moons[positionIndex].Vel.Z,
         )
      }
   }

   fmt.Println()
   fmt.Printf("Energy After %d steps:\n", steps)

   var rowTotalLabel []string
   energyTotal := 0

   for _, moon := range moons {
      px := int(math.Abs(float64(moon.Pos.X)))
      py := int(math.Abs(float64(moon.Pos.Y)))
      pz := int(math.Abs(float64(moon.Pos.Z)))

      vx := int(math.Abs(float64(moon.Vel.X)))
      vy := int(math.Abs(float64(moon.Vel.Y)))
      vz := int(math.Abs(float64(moon.Vel.Z)))

      pot := px + py + pz
      kin := vx + vy + vz
      rowTotal := pot * kin
      energyTotal += int(rowTotal)

      rowTotalStr :=  strconv.Itoa(rowTotal)
      rowTotalLabel = append(rowTotalLabel, rowTotalStr)

      fmt.Printf(
         "pot: %d + %d + %d =  %d;   kin: %d + %d + %d = %d;   total:  %d * %d = %d\n",
         px,
         py,
         pz,
         pot,
         vx,
         vy,
         vx,
         kin,
         pot,
         kin,
         rowTotal,
      )
   }

   fmt.Printf(
      "Sum of total energy: %s = %d",
      strings.Join(rowTotalLabel, " + "),
      energyTotal,
   )
}



