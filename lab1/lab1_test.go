package lab1

import (
	"reflect"
	"sort"
	"testing"
)

func sumElem(array []int) int {
	return len(array)
}
func TestSumElem(t *testing.T) {
	sum := sumElem([]int{1, 2, 3, 4})
	if sum != 4 {
		t.Fatal()
	}
}
func sumPositiveNegative(array []int) (positive int, negative int) {
	for _, elem := range array {
		if elem < 0 {
			negative++
		} else {
			positive++
		}
	}
	return positive, negative
}
func TestSumPosNeg(t *testing.T) {
	positive, negative := sumPositiveNegative([]int{1, -1, 2, -2, 3, -3})
	if positive != 3 {
		t.Fatal()
	}
	if negative != 3 {
		t.Fatal()
	}
}
func getMin(array []int, k int) []int {
	sort.Ints(array)
	return array[:k]
}

func getMax(array []int, k int) []int {
	sort.Ints(array)
	return array[len(array)-k:]
}

func TestMinMax(t *testing.T) {
	array := []int{
		1, 5, 2, 3, 6, 7,
	}
	min := getMin(array, 2)
	if !reflect.DeepEqual(min, []int{1, 2}) {
		t.Fatal()
	}
	max := getMax(array, 2)
	if !reflect.DeepEqual(max, []int{6, 7}) {
		t.Fatal()
	}

}
