const data = Deno.readTextFileSync('./data.txt');

const input = data.split('\n');

const res1 = input.reduce((accumulator, currentValue, index, arr) => {
  if (index === 0) {
    return accumulator;
  }

  const a = Number(currentValue);
  const b = Number(arr.at(index + 1));

  if (a < b) {
    accumulator.increased.push(b);
  }

  if (a > b) {
    accumulator.decreased.push(b);
  }

  return accumulator;
}, {
  increased: [],
  decreased: [],
} as { increased: number[], decreased: number[] });

const initialMeasurementCount = 1;
const part1 = res1.increased.length + initialMeasurementCount;
console.log('part1', part1);

const getSumFromArr = (arr: number[]) => arr.reduce((accumulator, currentValue) => accumulator += currentValue, 0);

const res2 = input.reduce((accumulator, currentValue, index, arr) => {
  const firstSlide3 = [Number(currentValue), Number(arr.at(index + 1)), Number(arr.at(index + 2))];
  const secondSlide3 = [Number(arr.at(index + 1)), Number(arr.at(index + 2)), Number(arr.at(index + 3))];

  if (!firstSlide3.every(Boolean) || !secondSlide3.every(Boolean)) {
    return accumulator;
  }

  const a = getSumFromArr(firstSlide3);
  const b = getSumFromArr(secondSlide3);

  if (a < b) {
    accumulator.increased.push(b);
  }

  if (a > b) {
    accumulator.decreased.push(b);
  }

  return accumulator;
}, {
  increased: [],
  decreased: [],
} as { increased: number[], decreased: number[] });

const part2 = res2.increased.length;
console.log('part2', part2);