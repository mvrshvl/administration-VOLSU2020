package labs

import (
	common "ADMINS/lab5/common"
	"log"
	"regexp"
)

//Реализовать параллельное вычисление слова максимальной длины в заданном тексте.

func GetWordMaxLen(path, name string) string {
	goFunc := func(data interface{}, resultChan chan interface{}) {
		dataString := common.InterfaceToString(data)
		pattern := `\b[[:alnum:]]*\b`
		re := regexp.MustCompile(pattern)
		array := re.FindAllString(dataString, -1)
		var result string
		for _, word := range array {
			if len(word) > len(result) {
				result = word
			}
		}
		log.Println("GOROUTINE RESULT:", result)
		resultChan <- result
		return
	}
	dataChan := make(chan interface{})
	go common.ReadFile(path, name, dataChan)
	results := common.ConcurencyGetResult(dataChan, goFunc)

	var result string
	for _, word := range results {
		w := common.InterfaceToString(word)
		if len(w) > len(result) {
			result = w
		}
	}
	return result
}
