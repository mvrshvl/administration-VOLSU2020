package labs

import (
	"ADMINS/lab5/common"
	"log"
	"sort"
)

func GetMinMaxElem(array []int, threads int) (min, max int) {
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
		sort.Ints(dataArray)
		dataArray = append(dataArray[:1], dataArray[len(dataArray)-1])
		log.Println("GOROUTINE RESULT:", dataArray)
		resultChan <- dataArray
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
		resultInts = append(resultInts, common.InterfaceToInts(r)...)
	}
	sort.Ints(resultInts)
	return resultInts[0], resultInts[len(resultInts)-1]
}
