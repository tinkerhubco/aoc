package main
// 89603079
import (
   "fmt"
   "strings"
   "strconv"
)

// var input = `3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10`
// var input = `3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5`
var input = `3,8,1001,8,10,8,105,1,0,0,21,30,47,60,81,102,183,264,345,426,99999,3,9,1002,9,5,9,4,9,99,3,9,1002,9,5,9,1001,9,4,9,1002,9,4,9,4,9,99,3,9,101,2,9,9,1002,9,4,9,4,9,99,3,9,1001,9,3,9,1002,9,2,9,101,5,9,9,1002,9,2,9,4,9,99,3,9,102,4,9,9,101,4,9,9,1002,9,3,9,101,2,9,9,4,9,99,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,99,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,99`

var used = make(map[int]bool)

func main() {

   var ip []string = strings.Split(input, ",")
   var o []int

   for _, v := range ip {
     n, _ := strconv.Atoi(v)
     o = append(o, n)
   }

   max  := 0
   f := false

   for i := 50000; i <= 98765 ; i++ {
     ii := strconv.Itoa(i)
     // if i < 10000 {
     //     ii = "0" + ii
     // }
     if !isValid(ii) {
      continue
     }

     if isUniq(ii) {

         ampAIP := 0
         ampBIP := 0
         ampCIP := 0
         ampDIP := 0
         ampEIP := 0

         ampAOutput := []int{}
         ampBOutput := []int{}
         ampCOutput := []int{}
         ampDOutput := []int{}
         ampEOutput := []int{ 0 }

         o1 := append(o[:0:0], o...)
         o2 := append(o[:0:0], o...)
         o3 := append(o[:0:0], o...)
         o4 := append(o[:0:0], o...)
         o5 := append(o[:0:0], o...)

         ampASetting, _ := strconv.Atoi(string(ii[0]))
         ampBSetting, _ := strconv.Atoi(string(ii[1]))
         ampCSetting, _ := strconv.Atoi(string(ii[2]))
         ampDSetting, _ := strconv.Atoi(string(ii[3]))
         ampESetting, _ := strconv.Atoi(string(ii[4]))

         for ;  ; {

            ampAOutput, ampAIP = get(o1, ampASetting, ampEOutput[0], ampAIP)
            if o1[ampAIP] == 99 {
               break
            }

            ampBOutput, ampBIP = get(o2, ampBSetting, ampAOutput[0], ampBIP)
            if o2[ampBIP] == 99 {
               break
            }

            ampCOutput, ampCIP = get(o3, ampCSetting, ampBOutput[0], ampCIP)
            if o3[ampCIP] == 99 {
               break
            }

            ampDOutput, ampDIP = get(o4, ampDSetting, ampCOutput[0], ampDIP)
            if o4[ampDIP] == 99 {
               break
            }

            ampEOutput, ampEIP = get(o5, ampESetting, ampDOutput[0], ampEIP)
            if o5[ampEIP] == 99 {
               break
            }
         }

         if (max < ampEOutput[0] || !f) {
            f = true
            max = ampEOutput[0]
         }
         
     }
     
   }   
   fmt.Print("Max:", max)

   fmt.Println()
}

func isUniq(str string) bool {
   m := make(map[string]int)
   for i := 0; i < len(str) ; i++ {
      // _, exists := m[string(str[i])]
      // if !exists {
      //    m[]
      // }
      
      m[string(str[i])]++
      if m[string(str[i])] > 1 {
         return false
      }

   }
      
   return true
}

func isValid(str string) bool {
   for i := 0; i < len(str) ; i++ {
      v := string(str[i])
      if v == "0" || v == "1" || v == "2" || v == "3" || v == "4" {
         return false
      }
   }     
   return true
}

func get(o []int, puts int, settings int, start int) ([]int, int) {
   var e int = len(o)
   var out []int = make([]int, 0)
   i := start
   for ; i < e; i++ {

      in := strconv.Itoa(o[i])
      l := len(in)

      opCode := -1
      mode1 := 0
      mode2 := 0

      if l < 5 {
         in = strings.Repeat("0", 5 - l) + in
      }

      opCode, _ = strconv.Atoi(in[3:])
      mode1, _ = strconv.Atoi(string(in[2]))
      mode2, _ = strconv.Atoi(string(in[1]))

      if opCode == 99 {
         break
      } else if opCode == 1 {
         p1 := o[i + 1]
         p2 := o[i + 2]
         p3 := o[i + 3]

         if mode1 == 0 {
            p1 = o[p1]
         }

         if mode2 == 0 {
            p2 = o[p2]
         }

         o[p3] = p1 + p2
         i+=3
      } else if opCode == 2 {
         p1 := o[i + 1]
         p2 := o[i + 2]
         p3 := o[i + 3]

         if mode1 == 0 {
            p1 = o[p1]
         }

         if mode2 == 0 {
            p2 = o[p2]
         }

         o[p3] = p1 * p2
         i+=3
      } else if opCode == 3 {
         p1 := o[i + 1]

         if i == 0  {
            o[p1] = puts
         }
         if i > 0 {
            o[p1] = settings
         }

         i+=1
      } else if opCode == 4 {
         p1 := o[i + 1]

         if mode1 == 0 {
            p1 = o[p1]
         }
         out = append(out, p1)
         i++
         return out, i
         // fmt.Print(p1)
      } else if opCode == 5 {
         p1 := o[i + 1]
         p2 := o[i + 2]

         if mode1 == 0 {
            p1 = o[p1]
         }

         if mode2 == 0 {
            p2 = o[p2]
         }

         if p1 != 0 {
            i = p2
            i--
         } else {
            i+=2
         }
      }  else if opCode == 6 {
         p1 := o[i + 1]
         p2 := o[i + 2]

         if mode1 == 0 {
            p1 = o[p1]
         }

         if mode2 == 0 {
            p2 = o[p2]
         }

         if p1 == 0 {
            i = p2
            i--
         } else {
            i+=2
         }
      }  else if opCode == 7 {
         p1 := o[i + 1]
         p2 := o[i + 2]
         p3 := o[i + 3]

         if mode1 == 0 {
            p1 = o[p1]
         }

         if mode2 == 0 {
            p2 = o[p2]
         }

         if p1 < p2 {
            o[p3] = 1
         } else {
            o[p3] = 0
         }
         i+=3
      } else if opCode == 8 {
         p1 := o[i + 1]
         p2 := o[i + 2]
         p3 := o[i + 3]

         if mode1 == 0 {
            p1 = o[p1]
         }

         if mode2 == 0 {
            p2 = o[p2]
         }

         if p1 == p2 {
            o[p3] = 1
         } else {
            o[p3] = 0
         }
         i+=3
      } else {
         // fmt.Print("invalid code")
      }

   }

   return out, i
}
