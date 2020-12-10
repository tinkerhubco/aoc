const { parseToNumberMapPredicate, processInputData } = require("./utils");

(async function () {
  const inputData = (await processInputData()).map(parseToNumberMapPredicate);

  const EXPENSE = 2020;

  function getAnswer(input) {
    const entries = getEntries(input);

    const answer = entries.reduce((acc, current) => (acc = acc * current), 1);

    return answer;
  }

  function getEntries(input) {
    let entries;

    for (const firstEntry of input) {
      const secondEntry = EXPENSE - firstEntry;

      const isSecondEntryExist = !!input.find((num) => num === secondEntry);

      if (!isSecondEntryExist) {
        continue;
      }

      entries = [firstEntry, secondEntry];

      break;
    }

    return entries;
  }

  const answer = await getAnswer(inputData);

  console.log(answer);
})();

// 181044
