package common

import (
	"bufio"
	"log"
	"os"
)

// Основа для запуска конкурентного выполнения
func ReadFile(path, filename string, export chan interface{}) {
	defer close(export)
	file, err := os.Open(path + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		export <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func ConcurencyGetResult(importChan chan interface{}, f func(data interface{}, ch chan interface{})) (results []interface{}) {
	resultChan := make(chan interface{})
	var process uint64
	for data := range importChan {
		process++
		go f(data, resultChan)
	}

	for {
		process--
		result := <-resultChan
		results = append(results, result)
		if process == 0 {
			break
		}
	}
	return results
}

func InterfaceToString(data interface{}) string {
	switch result := data.(type) {
	case string:
		return result
	default:
		return ""
	}
}
func InterfaceToInt(data interface{}) int {
	switch result := data.(type) {
	case int:
		return result
	default:
		return 0
	}
}
func InterfaceToInts(data interface{}) []int {
	switch result := data.(type) {
	case []int:
		return result
	default:
		return []int{}
	}
}
