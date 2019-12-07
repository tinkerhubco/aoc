const looper = () => {
  var results = []
  for (var a = 137683; a < 596253; a++) {
    var pass = a.toString()
    if (pass.length === 6 && 
        hasTwoAdjacent(pass) && 
        isIncreasing(pass)
    ) {
      console.log(pass)
      results.push(pass)
    }
  }
  return results
}

const hasAdjacent = (pass) => {
  var adjacent = false
  pass = pass.split('')
  pass.forEach((digit, index) => {
    var next = pass[index+1]
    if (index !== pass.length-1)
      if (digit === next) adjacent = true
  });
  return adjacent
}

const hasTwoAdjacent = (pass) => {
  return (pass.includes('11') && !pass.includes('111')) ||
      (pass.includes('22') && !pass.includes('222')) ||
      (pass.includes('33') && !pass.includes('333')) ||
      (pass.includes('44') && !pass.includes('444')) ||
      (pass.includes('55') && !pass.includes('555')) ||
      (pass.includes('66') && !pass.includes('666')) ||
      (pass.includes('77') && !pass.includes('777')) ||
      (pass.includes('88') && !pass.includes('888')) ||
      (pass.includes('99') && !pass.includes('999'))
}

const isIncreasing = (pass) => {
  var sorted = pass.split('').map((digit) => parseInt(digit)).sort().join('')
  return pass === sorted
}

console.log(looper().length)