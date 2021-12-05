const data = Deno.readTextFileSync('./data.txt');

/**
 * If the lines overlap with the same amount of threshold
 * it is considered a dangerous area
 * horizontal and vertical only
 */
const countDangerousAreas = (threshold = 2) => (map: Map<string, { count: number }>) => {
  return [...map].flatMap((x) => x[1])
    .filter((x) => x.count >= threshold)
    .length;
};

const input = data.split('\n')
  .map((line) => {
    const [part1, part2] = line
      .split('->')
      .map((x) => x.trim());

    const [x1, y1] = part1
      .split(',')
      .map(Number);
    const [x2, y2] = part2
      .split(',')
      .map(Number);

    return {
      x1,
      x2,
      y1,
      y2,
      isIgnored: false,
    };
  });


const playPart1 = () => {
  const allPoints = input
    .map((point) => {
      const isIgnored = point.x1 !== point.x2 && point.y1 !== point.y2;
      // part 1
      // For now, only consider horizontal and vertical lines: lines where either x1 = x2 or y1 = y2.
      return {
        ...point,
        isIgnored,
      };
    })
    .filter((x) => !x.isIgnored);

  // contains all the points
  const map = new Map();

  const isExistingPoint = (key: string) => {
    const point = map.has(key);
    if (!point) {
      return false;
    }
    return true;
  };

  // TODO - Fix typings
  const markPoint = (line: typeof input[0]) => {
    const options = {
      start: 0,
      end: 0,
      commonPoint: 0,
      commonPointAxis: '',
    };

    if (line.y1 === line.y2) {
      options.commonPoint = line.y1;
      options.commonPointAxis = 'y';
      if (line.x1 < line.x2) {
        // asc
        options.start = line.x1;
        options.end = line.x2;
      } else {
        options.start = line.x2;
        options.end = line.x1;
      }
    }

    if (line.x1 === line.x2) {
      options.commonPoint = line.x1;
      options.commonPointAxis = 'x';
      if (line.y1 < line.y2) {
        // asc
        options.start = line.y1;
        options.end = line.y2;
      } else {
        // desc
        options.start = line.y2;
        options.end = line.y1;
      }
    }

    for (let i = options.start; i <= options.end; i++) {
      const key = options.commonPointAxis === 'x'
        ? `${options.commonPoint},${i}`
        : `${i},${options.commonPoint}`;

      if (isExistingPoint(key)) {
        const point = map.get(key);
        // increment if existing point
        map.set(key, {
          count: point.count += 1
        })
      } else {
        map.set(key, {
          count: 1,
        });
      }
    }
  };

  allPoints.forEach(markPoint);
  const res1 = countDangerousAreas()(map);
  const part1 = res1;

  console.log('part1', part1);
};

playPart1();

const playPart2 = () => {
  const allPoints = input;

  // contains all the points
  const map = new Map();

  const isExistingPoint = (key: string) => {
    const point = map.has(key);
    if (!point) {
      return false;
    }
    return true;
  };

  // TODO - Fix typings
  const markPoint = (line: typeof input[0]) => {
    const newLine = {
      ...line,
    };

    const initialKey = `${newLine.x1},${newLine.y1}`;
    if (isExistingPoint(initialKey)) {
      const point = map.get(initialKey);
      map.set(initialKey, {
        count: point.count + 1,
      });
    } else {
      map.set(`${newLine.x1},${newLine.y1}`, {
        count: 1
      });
    }

    while (
      newLine.x1 !== newLine.x2
      || newLine.y1 !== newLine.y2
    ) {
      if (newLine.x1 > newLine.x2) {
        newLine.x1 -= 1;
      }
      if (newLine.x1 < newLine.x2) {
        newLine.x1 += 1;
      }

      if (newLine.y1 > newLine.y2) {
        newLine.y1 -= 1;
      }
      if (newLine.y1 < newLine.y2) {
        newLine.y1 += 1;
      }

      const key = `${newLine.x1},${newLine.y1}`;

      if (isExistingPoint(key)) {
        const point = map.get(key);
        map.set(key, {
          count: point.count + 1,
        });
      } else {
        map.set(key, {
          count: 1,
        });
      }
    }
  };

  allPoints.forEach(markPoint);
  const res2 = countDangerousAreas()(map);
  const part2 = res2;

  console.log('part2', part2);
  // console.log('map', map);
};

playPart2();