const data = Deno.readTextFileSync('./data.txt');

const input = data.split('\n');

const res1 = input.reduce((accumulator, currentValue) => {
  const [command, commandValue] = currentValue.split(' ');

  if (command === 'forward') {
    accumulator.x += Number(commandValue);
  }

  if (command === 'up') {
    accumulator.y -= Number(commandValue);
  }

  if (command === 'down') {
    accumulator.y += Number(commandValue);
  }

  return accumulator;
}, {
  x: 0,
  y: 0,
} as { x: number, y: number });

const product1 = res1.x * res1.y;
const part1 = product1;

console.log('part1', part1);

const res2 = input.reduce((accumulator, currentValue) => {
  const [command, commandValue] = currentValue.split(' ');

  if (command === 'forward') {
    accumulator.x += Number(commandValue);

    // It increases your depth by your aim multiplied by X
    if (accumulator.aim) {
      accumulator.y += accumulator.aim * Number(commandValue);
    }
  }

  if (command === 'up') {
    accumulator.aim -= Number(commandValue);
  }

  if (command === 'down') {
    accumulator.aim += Number(commandValue);
  }

  return accumulator;
}, {
  x: 0,
  y: 0,
  aim: 0,
} as { x: number, y: number, aim: number });

const product2 = res2.x * res2.y;
const part2 = product2;

console.log('part2', part2);
