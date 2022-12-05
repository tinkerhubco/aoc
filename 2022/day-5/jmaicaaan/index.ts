const input = `
move 3 from 9 to 6
move 7 from 6 to 2
move 1 from 1 to 5
move 7 from 7 to 1
move 3 from 9 to 7
move 1 from 9 to 1
move 1 from 7 to 2
move 11 from 1 to 8
move 9 from 8 to 2
move 1 from 6 to 7
move 3 from 7 to 3
move 7 from 3 to 4
move 9 from 8 to 7
move 3 from 3 to 1
move 2 from 5 to 2
move 6 from 7 to 3
move 1 from 1 to 7
move 1 from 9 to 2
move 1 from 5 to 3
move 1 from 8 to 2
move 2 from 7 to 5
move 1 from 1 to 4
move 3 from 5 to 8
move 2 from 8 to 7
move 1 from 8 to 9
move 7 from 3 to 1
move 8 from 2 to 5
move 3 from 7 to 3
move 1 from 5 to 1
move 1 from 9 to 6
move 1 from 7 to 4
move 1 from 6 to 3
move 1 from 7 to 1
move 9 from 4 to 5
move 8 from 1 to 2
move 3 from 3 to 2
move 1 from 1 to 6
move 7 from 5 to 6
move 1 from 1 to 5
move 1 from 3 to 5
move 21 from 2 to 3
move 8 from 6 to 3
move 5 from 4 to 9
move 9 from 3 to 8
move 17 from 3 to 5
move 6 from 2 to 1
move 2 from 9 to 1
move 3 from 3 to 6
move 3 from 2 to 5
move 7 from 8 to 2
move 3 from 6 to 9
move 2 from 2 to 4
move 1 from 2 to 6
move 2 from 2 to 6
move 2 from 6 to 5
move 1 from 6 to 1
move 2 from 2 to 7
move 1 from 8 to 2
move 4 from 9 to 1
move 4 from 1 to 6
move 1 from 8 to 5
move 3 from 6 to 9
move 1 from 9 to 1
move 2 from 9 to 2
move 4 from 4 to 5
move 1 from 7 to 8
move 1 from 7 to 5
move 8 from 1 to 8
move 1 from 1 to 9
move 1 from 6 to 8
move 2 from 2 to 6
move 1 from 1 to 3
move 1 from 2 to 5
move 1 from 3 to 4
move 3 from 9 to 4
move 4 from 4 to 1
move 29 from 5 to 1
move 2 from 6 to 3
move 2 from 3 to 5
move 2 from 5 to 9
move 7 from 8 to 1
move 3 from 8 to 6
move 6 from 1 to 6
move 2 from 9 to 8
move 2 from 5 to 3
move 3 from 5 to 6
move 2 from 5 to 6
move 1 from 5 to 1
move 2 from 3 to 9
move 1 from 8 to 6
move 1 from 8 to 3
move 1 from 3 to 5
move 5 from 1 to 5
move 5 from 6 to 2
move 25 from 1 to 9
move 9 from 9 to 3
move 7 from 6 to 8
move 9 from 5 to 9
move 2 from 6 to 5
move 6 from 9 to 7
move 1 from 6 to 8
move 3 from 2 to 1
move 3 from 8 to 1
move 5 from 9 to 6
move 3 from 9 to 1
move 4 from 6 to 9
move 2 from 7 to 4
move 1 from 4 to 1
move 1 from 6 to 2
move 7 from 1 to 6
move 1 from 9 to 8
move 9 from 3 to 9
move 5 from 1 to 7
move 1 from 5 to 7
move 3 from 1 to 7
move 3 from 6 to 7
move 8 from 9 to 1
move 3 from 7 to 3
move 1 from 5 to 6
move 3 from 1 to 7
move 4 from 1 to 4
move 2 from 8 to 5
move 1 from 4 to 2
move 3 from 2 to 7
move 2 from 6 to 4
move 1 from 1 to 2
move 18 from 7 to 5
move 1 from 7 to 5
move 1 from 2 to 3
move 4 from 5 to 9
move 1 from 2 to 1
move 2 from 3 to 9
move 2 from 3 to 4
move 2 from 6 to 5
move 1 from 8 to 3
move 4 from 9 to 7
move 1 from 1 to 9
move 3 from 5 to 2
move 2 from 8 to 6
move 2 from 6 to 1
move 5 from 5 to 7
move 7 from 9 to 7
move 11 from 5 to 9
move 3 from 7 to 6
move 6 from 4 to 9
move 5 from 7 to 3
move 6 from 3 to 6
move 2 from 1 to 2
move 2 from 4 to 9
move 6 from 9 to 2
move 1 from 7 to 5
move 10 from 2 to 9
move 4 from 9 to 4
move 1 from 4 to 3
move 31 from 9 to 3
move 1 from 9 to 4
move 6 from 3 to 8
move 1 from 5 to 8
move 5 from 6 to 4
move 4 from 3 to 2
move 1 from 4 to 6
move 22 from 3 to 7
move 6 from 4 to 7
move 4 from 6 to 2
move 8 from 8 to 1
move 3 from 2 to 8
move 2 from 1 to 9
move 1 from 2 to 6
move 3 from 2 to 5
move 2 from 5 to 4
move 2 from 6 to 4
move 24 from 7 to 4
move 1 from 7 to 4
move 2 from 1 to 5
move 2 from 9 to 6
move 10 from 4 to 6
move 3 from 1 to 6
move 6 from 7 to 1
move 2 from 2 to 3
move 1 from 7 to 4
move 2 from 8 to 4
move 1 from 8 to 5
move 4 from 5 to 2
move 5 from 4 to 1
move 2 from 7 to 8
move 2 from 8 to 4
move 5 from 6 to 3
move 2 from 4 to 3
move 1 from 7 to 5
move 2 from 3 to 6
move 1 from 5 to 1
move 3 from 6 to 8
move 11 from 4 to 3
move 7 from 6 to 1
move 3 from 8 to 1
move 1 from 2 to 3
move 2 from 6 to 9
move 2 from 2 to 3
move 3 from 4 to 3
move 2 from 9 to 4
move 1 from 6 to 3
move 5 from 1 to 2
move 2 from 4 to 3
move 24 from 3 to 7
move 3 from 3 to 9
move 1 from 2 to 6
move 1 from 2 to 5
move 1 from 6 to 1
move 4 from 2 to 1
move 2 from 9 to 2
move 1 from 2 to 4
move 18 from 7 to 5
move 1 from 2 to 1
move 1 from 9 to 1
move 2 from 5 to 7
move 13 from 1 to 8
move 3 from 4 to 9
move 7 from 1 to 7
move 13 from 7 to 6
move 1 from 9 to 5
move 3 from 4 to 3
move 1 from 9 to 8
move 3 from 1 to 3
move 1 from 9 to 5
move 2 from 1 to 4
move 2 from 7 to 3
move 4 from 3 to 1
move 1 from 1 to 5
move 9 from 6 to 7
move 5 from 7 to 1
move 2 from 4 to 1
move 4 from 6 to 1
move 3 from 5 to 3
move 3 from 3 to 5
move 7 from 1 to 6
move 6 from 6 to 1
move 1 from 6 to 8
move 2 from 7 to 9
move 2 from 1 to 5
move 1 from 3 to 7
move 7 from 5 to 9
move 10 from 1 to 5
move 8 from 8 to 4
move 6 from 4 to 8
move 1 from 4 to 1
move 2 from 9 to 8
move 2 from 1 to 3
move 2 from 7 to 3
move 1 from 7 to 8
move 4 from 3 to 8
move 1 from 3 to 2
move 20 from 5 to 8
move 1 from 2 to 4
move 4 from 9 to 4
move 4 from 4 to 5
move 18 from 8 to 6
move 3 from 9 to 6
move 1 from 3 to 9
move 10 from 8 to 7
move 7 from 7 to 9
move 7 from 8 to 5
move 3 from 7 to 8
move 6 from 5 to 1
move 6 from 9 to 4
move 1 from 9 to 6
move 1 from 3 to 6
move 1 from 8 to 5
move 1 from 9 to 4
move 12 from 6 to 7
move 5 from 7 to 1
move 6 from 8 to 5
move 1 from 5 to 1
move 3 from 5 to 3
move 8 from 4 to 9
move 2 from 3 to 7
move 4 from 7 to 2
move 10 from 5 to 6
move 11 from 1 to 6
move 4 from 2 to 5
move 1 from 3 to 8
move 1 from 8 to 9
move 1 from 4 to 7
move 3 from 7 to 4
move 1 from 1 to 6
move 1 from 4 to 7
move 1 from 7 to 1
move 4 from 5 to 2
move 3 from 7 to 1
move 2 from 4 to 8
move 20 from 6 to 8
move 4 from 1 to 5
move 2 from 5 to 2
move 6 from 6 to 1
move 5 from 1 to 8
move 7 from 6 to 2
move 6 from 9 to 7
move 2 from 9 to 8
move 2 from 7 to 4
move 4 from 2 to 6
move 3 from 5 to 8
move 12 from 8 to 7
move 1 from 4 to 3
move 1 from 2 to 9
move 1 from 9 to 2
move 1 from 6 to 8
move 1 from 3 to 1
move 2 from 1 to 6
move 1 from 4 to 2
move 3 from 6 to 2
move 2 from 5 to 7
move 1 from 9 to 8
move 6 from 2 to 4
move 17 from 7 to 1
move 10 from 1 to 7
move 4 from 2 to 6
move 10 from 7 to 8
move 3 from 6 to 2
move 4 from 4 to 1
move 2 from 6 to 4
move 4 from 2 to 6
move 1 from 7 to 1
move 2 from 4 to 3
move 12 from 1 to 7
move 5 from 6 to 3
move 17 from 8 to 2
move 4 from 3 to 8
move 1 from 4 to 2
move 20 from 8 to 7
move 19 from 2 to 6
move 7 from 6 to 3
move 7 from 3 to 5
move 2 from 5 to 7
move 4 from 6 to 9
move 1 from 4 to 2
move 1 from 2 to 1
move 2 from 3 to 6
move 1 from 2 to 6
move 1 from 3 to 1
move 4 from 6 to 2
move 1 from 5 to 9
move 7 from 7 to 3
move 7 from 3 to 8
move 5 from 8 to 1
move 2 from 8 to 3
move 1 from 2 to 1
move 3 from 5 to 6
move 1 from 3 to 9
move 2 from 9 to 2
move 8 from 1 to 7
move 3 from 7 to 6
move 2 from 2 to 4
move 21 from 7 to 3
move 10 from 3 to 1
move 2 from 9 to 2
move 7 from 3 to 4
move 3 from 3 to 7
move 4 from 2 to 3
move 3 from 7 to 8
move 1 from 3 to 6
move 1 from 3 to 2
move 4 from 7 to 9
move 10 from 1 to 6
move 1 from 5 to 9
move 6 from 7 to 2
move 24 from 6 to 5
move 2 from 8 to 4
move 1 from 8 to 6
move 2 from 2 to 9
move 5 from 2 to 7
move 1 from 2 to 9
move 11 from 4 to 1
move 3 from 3 to 2
move 4 from 9 to 7
move 1 from 1 to 5
move 1 from 6 to 1
move 5 from 1 to 9
move 5 from 9 to 7
move 5 from 7 to 5
move 23 from 5 to 2
move 5 from 7 to 8
move 6 from 5 to 6
move 1 from 3 to 7
move 1 from 5 to 7
move 6 from 7 to 8
move 3 from 6 to 1
move 2 from 8 to 7
move 4 from 2 to 1
move 4 from 8 to 5
move 7 from 2 to 3
move 1 from 7 to 4
move 1 from 4 to 7
move 4 from 3 to 8
move 6 from 1 to 9
move 4 from 8 to 6
move 3 from 1 to 5
move 3 from 8 to 5
move 1 from 1 to 8
move 3 from 9 to 1
move 3 from 6 to 7
move 1 from 7 to 9
move 3 from 8 to 3
move 8 from 5 to 7
move 11 from 2 to 8
move 5 from 8 to 3
move 1 from 8 to 7
move 10 from 3 to 4
move 2 from 5 to 8
move 3 from 9 to 2
move 1 from 9 to 6
move 7 from 2 to 7
move 6 from 9 to 4
move 1 from 8 to 5
move 3 from 6 to 8
move 1 from 5 to 3
move 2 from 3 to 1
move 6 from 1 to 3
move 13 from 7 to 5
move 16 from 4 to 3
move 2 from 1 to 5
move 5 from 5 to 4
move 11 from 3 to 4
move 2 from 7 to 1
move 7 from 3 to 1
move 2 from 8 to 3
move 8 from 1 to 9
move 12 from 4 to 8
move 1 from 1 to 4
move 2 from 6 to 2
move 3 from 7 to 8
move 2 from 4 to 6
move 5 from 8 to 1
move 3 from 7 to 5
move 6 from 5 to 7
move 2 from 2 to 5
move 1 from 4 to 9
move 5 from 1 to 8
move 6 from 3 to 1
move 7 from 5 to 7
move 7 from 9 to 2
move 1 from 6 to 7
move 1 from 1 to 9
move 2 from 5 to 3
move 2 from 9 to 6
move 13 from 7 to 3
move 2 from 6 to 1
move 1 from 9 to 2
move 16 from 8 to 7
move 6 from 8 to 5
move 3 from 2 to 5
move 4 from 2 to 1
move 3 from 1 to 8
move 2 from 8 to 9
move 1 from 8 to 7
move 1 from 2 to 1
move 8 from 3 to 1
move 1 from 4 to 5
move 1 from 6 to 3
move 2 from 9 to 7
move 5 from 1 to 4
move 15 from 7 to 9
move 11 from 9 to 3
move 7 from 1 to 3
move 2 from 1 to 6
move 1 from 6 to 3
move 2 from 4 to 5
move 2 from 4 to 9
move 7 from 5 to 9
move 5 from 9 to 3
move 1 from 1 to 6
move 5 from 5 to 9
move 1 from 4 to 8
move 1 from 8 to 4
move 3 from 7 to 4
move 8 from 9 to 5
move 1 from 6 to 4
move 4 from 9 to 3
move 1 from 9 to 3
move 23 from 3 to 1
move 12 from 1 to 2
move 6 from 1 to 9
move 5 from 9 to 7
move 3 from 3 to 7
move 6 from 4 to 3
move 1 from 6 to 8
move 6 from 1 to 2
move 3 from 7 to 3
move 3 from 2 to 5
move 10 from 3 to 5
move 1 from 1 to 8
move 12 from 2 to 5
move 3 from 2 to 9
move 2 from 8 to 4
move 13 from 5 to 1
move 2 from 9 to 2
move 2 from 1 to 3
move 11 from 3 to 1
move 2 from 2 to 1
move 2 from 1 to 9
move 16 from 1 to 7
move 17 from 5 to 8
move 1 from 1 to 2
move 3 from 9 to 6
`;

