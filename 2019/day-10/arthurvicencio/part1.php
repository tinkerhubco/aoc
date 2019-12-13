<?php

$input = 
'.#..#
.....
#####
....#
...##';
$input = 
'#.........
...#......
...#..#...
.####....#
..#.#.#...
.....#....
..###.#.##
.......#..
....#...#.
...#..#..#';
$input = 
'.###.#...#.#.##.#.####..
.#....#####...#.######..
#.#.###.###.#.....#.####
##.###..##..####.#.####.
###########.#######.##.#
##########.#########.##.
.#.##.########.##...###.
###.#.##.#####.#.###.###
##.#####.##..###.#.##.#.
.#.#.#####.####.#..#####
.###.#####.#..#..##.#.##
########.##.#...########
.####..##..#.###.###.#.#
....######.##.#.######.#
###.####.######.#....###
############.#.#.##.####
##...##..####.####.#..##
.###.#########.###..#.##
#.##.#.#...##...#####..#
##.#..###############.##
##.###.#####.##.######..
##.#####.#.#.##..#######
...#######.######...####
#....#.#.#.####.#.#.#.##';

$p = explode(PHP_EOL, $input);
$mmap = [];
$omap = [];
$track = [];

foreach ($p as $dex => $i) {
   $str = str_split(trim($i));
   foreach ($str as $val) {
      $mmap[$dex][] = $val;
   }
}


$max = '';
$m = null;

foreach ($mmap as $i => $r) {
   foreach ($r as $j => $v) {
      if ($mmap[$i][$j] !== '.') {
         // $track = [];
         $omap[$i][$j] = countInSight($j, $i);
         // echo $omap[$i][$j];
         if ($m < $omap[$i][$j] || $m === null) {
            $m = $omap[$i][$j];
            $max = "{$j},{$i} = {$m}";
         }
      } else {
         // echo '.';
         $omap[$i][$j] = 0;
      }
   }
   // echo '<br>';
}

echo $max;

function countInSight($x, $y, $p = false, $d = false) {
   global $mmap;
   $inSight = 0;
   foreach ($mmap as $i => $r) {
      foreach ($r as $j => $v) {
         if ("{$y},{$x}" === "{$i},{$j}") {
            if ($p) echo '0';
            continue;
         }
         if ($mmap[$i][$j] !== '.') {
            // $track = [];
            if(checkInsight($x, $y, $j, $i)) {
               if ($p) echo '#';
               $inSight++;
            } else {
               // echo "{$i},{$j}", '<br>';
               if ($p) echo '.';
            }
         } else {
            if ($p) echo '.';
         }
      }
      if ($p) echo PHP_EOL;
   }
   if ($d) {
      die;
   }
   return $inSight;
}

function checkInsight($x, $y, $cx, $cy) {
   global $mmap;
   $track = [];

   $xDir = 1;
   $yDir = 1;

   if ($x > $cx) {
      $xDir = -1;
   }

   if ($y > $cy) {
      $yDir = -1;
   }

   // Horizontal
   $h = checkHorizontal($x, $y, $cx, $cy, $xDir, $yDir);
   if ($h) {
      return true;
   }

   // Vertical
   $v = checkVertial($x, $y, $cx, $cy, $xDir, $yDir);
   if ($v) {
      return true;
   }

   // Diagonal
   $d = checkDiagonal($x, $y, $cx, $cy, $xDir, $yDir);
   if ($d) {
      // echo "TRUE";
      return true;
   }

   $lenx = abs($x - $cx);
   if ($xDir === -1) {
      $lenx = -$lenx;
   }
   $leny = abs($y - $cy);
   if ($yDir === -1) {
      $leny = -$leny;
   }
   if ($lenx === $leny) {
      return false;
   }
   for ($ySkip = $yDir; ; $ySkip+=$yDir) {
      if ($yDir === -1) {
         if ($ySkip < $leny) {
            break;
         }
      } else if ($ySkip > $leny) {
         break;
      }
      for ($xSkip = $xDir; ; $xSkip+=$xDir) {
         if ($xDir === -1) {
            if ($xSkip < $lenx) {
               break;
            }
         } else if ($xSkip > $lenx) {
            break;
         }
         for ($yy = $y, $xx = $x; ; $yy+=$ySkip, $xx+=$xSkip) {
            if ($xDir === 1) {
               if ($xx > $cx) {
                  break;
               }
            }
            if ($xDir === -1) {
               if ($xx < $cx) {
                  break;
               }
            }
            if ($yDir === 1) {
               if ($yy > $cy) {
                  break;
               }
            }
            if ($yDir === -1) {
               if ($yy < $cy) {
                  break;
               }
            }

            if ("{$yy},{$xx}" !== "{$y},{$x}" && "{$yy},{$xx}" !== "{$cy},{$cx}" && $mmap[$yy][$xx] === '#') {
               if (!isset($track["{$y},{$x}-{$ySkip},{$xSkip}"])) {
                  $track["{$y},{$x}-{$ySkip},{$xSkip}"] = 0;
               }

               $track["{$y},{$x}-{$ySkip},{$xSkip}"]++;
            }

            if ($xx === $cx && $yy === $cy && $mmap[$yy][$xx] === '#') {

               if (!isset($track["{$y},{$x}-{$ySkip},{$xSkip}"])) {
                  $track["{$y},{$x}-{$ySkip},{$xSkip}"] = 0;
               }

               $track["{$y},{$x}-{$ySkip},{$xSkip}"]++;

               if ($track["{$y},{$x}-{$ySkip},{$xSkip}"] > 1) {
                  return false;
               }

               return true;
            }

         }

      }
   }

   return false;
}

function checkHorizontal($x, $y, $cx, $cy, $xDir, $yDir) {
   global $mmap;
   for ($xx = $x; $xx !== $cx + $xDir ; $xx+=$xDir) {
      if ($mmap[$y][$xx] === '#') {
         if ("{$y},{$xx}" !== "{$y},{$x}" && "{$y},{$xx}" !== "{$cy},{$cx}") {
            break;
         }
         if ("{$y},{$xx}" === "{$cy},{$cx}") {
            return true;
         }
      }
   }

   return false;
}

function checkVertial($x, $y, $cx, $cy, $xDir, $yDir) {
   global $mmap;
   for ($yy = $y; $yy !== $cy + $yDir ; $yy+=$yDir) {
      if ($mmap[$yy][$x] === '#') {
         if ("{$yy},{$x}" !== "{$y},{$x}" && "{$yy},{$x}" !== "{$cy},{$cx}") {
            break;
         }
         if ("{$yy},{$x}" === "{$cy},{$cx}") {
            return true;
         }
      }
   }
   return false;
}

function checkDiagonal($x, $y, $cx, $cy, $xDir, $yDir) {
   global $mmap;

   for ($yy = $y, $xx = $x; $yy !== $cy + $yDir && $xx !== $cx + $xDir ; $yy+=$yDir, $xx+=$xDir) {
      if ($mmap[$yy][$xx] === '#') {
         if ("{$yy},{$xx}" !== "{$y},{$x}" && "{$yy},{$xx}" !== "{$cy},{$cx}") {
            break;
         }
         if ("{$yy},{$xx}" === "{$cy},{$cx}") {
            return true;
         }
      }
   }
   return false;
}