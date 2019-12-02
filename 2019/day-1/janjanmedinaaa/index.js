const list = require('./data.json')
const partOne = (data) => {
  var sum = 0
  data.forEach(element => {
    sum += Math.floor(parseInt(element) / 3) - 2
  });

  return sum
}

const partTwo = (data) => {
  var sum = 0
  data.forEach(element => {
    var singleFuel = 0
    var lastFuel = element

    while (true) {
      var fuel = Math.floor(parseInt(lastFuel) / 3) - 2
      if (fuel <= 0) {
        break;
      } else {
        singleFuel += fuel
        lastFuel = fuel
      }
    }

    sum += singleFuel
  });

  return sum
}