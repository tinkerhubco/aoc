const data = Deno.readTextFileSync('./data.txt');

const input = data.split('\n');

// get the most common bit to produce a new binary = gamma rate binary
// get the least common bit to produce a new binary = epsilon rate binary
// convert the gammar rate binary and epsilon rate binary to decimanl
// multiply the two decimals

const binaryLen = input[0].length;

const combineAllBitsByPosition = (
  accumulator: Array<Array<string>>, currentValue: string
) => {
  new Array(binaryLen)
    .fill(undefined)
    .forEach((_, positionIndex) => {
      const a = currentValue[positionIndex];

      if (!accumulator[positionIndex]) {
        accumulator[positionIndex] = [];
      }

      accumulator[positionIndex].push(a)
      return a;
    });

  return accumulator;
};

const countBit = (bit: string) => (bits: string[]) => {
  return bits.reduce((accumulator, currentValue) => {
    if (currentValue !== bit) {
      return accumulator;
    }

    return accumulator += 1;
  }, 0);
};

const mapBitsToGammaAndEpsilon = (
  accumulator: Array<Array<string>>,
  currentValue: string[]
) => {
  const zeroBitCount = countBit('0')(currentValue);
  const oneBitCount = countBit('1')(currentValue);

  if (zeroBitCount > oneBitCount) {
    accumulator.push(['0', '1']);
    return accumulator;
  }

  accumulator.push(['1', '0']);
  return accumulator;
};

const mapToGammaAndEpsilon = (
  accumulator: string[],
  currentValue: string[],
) => {
  if (!accumulator[0]) {
    accumulator = ['', ''];
  }
  accumulator[0] = accumulator[0] + currentValue[0];
  accumulator[1] = accumulator[1] + currentValue[1];

  return accumulator;
};

const binaryToDecimal = (binary: string) => parseInt(binary, 2);

const getTotalConsumption = (
  accumulator: number,
  currentValue: number
) => {
  if (!accumulator) {
    accumulator = currentValue;
    return accumulator;
  }
  accumulator *= currentValue;
  return accumulator;
};

const filterByBit = (bit: string, position: number) => (source: string[]) => {
  return source
    .reduce((accumulator, currentValue) => {
      if (currentValue[position] !== bit) {
        return accumulator;
      }

      accumulator.push(currentValue);
      return accumulator;
    }, [] as Array<string>)
};


const res1 = input
  .reduce(combineAllBitsByPosition, [] as Array<Array<string>>)
  .reduce(mapBitsToGammaAndEpsilon, [] as Array<Array<string>>)
  .reduce(mapToGammaAndEpsilon, [] as Array<string>)
  .map(binaryToDecimal) // convert to a reduce and combine with `getTotalConsumption`?
  .reduce(getTotalConsumption, 0);

const part1 = res1;
// 4006064
console.log('part1', part1);

/**
 * part 2
 */

const getOxygenGeneratorRating = () => {
  const newBitsSource: string[] = [];

  const getBit = (source: string[]) => (position: number) => {
    return source
      .reduce(combineAllBitsByPosition, [] as Array<Array<string>>)
      .reduce((accumulator, _currentValue, _, arr) => {
        if (arr[position]) {
          const zeroBitCount = countBit('0')(arr[position]);
          const oneBitCount = countBit('1')(arr[position]);
          /**
           * For retrieving oxygen generator rating only
           */
          if (zeroBitCount === oneBitCount) {
            return '1';
          }
          /**
           * Note on the operator used here
           * Oxygen generator rating formula will use 
           * `0` if zero bits is higher
           * `1` if zero bits is lower
           */
          return zeroBitCount > oneBitCount ? '0' : '1';
        }

        return accumulator;
      }, '');
  };

  let inputSource = input;
  let outputSource = 0;
  new Array(binaryLen)
    .fill(undefined)
    .forEach((_, positionIndex) => {
      const bit = getBit(inputSource)(positionIndex);
      newBitsSource.push(bit);
      const newFilteredInput = filterByBit(
        newBitsSource[positionIndex],
        positionIndex,
      )(inputSource);

      inputSource = newFilteredInput;
      if (inputSource.length === 1 && inputSource[0]) {
        outputSource = binaryToDecimal(inputSource[0]);
        return;
      }
    });

  return outputSource;
};

const getC02ScrubberRating = () => {
  const newBitsSource: string[] = [];

  const getBit = (source: string[]) => (position: number) => {
    return source
      .reduce(combineAllBitsByPosition, [] as Array<Array<string>>)
      .reduce((accumulator, _currentValue, _, arr) => {
        if (arr[position]) {
          const zeroBitCount = countBit('0')(arr[position]);
          const oneBitCount = countBit('1')(arr[position]);
          /**
           * For retrieving c02 scrubber rating only
           */
          if (zeroBitCount === oneBitCount) {
            return '0';
          }
          /**
           * Note on the operator used here
           * C02 scrubber rating formula will use 
           * `0` if zero bits is lower
           * `1` if zero bits is higher
           */
          return zeroBitCount < oneBitCount ? '0' : '1';
        }

        return accumulator;
      }, '');
  };

  let inputSource = input;
  let outputSource = 0;
  new Array(binaryLen)
    .fill(undefined)
    .forEach((_, positionIndex) => {
      const bit = getBit(inputSource)(positionIndex);
      newBitsSource.push(bit);
      const newFilteredInput = filterByBit(
        newBitsSource[positionIndex],
        positionIndex,
      )(inputSource);

      inputSource = newFilteredInput;
      if (inputSource.length === 1 && inputSource[0]) {
        outputSource = binaryToDecimal(inputSource[0]);
        return;
      }
    });

  return outputSource;
};

const oxygenRating = getOxygenGeneratorRating();
const c02ScrubberRating = getC02ScrubberRating();
const res2 = oxygenRating * c02ScrubberRating
const part2 = res2;
console.log('part2', part2);
// 5941884