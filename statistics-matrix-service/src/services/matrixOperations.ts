import { create, all, Matrix, BigNumber } from "mathjs";

// Configuración de Math.js con alta precisión
const math = create(all, { number: "BigNumber", precision: 20 });

export const findMaxValue = (matrix: Matrix): BigNumber => {
  const values = math.flatten(matrix).valueOf() as number[];
  return math.bignumber(math.max(...values.map(num => math.bignumber(num))));
};

export const findMinValue = (matrix: Matrix): BigNumber => {
  const values = math.flatten(matrix).valueOf() as number[];
  return math.bignumber(math.min(...values.map(num => math.bignumber(num))));
};

export const calculateGlobalAverage = (matrix1: Matrix, matrix2: Matrix): BigNumber => {
  const sum1 = calculateTotalSum(matrix1);
  const sum2 = calculateTotalSum(matrix2);

  const totalElements = math.bignumber(
    matrix1.size().reduce((acc, val) => acc * val, 1) +
    matrix2.size().reduce((acc, val) => acc * val, 1)
  );

  return math.divide(math.add(sum1, sum2), totalElements) as BigNumber;
};

export const calculateTotalSum = (matrix: Matrix): BigNumber => {
  const values = math.flatten(matrix).valueOf() as number[];

  return values.reduce(
    (acc, num) => math.add(acc, math.bignumber(num)) as BigNumber,
    math.bignumber(0)
  );
};

export const isDiagonalMatrix = (matrix: Matrix): boolean => {
  const arr = matrix.valueOf() as number[][];
  const size = arr.length;
  for (let i = 0; i < size; i++) {
    for (let j = 0; j < size; j++) {
      if (i !== j && arr[i][j] !== 0) {
        return false;
      }
    }
  }
  return true;
};