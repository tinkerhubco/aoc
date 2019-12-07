const fs = require('fs')
var collection = fs.existsSync('./data.json') ? require('./data.json') : {}

const list = `` 

const generator = (year, day, data) => {
  collection[year] = {}
  collection[year][`day${day}`] = data.split('\n')
  fs.writeFileSync('data.json', JSON.stringify(collection, null, 2))
}

generator(2019, 1, list)