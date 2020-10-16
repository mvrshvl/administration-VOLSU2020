package lab5

import (
	"ADMINS/lab5/labs"
	"fmt"
	"testing"
)

func TestGetWordMaxLen(t *testing.T) {
	result := labs.GetWordMaxLen("TEXTS/", "t9.txt")
	if result != "1010101010" {
		t.Fatal(result)
	}
}

func TestGetCountWords(t *testing.T) {
	result := labs.GetCountWords("TEXTS/", "t7.txt")
	fmt.Println(result)
	if result != 10 {
		t.Fatal(result)
	}
}

func TestGetMinMaxElem(t *testing.T) {
	arr := []int{1, 0, -5, 24141, 12, -155, 244, 9, 12, -5541, 122423, 56575, 89, -2}
	min, max := labs.GetMinMaxElem(arr, 5)
	fmt.Println(min, max)

}

func TestGetMinPositiveElem(t *testing.T) {
	arr := []int{1, 123, 512, 24141, 12, 155, 244, 9, 12, 5541, 122423, 56575, 89, -2}
	min := labs.GetMinPositiveNumber(arr, 4)
	fmt.Println(min)

}
