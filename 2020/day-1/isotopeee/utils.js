const { once } = require("events");
const { createReadStream } = require("fs");
const { createInterface } = require("readline");

async function processInputData({ inputFilePath = "./input.txt" } = {}) {
  const readStream = createReadStream(inputFilePath);

  const readInterface = createInterface({
    input: readStream,
    crlfDelay: Infinity,
  });

  const data = [];

  readInterface.on("line", (line) => {
    data.push(line);
  });

  await once(readInterface, "close");

  return data;
}

function parseToNumberMapPredicate(item) {
  return +item;
}

module.exports = {
  parseToNumberMapPredicate,
  processInputData,
};
