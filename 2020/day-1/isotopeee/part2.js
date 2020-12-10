const { parseToNumberMapPredicate, processInputData } = require("./utils");

(async function () {
  const inputData = (await processInputData()).map(parseToNumberMapPredicate);

  const EXPENSE = 2020;

  function getAnswer(input) {
    const entries = getEntries(input);

    const answer = entries.reduce((acc, current) => (acc = acc * current), 1);

    return answer;
  }

  function getEntries(inputData) {
    let entries;

    for (const firstEntry of inputData) {
      const secondAndThirdEntrySum = EXPENSE - firstEntry;

      const isSecondAndThirdEntrySumFeasible = secondAndThirdEntrySum > 0;
      /**
       * ! This check is not necessary but just to guard if the firstEntry is greater than `2020`
       */
      if (!isSecondAndThirdEntrySumFeasible) {
        continue;
      }

      const valuesLowerThanSecondAndThirdEntrySum = inputData.filter(
        (num) => num < secondAndThirdEntrySum
      );

      const secondAndThirdEntryMedian = secondAndThirdEntrySum / 2;

      const valuesLowerThanMedian = valuesLowerThanSecondAndThirdEntrySum.filter(
        (num) => num < secondAndThirdEntryMedian
      );
      const valuesHigherThanMedian = valuesLowerThanSecondAndThirdEntrySum.filter(
        (num) => num > secondAndThirdEntryMedian
      );

      let secondEntry;
      let thirdEntry;
      for (const possibleSecondEntry of valuesLowerThanMedian) {
        for (const possibleThirdEntry of valuesHigherThanMedian) {
          const sum = possibleSecondEntry + possibleThirdEntry;

          if (sum !== secondAndThirdEntrySum) {
            continue;
          }

          secondEntry = possibleSecondEntry;
          thirdEntry = possibleThirdEntry;
          break;
        }
      }

      if (!secondEntry || !thirdEntry) {
        continue;
      }

      entries = [firstEntry, secondEntry, thirdEntry];

      break;
    }

    return entries;
  }

  const answer = getAnswer(inputData);

  console.log(answer);
})();
