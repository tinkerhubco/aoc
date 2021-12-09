const data = Deno.readTextFileSync('./data.txt');

const input = data.split(',').map(Number);

const pipe = (...fns: Function[]) => (i: unknown) => fns.reduce((f, d) => d(f), i);

// Median
const findLeastDist = (arr: number[]) => {
  const items = [...arr].sort((a, b) => a - b);
  const numberOfItems = items.length;

  const isEven = numberOfItems % 2 === 0;

  if (isEven) {
    const median = Math.round((numberOfItems - 1) / 2);
    return items[median];
  }

  const median = numberOfItems / 2;
  return items[median];
};

const computeTotalForLeastDistanceOnAllPoints = (arr: number[]) => (leastDistance: number) =>
  arr.map((point) => Math.abs(point - leastDistance));

const getSum = (arr: number[]) => arr.reduce((s, n) => s + n, 0);
const getSumUsingGauss = (n: number) => (n * (n + 1)) / 2;

const playPart1 = () => {
  const res1 = pipe(
    findLeastDist,
    computeTotalForLeastDistanceOnAllPoints(input),
    getSum,
  )(input);

  const part1 = res1;
  console.log('part1', part1);
};

playPart1();

const playPart2 = () => {
  const [max] = input.sort((a, b) => b - a);

  const getCostUsingGaussFromPoint = (point: number) => {
    return input.reduce((accumulator, currentValue) => {
      const distance = Math.abs(currentValue - point);
      return accumulator += getSumUsingGauss(distance);
    }, 0);
  };

  const res2 = new Array(max).fill(undefined)
    .reduce((accumulator, _currentValue, currentIndex) => {
      const cost = getCostUsingGaussFromPoint(currentIndex);
      if (cost < accumulator) {
        return cost;
      }
      return accumulator;
    }, Number.MAX_VALUE);

  const part2 = res2;
  console.log('part2', part2);
};

playPart2();