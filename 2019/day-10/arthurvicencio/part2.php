<?php

$input = '.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##';

$input = '.#....#####...#..
##...##.#####..##
##...#...#.#####.
..#.....#...###..
..#.#.....#....##';

$input = '.###.#...#.#.##.#.####..
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
 
$pasrsed = explode(PHP_EOL, $input);
$mmap = [];


foreach ($pasrsed as $row) {
   $mmap[] = str_split($row);
}

$max = '';
$maxIndex = [];
$m = null;
$omap = [];

foreach ($mmap as $i => $r) {
   foreach ($r as $j => $v) {
   // var_dump($mmap[$i][$j]);
      // die;
      if ($mmap[$i][$j] !== '.') {
         // $track = [];
         $omap[$i][$j] = countInSight($j, $i);
         // echo $omap[$i][$j];
         if ($m < $omap[$i][$j] || $m === null) {
            $m = $omap[$i][$j];
            $maxIndex = [$i, $j];
            $max = "{$j},{$i} = {$m}";
         }
      } else {
         // echo '.';
         $omap[$i][$j] = 0;
      }
   }
   // echo '<br>';
}
// die;
echo $max;
echo PHP_EOL;

$maxY = count($mmap) - 1;
$maxX = count($mmap[0]) - 1;


$nthX = 0;
$nthY = 0;
$points = [];
$done = [];
$target = 200;
$a = 0;

for ($r = 0; ; $r++) {
   countInSight($maxIndex[1], $maxIndex[0], false, false, true);

   $straightUp = getQuadrant($points, $maxIndex[1], 0, $maxIndex[1], $maxIndex[0]);
   $stop = laser($straightUp);
   if ($stop) {
      break;
   }

   $q1 = getQuadrant($points, $maxIndex[1] + 1, 0, $maxX, $maxIndex[0] - 1);
   sortByAngleFrom($q1, $maxIndex[1], $maxIndex[0]);
   $stop = laser($q1);
   if ($stop) {
      break;
   }

   $straightRight = getQuadrant($points, $maxIndex[1], $maxIndex[0], $maxX, $maxIndex[0]);
   $stop = laser($straightRight);
   if ($stop) {
      break;
   }

   $q2 = getQuadrant($points, $maxIndex[1] + 1, $maxIndex[0] + 1, $maxX, $maxY);
   sortByAngleFrom($q2, $maxIndex[1], $maxIndex[0]);
   $stop = laser($q2);
   if ($stop) {
      break;
   }

   $straightDown = getQuadrant($points, $maxIndex[1], $maxIndex[0], $maxIndex[1], $maxY);
   $stop = laser($straightDown);
   if ($stop) {
      break;
   }

   $q3 = getQuadrant($points, 0, $maxIndex[0] + 1, $maxIndex[1] - 1, $maxY);
   sortByAngleFrom($q3, $maxIndex[1], $maxIndex[0]);
   $stop = laser($q3);
   if ($stop) {
      break;
   }

   $straightLeft = getQuadrant($points, 0, $maxIndex[0], $maxIndex[1], $maxIndex[0]);
   $stop = laser($straightLeft);
   if ($stop) {
      break;
   }

   $q4 = getQuadrant($points, 0,0, $maxIndex[1] - 1, $maxIndex[0] - 1);
   sortByAngleFrom($q4, $maxIndex[1], $maxIndex[0]);
   $stop = laser($q4);
   if ($stop) {
      break;
   }
}

echo "{$nthX},{$nthY}", PHP_EOL;
echo "X * 100 + Y = ", $nthX * 100 + $nthY, PHP_EOL;

foreach ($mmap as $row) {
   foreach ($row as $v) {
      echo $v;
   }
   echo PHP_EOL;
}

die;

function sortByAngleFrom(&$points, $centerX, $centerY) {
   $ts = [];
   foreach ($points as $p => $v) {
      $p = explode(",", $p);
      $x = $p[1];
      $y = $p[0];
      $dx = $x - $centerX;
      $dy = $y - $centerY;
      $theta = atan2($dy, $dx);
      $ts["{$y},{$x}"] = $theta;
   }
   asort($ts);
   $points = $ts;
}

function laser($css) {

   global $done, $mmap, $a, $target, $nthY, $nthX;

   foreach ($css as $val => $csss) {
      $p = explode(",", $val);
      $x = $p[1];
      $y = $p[0];

      if (isset($done["{$y},{$x}"])) {
         continue;
      }

      $done["{$y},{$x}"] = true;
      $mmap[$y][$x] = '.';
      $a++;
      if ($a === $target) {
         $nthY = $y;
         $nthX = $x;
         return true;
      }
   }
   return false;
}

function getQuadrant($points, $x1, $y1, $x2, $y2) {
   $part = [];

   for ($x = $x1; $x <= $x2; $x++) {
      for ($y = $y1; $y <= $y2; $y++) {
         if (isset($points["{$y},{$x}"])) {
            $part["{$y},{$x}"] = [$y, $x];
         }
      }  
   }

   return $part;
}

die;

function countInSight($x, $y, $p = false, $d = false, $store = false) {
   global $mmap, $points;
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
               if ($store) {
                  $points["{$i},{$j}"] = true;
               }
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