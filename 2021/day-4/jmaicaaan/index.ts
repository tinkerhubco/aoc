const data = Deno.readTextFileSync('./data.txt');

const input = data.split('\n');

const drawNumbers = input[0].split(',');
const DRAW_NUMBERS_INDEX = 0;

type BoardMapItem = {
  number: number,
  row: number,
  column: number,
  status: 'open' | 'closed', // open = not marked, closed = marked
};
type BoardMap = Record<string, BoardMapItem>;

const pipe = (...fns: Function[]) => (i: unknown) =>
  fns.reduce((f, d) => d(f), i);


const getAllBoards = () => {
  const copyInput = [...input].filter((i, index) => index !== DRAW_NUMBERS_INDEX && i);
  const boards = [];

  while (copyInput.length) {
    const board = copyInput.splice(0, 5);
    boards.push(board);
  }

  return boards;
};

const buildBoardMap = (board: string[]) => {
  const lookup: BoardMap = {};

  board.forEach((boardItem, boardItemIndex) => {
    const regex = new RegExp(/\D+/)
    const itemNumbers = boardItem.trim().split(regex);
    itemNumbers.forEach((itemNumber, itemNumberIndex) => {
      lookup[itemNumber] = {
        number: Number(itemNumber),
        row: boardItemIndex,
        column: itemNumberIndex,
        status: 'open',
      };
    });
  });

  return lookup;
};

const markNumber = (num: number) => (boardMap: BoardMap) => {
  if (boardMap[num]) {
    const newBoardMap = { ...boardMap };
    newBoardMap[num].status = 'closed';
  }

  return boardMap;
};

type CheckResult = {
  rows: {
    [key: string]: {
      count: number;
    };
  },
  columns: {
    [key: string]: {
      count: number;
    };
  },
};
const checkWinner = (boardMap: BoardMap) => {
  const getClosedNumbers = (_key: string, value: BoardMapItem) => {
    return value.status === 'closed';
  };

  const a = Object.entries(boardMap)
    .filter(([key, value]) => getClosedNumbers(key, value))
    .reduce((accumulator, currentValue) => {
      const [, boardItem] = currentValue;

      if (accumulator.rows[boardItem.row]) {
        if (!accumulator.rows[boardItem.row].count) {
          accumulator.rows[boardItem.row].count = 1;
        } else {
          accumulator.rows[boardItem.row].count += 1;
        }
      } else {
        accumulator.rows[boardItem.row] = {
          count: 1,
        };
      }

      if (accumulator.columns[boardItem.column]) {
        if (!accumulator.columns[boardItem.column].count) {
          accumulator.columns[boardItem.column].count = 1;
        } else {
          accumulator.columns[boardItem.column].count += 1;
        }
      } else {
        accumulator.columns[boardItem.column] = {
          count: 1,
        };
      }

      return accumulator;
    }, {
      'rows': {},
      'columns': {},
    } as CheckResult)

  return a;
};

const hasWinner = (checkResult: CheckResult) => {
  const rowWinner = Object.values(checkResult.rows)
    .find((value) => value.count === 5);

  const columnWinner = Object.values(checkResult.columns)
    .find((value) => value.count === 5);

  if (rowWinner) {
    return {
      direction: 'row',
    };
  }

  if (columnWinner) {
    return {
      direction: 'column',
    };
  }

  return;
};

const getAllOpenItems = (boardMap: BoardMap) => {
  return Object.entries(boardMap)
    .filter(([, boardItem]) => boardItem.status === 'open');
};

const getSum = (nums: number[]) => nums.reduce((accumulator, currentValue) => {
  accumulator += currentValue;
  return accumulator;
}, 0);

const playPart1 = () => {
  const boards = getAllBoards();

  const gameStatus: {
    shouldStopDrawingNumbers: boolean;
    lastDrawNumber: number;
    winnerBoard: BoardMap | undefined;
    winnerBoardIndex: number;
    winnerDirection: string;
  } = {
    shouldStopDrawingNumbers: false,
    lastDrawNumber: 0,
    winnerBoard: undefined,
    winnerBoardIndex: 0,
    winnerDirection: '',
  };

  const boardMaps = boards.map(buildBoardMap);

  drawNumbers.forEach((drawNumber) => {
    if (gameStatus.shouldStopDrawingNumbers) {
      return;
    }

    boardMaps.forEach((board, boardIndex) => {
      const res = pipe(
        markNumber(Number(drawNumber)),
        checkWinner,
        hasWinner,
      )(board) as { direction: string };

      if (res) {
        gameStatus.shouldStopDrawingNumbers = true;
        gameStatus.lastDrawNumber = Number(drawNumber);
        gameStatus.winnerDirection = res.direction;
        gameStatus.winnerBoard = board;
        gameStatus.winnerBoardIndex = boardIndex;
        return;
      }
    });
  });

  if (gameStatus.winnerBoard) {
    const allOpenItems = getAllOpenItems(gameStatus.winnerBoard);
    const sumOfAllOpenItems = getSum(
      allOpenItems.map(([, boardItem]) => boardItem.number)
    );
    const res1 = sumOfAllOpenItems * gameStatus.lastDrawNumber;
    const part1 = res1;
    console.log('part1', part1);
  }
};

playPart1();

const playPart2 = () => {
  const boards = getAllBoards();

  const gameStatus: {
    shouldStopDrawingNumbers: boolean;
    lastDrawNumber: number;
    winnerBoard: BoardMap | undefined;
    winnerBoardIndex: number;
    winnerDirection: string;
  } = {
    shouldStopDrawingNumbers: false,
    lastDrawNumber: 0,
    winnerBoard: undefined,
    winnerBoardIndex: 0,
    winnerDirection: '',
  };

  const boardMaps = boards.map(buildBoardMap);
  const indexes: number[] = [];

  drawNumbers.forEach((drawNumber) => {
    boardMaps.forEach((board, boardIndex) => {
      if (gameStatus.shouldStopDrawingNumbers) {
        return;
      }

      const res = pipe(
        markNumber(Number(drawNumber)),
        checkWinner,
        hasWinner,
      )(board) as { direction: string };

      // change the game status when all board maps have been evaluated
      if (indexes.length === boardMaps.length) {
        gameStatus.shouldStopDrawingNumbers = true;
        gameStatus.lastDrawNumber = Number(drawNumber);
        gameStatus.winnerDirection = res.direction;
        // do not use the `board` but use the last index then get the appropriate board
        gameStatus.winnerBoard = boardMaps[Number([...indexes].pop())];
        gameStatus.winnerBoardIndex = Number([...indexes].pop());

        return;
      }

      if (res) {
        // if this board has a winner, push it
        if (!indexes.includes(boardIndex)) {
          indexes.push(boardIndex);
        }
      }
    });
  });

  if (gameStatus.winnerBoard) {
    const allOpenItems = getAllOpenItems(gameStatus.winnerBoard);
    const sumOfAllOpenItems = getSum(
      allOpenItems.map(([, boardItem]) => boardItem.number)
    );
    const res2 = sumOfAllOpenItems * gameStatus.lastDrawNumber;
    const part2 = res2;
    console.log('part2', part2);
    // 12080
  }
};

playPart2();