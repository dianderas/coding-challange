import { create, all, Matrix } from "mathjs";
import {
  findMaxValue,
  findMinValue,
  calculateGlobalAverage,
  calculateTotalSum,
  isDiagonalMatrix
} from "./matrixOperations";

const math = create(all, { number: "BigNumber", precision: 20 });

describe("Matrix Operations", () => {
  let matrix1: Matrix, matrix2: Matrix, diagonalMatrix: Matrix, nonDiagonalMatrix: Matrix;

  beforeEach(() => {
    matrix1 = math.matrix([
      [-0.16903085094570325, 0.8970852271450604],
      [-0.50709255283711, 0.2760262237369414],
      [-0.8451542547285166, -0.34503277967117696]
    ]);

    matrix2 = math.matrix([
      [-5.916079783099615, -7.437357441610946],
      [0, 0.8280786712108259]
    ]);

    diagonalMatrix = math.matrix([
      [1, 0, 0],
      [0, 2, 0],
      [0, 0, 3]
    ]);

    nonDiagonalMatrix = math.matrix([
      [1, 2, 0],
      [0, 2, 0],
      [0, 0, 3]
    ]);
  });

  test("findMaxValue should return the maximum value in a matrix", () => {
    expect(findMaxValue(matrix1).toNumber()).toBeCloseTo(0.8970852271450604);
    expect(findMaxValue(matrix2).toNumber()).toBeCloseTo(0.8280786712108259);
  });

  test("findMinValue should return the minimum value in a matrix", () => {
    expect(findMinValue(matrix1).toNumber()).toBeCloseTo(-0.8451542547285166);
    expect(findMinValue(matrix2).toNumber()).toBeCloseTo(-7.437357441610946);
  });

  test("calculateTotalSum should return the sum of all values in a matrix", () => {
    expect(calculateTotalSum(matrix1).toNumber()).toBeCloseTo(-0.6931989873005052);
    expect(calculateTotalSum(matrix2).toNumber()).toBeCloseTo(-12.525358553499735);
  });

  test("calculateGlobalAverage should return the correct average for two matrices", () => {
    const result = calculateGlobalAverage(matrix1, matrix2).toNumber();
    expect(result).toBeCloseTo(-1.321855754080024);
  });

  test("isDiagonalMatrix should correctly identify a diagonal matrix", () => {
    expect(isDiagonalMatrix(diagonalMatrix)).toBe(true);
    expect(isDiagonalMatrix(nonDiagonalMatrix)).toBe(false);
  });
});
