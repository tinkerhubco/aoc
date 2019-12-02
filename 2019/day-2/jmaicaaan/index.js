const inputs = [
    1,0,0,3,
    1,1,2,3,
    1,3,4,3,
    1,5,0,3,
    2,1,10,19,
    1,6,19,23,
    2,23,6,27,
    2,6,27,31,
    2,13,31,35,
    1,10,35,39,
    2,39,13,43,
    1,43,13,47,
    1,6,47,51,
    1,10,51,55,
    2,55,6,59,
    1,5,59,63,
    2,9,63,67,
    1,6,67,71,
    2,9,71,75,
    1,6,75,79,
    2,79,13,83,
    1,83,10,87,
    1,13,87,91,
    1,91,10,95,
    2,9,95,99,
    1,5,99,103,
    2,10,103,107,
    1,107,2,111,
    1,111,5,0,
    99,
    2,14,0,0
  ];
  
  const OPERATION = {
    1: 'add',
    2: 'multiply',
    99: 'halt',
  };
  
  const computeOperation = operationCode => (input1, input2) => {
    if (operationCode === 1) {
      return input1 + input2;
    } else if (operationCode === 2) {
      return input1 * input2;
    }
    return 0;
  };
  
  const generateNumArray = num => nonInclusive => 
    Array.from(Array(num).keys()).map(count => nonInclusive ? 0 : count + 1);
  
  const inspect = (intCodes = [], position = 0) => {
    const metadata = {
      totalHopSteps: 4,
      lastOperation: '', 
      outputPosition: 0,
      output: 0,
    };
    if (!intCodes[position]) {
      return Object.assign(metadata, {
        lastOperation: OPERATION[99],
      });
    }
  
    const opCode = intCodes[position];
    const inputPosition1 = intCodes[position + 1];
    const inputPosition2 = intCodes[position + 2];
    const outputPosition = intCodes[position + 3];
    
    const input1 = intCodes[inputPosition1];
    const input2 = intCodes[inputPosition2];
  
    return Object.assign(metadata, {
      lastOperation: OPERATION[opCode],
      outputPosition,
      output: computeOperation(opCode)(input1, input2)
    });
  };
  
  const inspectIntCodes = (intCodes = [], position = 0) => {
    const modifiedIntCodes = [...intCodes];
    const metadata = inspect(modifiedIntCodes, position);
  
    if (metadata.lastOperation !== OPERATION[99]) {
      const nextPosition = position + metadata.totalHopSteps;
      modifiedIntCodes[metadata.outputPosition] = metadata.output;
      return inspectIntCodes(modifiedIntCodes, nextPosition);
    }
    return modifiedIntCodes;
  };
  
  const getResult1 = (intCodes = []) => {
    intCodes[1] = 12;
    intCodes[2] = 2;
    return inspectIntCodes(intCodes)[0];
  };
  
  const getResult2 = (intCodes = []) => {
    const puzzleOutput = 19690720;
    const range = generateNumArray(99)();
  
    const check = (noun = 1, verb = 1) => {
      intCodes[1] = noun;
      intCodes[2] = verb;
      return inspectIntCodes(intCodes)[0];
    };
  
    return range.reduce((acc, a) => {
      range.forEach((b) => {
        const result = check(a, b);
        if (result === puzzleOutput) {
          acc = 100 * a + b;
        }
      });
      return acc;
    }, 0);
  };
  
  const result1 = getResult1(inputs);
  const result2 = getResult2(inputs);
  
  console.log('result1', result1);
  console.log('result2', result2);