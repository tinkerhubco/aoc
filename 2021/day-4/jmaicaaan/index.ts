const data = Deno.readTextFileSync('./data.txt');

const input = data.split('\n');

const drawNumbers = input[0].split(',');
const DRAW_NUMBERS_INDEX = 0;
// If it reaches 5 rows or 5 columns then that board wins!
const TOTAL_COUNT_OF_WINNER = 5;

type BoardMapItem = {
  number: number,
  row: number,
  column: number,
  status: 'open' | 'closed', // open = not marked, closed = marked
};
type BoardMap = Record<string, BoardMapItem>;

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

type GameStatus = {
  shouldStopDrawingNumbers: boolean;
  lastDrawNumber: number;
  winnerBoard: BoardMap | undefined;
  winnerBoardIndex: number;
  winnerDirection: string;
};

const pipe = (...fns: Function[]) => (i: unknown) =>
  fns.reduce((f, d) => d(f), i);

const getAllBoards = () => {
  const copyInput = [...input]
    .filter((line, index) => index !== DRAW_NUMBERS_INDEX && line);
  const boards = [];

  while (copyInput.length) {
    const board = copyInput.splice(0, 5);
    boards.push(board);
  }

  return boards;
};

const buildBoardMap = (board: string[]) => {
  return board.reduce((accumulator, currentValue, currentIndex) => {
    const regex = new RegExp(/\D+/)
    const lineNumbers = currentValue.trim().split(regex);
    const columnIndex = currentIndex;

    lineNumbers.forEach((lineNumber, lineNumberIndex) => {
      const rowIndex = lineNumberIndex;
      accumulator[lineNumber] = {
        number: Number(lineNumber),
        row: rowIndex,
        column: columnIndex,
        status: 'open',
      };
    });

    return accumulator;
  }, {} as BoardMap)
};

const markNumber = (num: number) => (boardMap: BoardMap) => {
  const isNumberNotExistingInBoard = !boardMap[num];

  if (isNumberNotExistingInBoard) {
    return boardMap;
  }

  // copy and mutate
  const newBoardMap = { ...boardMap };
  newBoardMap[num].status = 'closed';
  return newBoardMap;
};

const checkWinner = (boardMap: BoardMap) => {
  const getClosedNumbers = ([, value]: [string, BoardMapItem]) => {
    return value.status === 'closed';
  };

  const groupByCheckResult = (
    accumulator: CheckResult,
    currentValue: [string, BoardMapItem]
  ) => {
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
  };

  return Object.entries(boardMap)
    .filter(getClosedNumbers)
    .reduce(groupByCheckResult, {
      'rows': {},
      'columns': {},
    } as CheckResult)
};

const hasWinner = (checkResult: CheckResult) => {
  const hasTotalCount = (count: number) => count === TOTAL_COUNT_OF_WINNER;

  const rowWinner = Object.values(checkResult.rows)
    .find((value) => hasTotalCount(value.count));

  const columnWinner = Object.values(checkResult.columns)
    .find((value) => hasTotalCount(value.count));

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

  const gameStatus: GameStatus = {
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
      const result = pipe(
        markNumber(Number(drawNumber)),
        checkWinner,
        hasWinner,
      )(board) as { direction: string };

      if (result) {
        gameStatus.shouldStopDrawingNumbers = true;
        gameStatus.lastDrawNumber = Number(drawNumber);
        gameStatus.winnerDirection = result.direction;
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
    // 4662
  }
};

playPart1();

const playPart2 = () => {
  const boards = getAllBoards();

  const gameStatus: GameStatus = {
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

      const result = pipe(
        markNumber(Number(drawNumber)),
        checkWinner,
        hasWinner,
      )(board) as { direction: string };

      // change the game status when all board maps have been evaluated
      if (indexes.length === boardMaps.length) {
        gameStatus.shouldStopDrawingNumbers = true;
        gameStatus.lastDrawNumber = Number(drawNumber);
        gameStatus.winnerDirection = result.direction;
        // do not use the `board` but use the last index then get the appropriate board
        gameStatus.winnerBoard = boardMaps[Number([...indexes].pop())];
        gameStatus.winnerBoardIndex = Number([...indexes].pop());

        return;
      }

      if (result) {
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