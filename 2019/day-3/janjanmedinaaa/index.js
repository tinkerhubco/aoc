Array.prototype.hasMin = function(attrib) {
  return (this.length && this.reduce(function(prev, curr){
    return prev[attrib] < curr[attrib] ? prev : curr;
  })) || null;
}

const lineIntersection = (pointA, pointB, pointC, pointD) => {
  var z1 = (pointA.x - pointB.x);
  var z2 = (pointC.x - pointD.x);
  var z3 = (pointA.y - pointB.y);
  var z4 = (pointC.y - pointD.y);
  var dist = z1 * z4 - z3 * z2;
  if (dist == 0) {
    return null;
  }
  var tempA = (pointA.x * pointB.y - pointA.y * pointB.x);
  var tempB = (pointC.x * pointD.y - pointC.y * pointD.x);
  var xCoor = (tempA * z2 - z1 * tempB) / dist;
  var yCoor = (tempA * z4 - z3 * tempB) / dist;

  if (xCoor < Math.min(pointA.x, pointB.x) || xCoor > Math.max(pointA.x, pointB.x) ||
    xCoor < Math.min(pointC.x, pointD.x) || xCoor > Math.max(pointC.x, pointD.x)) {
    return null;
  }
  if (yCoor < Math.min(pointA.y, pointB.y) || yCoor > Math.max(pointA.y, pointB.y) ||
    yCoor < Math.min(pointC.y, pointD.y) || yCoor > Math.max(pointC.y, pointD.y)) {
    return null;
  }

  return { x: xCoor, y: yCoor }
}

// Get Manhattan Distance
const manhattan = (point1, point2) => {
  return {
    manhattan: Math.abs(point1.x - point2.x) + Math.abs(point1.y - point2.y),
    point1,
    point2
  }
}

// Check if Point is inside two Points
const checkCoordinate = (point1, point2, coor) => {
    var slope = (point2.y - point1.y) / (point2.x - point1.x);
    var newSlope = (point2.y - coor.y) / (point2.x - coor.x);
    if (coor.x > point1.x && coor.x < point2.x && coor.y > point1.y && coor.y < point2.y && slope == newSlope) {
        return true
    } else {
        return false
    }
}

const checkGoal = (point, goal) => {
  if (goal === null) return false
  else return point.x === goal.x && point.y === goal.y
}

const mover = (dir, currMoves, moves, currPoint, goal = null) => {
  var newMoves = []
  switch (dir) {
    case 'L':
      var a = 0;
      while(a < moves) {
        currMoves += 1
        var x = currPoint.x - 1
        var y = currPoint.y
        var point = { x, y }
        newMoves.push(point)
        currPoint = point
        if (checkGoal(point, goal)) {
          return { currMoves, currPoint, goal: true }
        }
        a++
      }
      break;
    case 'R':
      var a = 0;
      while(a < moves) {
        currMoves += 1
        var x = currPoint.x + 1
        var y = currPoint.y
        var point = { x, y }
        newMoves.push(point)
        currPoint = point
        if (checkGoal(point, goal)) {
          return { currMoves, currPoint, goal: true }
        }
        a++
      }
      break;
    case 'U':
      var a = 0;
      while(a < moves) {
        currMoves += 1
        var x = currPoint.x
        var y = currPoint.y + 1
        var point = { x, y }
        newMoves.push(point)
        currPoint = point
        if (checkGoal(point, goal)) {
          return { currMoves, currPoint, goal: true }
        }
        a++
      }
      break;
    case 'D':
      var a = 0;
      while(a < moves) {
        currMoves += 1
        var x = currPoint.x
        var y = currPoint.y - 1
        var point = { x, y }
        newMoves.push(point)
        currPoint = point
        if (checkGoal(point, goal)) {
          return { currMoves, currPoint, goal: true }
        }
        a++
      }
      break;
  }

  return { newMoves, currMoves, currPoint, goal: false }
}

const singleLooper = (data, goal = null) => {
  var split = data.split(',')
  var current = { x: 0, y: 0 }
  var moves = 0
  var points = []
  points.push(current)

  for (var a = 0; a < split.length; a++) {
    var e = split[a]
    var length = parseInt(e.substring(1))
    var moverResult = mover(e[0], moves, length, current, goal)
    if (moverResult.goal) {
      return moverResult.currMoves
    } else {
      moves = moverResult.currMoves
      points.push(moverResult.currPoint)
      current = moverResult.currPoint
    }
  }

  return points
}

const checker = (points1, points2) => {
  var common = []
  for (var a = 0; a < points1.length-1; a++) {
    // var move1 = { start: points1[a], end: points1[a+1] }
    for (var b = 0; b < points2.length-1; b++) {
      // var move2 = { start: points2[b], end: points2[b+1] }

      var insect = lineIntersection(points1[a], points1[a+1], points2[b], points2[b+1])
      if (insect !== null) {
        common.push(insect)
      }
    }
  }
  return common
}

const getDistance = (match) => {
  var distances = []
  match.forEach(p => {
    var central = { x: 0,y: 0 }
    distances.push(manhattan(central, p))
  })
  return distances
}


var first = singleLooper(data1)
var second = singleLooper(data2)
var matches = checker(first, second)
// console.log(getDistance(matches).hasMin('manhattan'))
var results = []
matches.forEach((match) => {
  var step1 = singleLooper(data1, match)
  var step2 = singleLooper(data2, match)
  console.log(step1, step2)
  results.push(step1+step2)
})

console.log(Math.min(...results))