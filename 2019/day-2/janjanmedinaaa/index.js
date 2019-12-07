const singleLooper = (data, noun, verb) => {
  var copy = [...data]
  copy[1] = noun
  copy[2] = verb

  for(var a = 0; a < copy.length; a += 4) {
    var opcode = copy[a]
    switch (opcode) {
      case 1: 
        copy[copy[a+3]] = copy[copy[a+1]] + copy[copy[a+2]]
        break;
      case 2:
        copy[copy[a+3]] = copy[copy[a+1]] * copy[copy[a+2]]
        break;
      default:
        return copy
    }
  }

  return copy
}

const greatLooper = (data) => {
  var replicate = [...data]

  for(var a = 1; a < 100; a++) {
    for(var b = 1; b < 100; b++) {
      var result = singleLooper(replicate, a, b)
      if (result[0] === 19690720) {
        return 100 * a + b
      }
    }
  }
}

console.log(singleLooper(RAW, 12, 1)[0])
console.log(greatLooper(RAW))