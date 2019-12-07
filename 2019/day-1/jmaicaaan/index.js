const inputs = [
    140170,
    75120,
    75645,
    134664,
    124948,
    137630,
    146662,
    116881,
    120030,
    94332,
    50473,
    59361,
    128237,
    84894,
    51368,
    128802,
    57275,
    129235,
    113481,
    66378,
    55842,
    90548,
    107696,
    53603,
    130458,
    80306,
    120820,
    131313,
    100303,
    59224,
    123369,
    140584,
    60642,
    68184,
    103101,
    82278,
    51968,
    51048,
    98139,
    60498,
    127082,
    71197,
    109478,
    71286,
    84840,
    141305,
    51800,
    72352,
    93147,
    73549,
    122739,
    62363,
    58453,
    59000,
    63564,
    63424,
    51053,
    120826,
    123337,
    130824,
    59053,
    77983,
    68977,
    67126,
    96051,
    53024,
    145647,
    139343,
    113236,
    59396,
    146174,
    148622,
    83384,
    86938,
    100673,
    80757,
    107675,
    147417,
    124538,
    136463,
    104609,
    149559,
    136037,
    54997,
    139674,
    101638,
    65739,
    70029,
    143847,
    122035,
    66256,
    78087,
    105045,
    108867,
    99630,
    127173,
    139021,
    139759,
    134171,
    104869,
  ];
  
  const getFuelRequired = (mass = 0) => Math.floor(mass / 3) - 2;
  const getSumFromArray = (arr = []) => arr.reduce((acc, curr) => acc += curr, 0);
  
  const result1 = inputs.reduce((acc, curr) => acc += getFuelRequired(curr), 0);
  
  const rocketDoubleChecker = (initialFuel = 0) => {
    const totalFuels = [initialFuel];
    const negativeFuel = 0;
    const compute = (fuel = initialFuel) => {
      const requiredFuel = getFuelRequired(fuel);
      if (requiredFuel > negativeFuel) {
        totalFuels.push(requiredFuel);
        compute(requiredFuel);
        return totalFuels;
      }
      return fuel;
    };
    return compute;
  };
  
  const result2 = inputs.map(
    mass => getFuelRequired(mass)
  ).map(
    fuel => rocketDoubleChecker(fuel)()
  ).reduce(
    (acc, curr) => acc += getSumFromArray(curr), 0
  );
  
  console.log('result1', result1);
  console.log('result2', result2);