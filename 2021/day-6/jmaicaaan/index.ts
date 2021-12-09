const data = Deno.readTextFileSync('./data.txt');

const input = data.split(',')
  .map(Number);

const generateDays = (numberOfDays: number) => {
  const days = new Array(numberOfDays)
    .fill(undefined)
    .map((_, index) => index);
  return days;
};


const playPart1 = () => {
  const days = generateDays(80);

  const res = days.reduce((accumulator) => {
    accumulator = accumulator.map((fishTimer) => {
      if (fishTimer === 0) {
        return [6, 8];
      }
      const newFishTimer = fishTimer - 1;
      return newFishTimer;
    }).flat();

    return accumulator;
  }, [...input]);

  const part1 = res.length;
  console.log('part1', part1);
};

playPart1();

const playPart2 = () => {

  const lifecycleMap = new Map();
  const lifeSpanDays = generateDays(256);

  const initializeLifecycleMap = (initialPopulations = input) => {
    initialPopulations.forEach((population) => {
      const existingEntry = lifecycleMap.get(population);

      if (!existingEntry) {
        lifecycleMap.set(population, 1)
        return;
      }
      lifecycleMap.set(population, existingEntry + 1);
    });
  };

  const update = (map: Map<number, number>) => (key: number, value: number) => {
    map.set(key, (map.get(key) || 0) + value);
  };

  const startLife = (spanDays = lifeSpanDays) => {
    return spanDays.reduce((accumulator) => {
      const copy = new Map();
      accumulator.forEach((count, key) => {
        const updateCopy = update(copy);

        if (key === 0) {
          updateCopy(6, count);
          updateCopy(8, count);
        } else {
          updateCopy(key - 1, count);
        }
      });

      return copy;
    }, lifecycleMap);
  };

  const countLife = (map: Map<number, number>) => {
    return [...map].reduce((accumulator, [, value]) => {
      return accumulator + value;
    }, 0);
  };

  initializeLifecycleMap();
  const newLifecycleMap = startLife();
  const res2 = countLife(newLifecycleMap);
  const part2 = res2;
  console.log('part2', part2);
};

playPart2();
