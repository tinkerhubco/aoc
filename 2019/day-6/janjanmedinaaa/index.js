const getSingles = (data, getFirst = true) => {
  var singles = []
  var reps = data.split('\n')
  for (var a = 0; a < reps.length; a++) {
    var count = 1
    var splitFirst = reps[a].split(')')
    for (var b = 0; b < reps.length; b++) {
      var splitSecond = reps[b].split(')')
      if (a !== b) {
        if (getFirst) {
          if (splitFirst[0] === splitSecond[1]) count += 1
        } else if ((splitFirst[1] === splitSecond[0])) {
          count += 1
        }
      }
    }
    if (count === 1) singles.push(reps[a])
  }
  return singles
}

const total = []
const builder = (current) => {
  var combinations = []
  var raw = RAW_DATA.split('\n')
  for (var a = 0; a < current.length; a++) {
    var splitCurrent = current[a].split(')')
    total.push(current[a])
    for(var b = 0; b < raw.length; b++) {
      var splitItem = raw[b].split(')')
      if (splitCurrent[splitCurrent.length-1] === splitItem[0]) {
        combinations.push(`${current[a]})${splitItem.slice(1, splitItem.length).join(')')}`)
        matched = true
      }
    }
  }

  if (combinations.length === 0) 
    return current 
  else 
    return builder(combinations)
}

const summation = (data) => {
  var count = 0
  data.forEach(i => count += i.split(')').length-1)
  return count
}

const removeFirstCommon = (you, san) => {
  var splitYou = you.split(')')
  var splitSan = san.split(')')

  for(var a = 0; a < splitYou.length; a++) {
    if (splitYou[a] !== splitSan[a]) {
      return [
        splitYou.slice(a-1).join(')'),
        splitSan.slice(a-1).join(')')
      ]
    }
  }
}

const getDistance = (you, san) => {
  var splitYou = you.split(')').length-2
  var splitSan = san.split(')').length-2
  return splitYou + splitSan
}

var build = builder(getSingles(RAW_DATA))[0]
var youSan = total.filter(item => {
  var split = item.split(')')
  return split[split.length-1] === 'YOU' || split[split.length-1] === 'SAN'
})

var removeFirstCommonResult = removeFirstCommon(youSan[1], youSan[0])
console.log(getDistance(removeFirstCommonResult[0], removeFirstCommonResult[1]))