const testInput = `
move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`;

// const testInput = `
// move 1 from 2 to 1
// move 3 from 1 to 3
// `;

const crateMap = new Map<number, string[]>();
// crateMap.set(1, ['Z', 'N']);
// crateMap.set(2, ['M', 'C', 'D']);
// crateMap.set(3, ['P']);

const getPart1 = () => {
  crateMap.set(1, ['S', 'C', 'V', 'N']);
  crateMap.set(2, ['Z', 'M', 'J', 'H', 'N', 'S']);
  crateMap.set(3, ['M', 'C', 'T', 'G', 'J', 'N', 'D']);
  crateMap.set(4, ['T', 'D', 'F', 'J', 'W', 'R', 'M']);
  crateMap.set(5, ['P', 'F', 'H']);
  crateMap.set(6, ['C', 'T', 'Z', 'H', 'J']);
  crateMap.set(7, ['D', 'P', 'R', 'Q', 'F', 'S', 'L', 'Z']);
  crateMap.set(8, ['C', 'S', 'L', 'H', 'D', 'F', 'P', 'W']);
  crateMap.set(9, ['D', 'S', 'M', 'P', 'F', 'N', 'G', 'Z']);

  input.trim().split('\n').forEach((procedure) => {
    const itemsToMove = Number(procedure.split('from')[0].split('move')[1].trim());
    const originCrate = Number(procedure.split('from')[1].trim().split(' to ')[0]);
    const destinationCrate = Number(procedure.split('from')[1].trim().split(' to ')[1]);

    for (let i = 1; i <= itemsToMove; i++) {
      const originCrateStacks = [...(crateMap.get(originCrate) || [])];
      const destinationCrateStacks = [...(crateMap.get(destinationCrate) || [])];

      const updatedDestinationCrateStacks = [...destinationCrateStacks, (originCrateStacks.pop() || '')];
      crateMap.set(originCrate, originCrateStacks);
      crateMap.set(destinationCrate, updatedDestinationCrateStacks);
    }
  });

  const topStacksOnAllCrates = [...crateMap].map((crate) => {
    const [, stacks] = crate;
    return [...stacks].pop();
  }).join('');

  return topStacksOnAllCrates;
};

