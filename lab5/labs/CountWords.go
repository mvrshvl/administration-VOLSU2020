package labs

import (
	"ADMINS/lab5/common"
	"log"
	"regexp"
)

func GetCountWords(path, name string) int {
	goFunc := func(data interface{}, resultChan chan interface{}) {
		dataString := common.InterfaceToString(data)
		pattern := `\b[[:alnum:]]*\b`
		re := regexp.MustCompile(pattern)
		array := re.FindAllString(dataString, -1)
		result := len(array)
		log.Println("GOROUTINE RESULT:", result)
		resultChan <- result
		return
	}

	dataChan := make(chan interface{})
	go common.ReadFile(path, name, dataChan)
	results := common.ConcurencyGetResult(dataChan, goFunc)
	var result int
	for _, res := range results {
		result += common.InterfaceToInt(res)
	}
	return result
}
