package tests

import (
	"reflect"
	"testing"

	rref "github.com/tokarzkj/rref-explorer/m/v2/domain"
)

func TestRowAddition(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {-2, -8, -7}}
	result := rref.CalculateRow(matrix[1], -2, matrix[0])
	expected_result := []int{0, -4, -1}

	if !reflect.DeepEqual(result, expected_result) {
		t.Error("The arrays do not match", result)
	}
}

func TestRowSubtraction(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {2, 8, 7}}
	result := rref.CalculateRow(matrix[1], 2, matrix[0])
	expected_result := []int{0, 4, 1}

	if !reflect.DeepEqual(result, expected_result) {
		t.Error("The arrays do not match", result)
	}
}

func TestMtxAddition(t *testing.T) {
	mtx := [][]int{{1, 2, 3}, {-2, -8, -7}}
	mtx = rref.UpdateMatrix(1, -2, 0, mtx)
	expected_mtx := [][]int{{1, 2, 3}, {0, -4, -1}}

	if !reflect.DeepEqual(mtx, expected_mtx) {
		t.Error("The arrays do not match", mtx)
	}
}

func TestMtxSubstraction(t *testing.T) {
	mtx := [][]int{{1, 2, 3}, {2, 8, 7}}
	mtx = rref.UpdateMatrix(1, 2, 0, mtx)
	expected_mtx := [][]int{{1, 2, 3}, {0, 4, 1}}

	if !reflect.DeepEqual(mtx, expected_mtx) {
		t.Error("The arrays do not match", mtx)
	}
}

func TestSwapRows(t *testing.T) {
	mtx := [][]int{{1, 2, 3}, {-2, -8, -7}}
	mtx = rref.SwapRows(1, 0, mtx)
	expected_mtx := [][]int{{-2, -8, -7}, {1, 2, 3}}

	if !reflect.DeepEqual(mtx, expected_mtx) {
		t.Error("The arrays do not match", mtx)
	}
}
