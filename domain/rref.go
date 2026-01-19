package domain

func UpdateMatrix(tar_row_idx int, op int, op_row_idx int, mtx [][]int) [][]int {
	tar_row := mtx[tar_row_idx]
	op_row := mtx[op_row_idx]

	tar_row = CalculateRow(tar_row, op, op_row)
	mtx[tar_row_idx] = tar_row

	return mtx
}

func CalculateRow(tar_row []int, op int, op_row []int) []int {
	var arraySize = len(tar_row)
	var results = make([]int, arraySize)
	for i := range tar_row {
		results[i] = tar_row[i] - (op * op_row[i])
	}

	return results
}

func SwapRows(r2 int, r1 int, matrix [][]int) [][]int {
	var tempRow = matrix[r1]
	matrix[r1] = matrix[r2]
	matrix[r2] = tempRow

	return matrix
}
