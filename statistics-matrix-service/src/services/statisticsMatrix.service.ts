import { all, create, Matrix } from "mathjs";
import { calculateGlobalAverage, calculateTotalSum, findMaxValue, findMinValue, isDiagonalMatrix } from "./matrixOperations";

const math = create(all, { number: "BigNumber", precision: 20 });

/**
 * Procesa dos matrices `Q` y `R`, realizando operaciones matemáticas avanzadas con `mathjs`.
 *
 * Calcula:
 * - Multiplicación `Q * R`
 * - Valor máximo y mínimo
 * - Suma total y promedio
 * - Verificación de matriz diagonal
 *
 * @param {any} matrixData - Objeto con matrices `q` y `r`
 * @returns {Promise<object>} Resultados de las operaciones matemáticas
 * @throws {Error} Si ocurre un error en el procesamiento
 */
export const processMatrix = async (matrixData: any) => {
  try {
    // Convertir matrices a `mathjs.Matrix` para cálculos precisos
    const q = math.matrix(matrixData.q) as Matrix;
    const r = math.matrix(matrixData.r) as Matrix;

    const result = math.multiply(q, r);

    const [maxQ, maxR, minQ, minR, sumQ, sumR, average, isDiagQ, isDiagR] = await Promise.all([
      findMaxValue(q),
      findMaxValue(r),
      findMinValue(q),
      findMinValue(r),
      calculateTotalSum(q),
      calculateTotalSum(r),
      calculateGlobalAverage(q, r),
      isDiagonalMatrix(q),
      isDiagonalMatrix(r)
    ]);

    return {
      max: math.max(maxQ, maxR).toNumber(),
      min: math.min(minQ, minR).toNumber(),
      average: average.toNumber(),
      totalSum: math.add(sumQ, sumR).toNumber(),
      isDiagonal: isDiagQ || isDiagR
    };

  } catch (error) {
    throw new Error("Error al procesar la matriz con precisión mejorada");
  }
};