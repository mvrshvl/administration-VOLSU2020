package labs

import (
	"ADMINS/lab5/common"
	"log"
	"sort"
)

/* 9
Реализовать параллельный поиск минимального положительного элемента числовой
последовательности.
*/
func GetMinPositiveNumber(array []int, threads int) (min int) {
	dataChan := make(chan interface{})
	maxSize := len(array) / 2
	if threads > maxSize {
		threads = maxSize
	} else if threads < 1 {
		threads = 1
	}
	part := len(array) / threads

	goFunc := func(data interface{}, resultChan chan interface{}) {
		dataArray := common.InterfaceToInts(data)
		positiveArray := []int{}
		for _, elem := range dataArray {
			if elem > 0 {
				positiveArray = append(positiveArray, elem)
			}
		}
		sort.Ints(positiveArray)
		log.Println("GOROUTINE RESULT:", positiveArray[0])
		resultChan <- positiveArray[0]
	}
	sendParts := func() {
		for i := 0; i < threads; i++ {
			from := part * i
			to := from + part
			if i == threads-1 {
				to = len(array)
			}
			dataChan <- array[from:to]
		}
		close(dataChan)
	}

	go sendParts()
	results := common.ConcurencyGetResult(dataChan, goFunc)
	resultInts := []int{}
	for _, r := range results {
		resultInts = append(resultInts, common.InterfaceToInt(r))
	}
	sort.Ints(resultInts)
	return resultInts[0]
}
