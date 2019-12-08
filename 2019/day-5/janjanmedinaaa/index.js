const getOpcode = (instruction) => {
  var zeroes = '00000'.substr(0, 5-instruction.toString().length)
  instruction = `${zeroes}${instruction}`.split('')
  return {
    opcode: parseInt(instruction.slice(instruction.length-2).join('')),
    modes: instruction.map(i => parseInt(i))
  }
}

const singleLooper = (data) => {
  var copy = [...data]
  var jump = 1
  var input = 5
  var output = 0

  for(var a = 0; a < copy.length; a++) {
    var opcode = getOpcode(copy[a])
    var mode1 = opcode.modes[2]
    var mode2 = opcode.modes[1]

    if (opcode.opcode == 99) {
      break
    } else if (opcode.opcode == 1) {
      var value1 = (mode1 == 0) ? copy[copy[a+1]] : copy[a+1]
      var value2 = (mode2 == 0) ? copy[copy[a+2]] : copy[a+2]
      copy[copy[a+3]] = value1 + value2
      a += 3
    } else if (opcode.opcode == 2) {
      var value1 = (mode1 == 0) ? copy[copy[a+1]] : copy[a+1]
      var value2 = (mode2 == 0) ? copy[copy[a+2]] : copy[a+2]
      copy[copy[a+3]] = value1 * value2
      a += 3
    } else if (opcode.opcode == 3) {
      copy[copy[a+1]] = input
      a += 1
    } else if (opcode.opcode == 4) {
      var value1 = (mode1 == 0) ? copy[copy[a+1]] : copy[a+1]
      output = value1
      a += 1
    }  else if (opcode.opcode == 5) {
      var value1 = (mode1 == 0) ? copy[copy[a+1]] : copy[a+1]
      var value2 = (mode2 == 0) ? copy[copy[a+2]] : copy[a+2]
      if (value1 != 0) {
        a = value2-1
      } else {
        a += 2
      }
    } else if (opcode.opcode == 6) {
      var value1 = (mode1 == 0) ? copy[copy[a+1]] : copy[a+1]
      var value2 = (mode2 == 0) ? copy[copy[a+2]] : copy[a+2]
      if (value1 == 0) {
        a = value2-1
      } else {
        a += 2
      }
    } else if (opcode.opcode == 7) {
      var value1 = (mode1 == 0) ? copy[copy[a+1]] : copy[a+1]
      var value2 = (mode2 == 0) ? copy[copy[a+2]] : copy[a+2]
      copy[copy[a+3]] = (value1 < value2) ? 1 : 0
      a += 3
    } else if (opcode.opcode == 8) {
      var value1 = (mode1 == 0) ? copy[copy[a+1]] : copy[a+1]
      var value2 = (mode2 == 0) ? copy[copy[a+2]] : copy[a+2]
      copy[copy[a+3]] = (value1 == value2) ? 1 : 0
      a += 3
    }
  }

  return { copy, output }
}

console.log(singleLooper(RAW).output)