const getPart2 = () => {
  crateMap.set(1, ['S', 'C', 'V', 'N']);
  crateMap.set(2, ['Z', 'M', 'J', 'H', 'N', 'S']);
  crateMap.set(3, ['M', 'C', 'T', 'G', 'J', 'N', 'D']);
  crateMap.set(4, ['T', 'D', 'F', 'J', 'W', 'R', 'M']);
  crateMap.set(5, ['P', 'F', 'H']);
  crateMap.set(6, ['C', 'T', 'Z', 'H', 'J']);
  crateMap.set(7, ['D', 'P', 'R', 'Q', 'F', 'S', 'L', 'Z']);
  crateMap.set(8, ['C', 'S', 'L', 'H', 'D', 'F', 'P', 'W']);
  crateMap.set(9, ['D', 'S', 'M', 'P', 'F', 'N', 'G', 'Z']);

  input.trim().split('\n').forEach((procedure) => {
    const itemsToMove = Number(procedure.split('from')[0].split('move')[1].trim());
    const originCrate = Number(procedure.split('from')[1].trim().split(' to ')[0]);
    const destinationCrate = Number(procedure.split('from')[1].trim().split(' to ')[1]);

    const unsortedPoppedStacks: string[] = [];

    for (let i = 1; i <= itemsToMove; i++) {
      const originCrateStacks = [...(crateMap.get(originCrate) || [])];
      const stack = originCrateStacks.pop() || '';
      
      unsortedPoppedStacks.push(stack);
      crateMap.set(originCrate, originCrateStacks);
    }

    const destinationCrateStacks = [...(crateMap.get(destinationCrate) || [])];
    const updatedDestinationCrateStacks = [...destinationCrateStacks, ...unsortedPoppedStacks.reverse()];
    crateMap.set(destinationCrate, updatedDestinationCrateStacks);
  });

  const topStacksOnAllCrates = [...crateMap].map((crate) => {
    const [, stacks] = crate;
    return [...stacks].pop();
  }).join('');

  return topStacksOnAllCrates;
};

const part1 = getPart1();
const part2 = getPart2();

// part 1 CNSZFDVLJ
// part 2 QNDWLMGNS

console.log('part1', part1);
console.log('part2', part2);

export { };
