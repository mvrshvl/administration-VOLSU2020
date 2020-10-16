package lab2

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"strconv"
	"testing"
)

func getCountWords(length int, filename string, path string) (int, error) {
	data, err := ioutil.ReadFile(path + filename)
	if err != nil {
		return 0, err
	}
	lenStr := strconv.Itoa(length)
	pattern := `\b[[:alnum:]]{` + lenStr + `}\b`
	re := regexp.MustCompile(pattern)
	array := re.FindAllString(string(data), -1)

	return len(array), nil
}

func TestGetCountWords(t *testing.T) {
	count, err := getCountWords(4, "GCW.txt", "TEXTS/")
	if err != nil {
		t.Fatal()
	}
	if count != 7 {
		t.Fatal()
	}
}

func getWordsWithNLen(length int, filename string, path string) ([]string, error) {
	data, err := ioutil.ReadFile(path + filename)
	if err != nil {
		return []string{}, err
	}
	lenStr := strconv.Itoa(length)
	pattern := `\b[[:alnum:]]{` + lenStr + `}\b`
	re := regexp.MustCompile(pattern)
	array := re.FindAllString(string(data), -1)

	return array, nil
}
func TestGetWords(t *testing.T) {
	words, err := getWordsWithNLen(5, "GCW.txt", "TEXTS/")
	if err != nil {
		t.Fatal()
	}
	waitResult := []string{"ttttt", "fffff", "eeeee"}
	if !reflect.DeepEqual(waitResult, words) {
		t.Fatal("WAIT", waitResult, "EXPECT", words)
	}
}

func getWordWithMaxLen(filename string, path string) (string, error) {
	data, err := ioutil.ReadFile(path + filename)
	if err != nil {
		return "", err
	}
	pattern := `\b[[:alnum:]]*\b`
	re := regexp.MustCompile(pattern)
	array := re.FindAllString(string(data), -1)
	var result string
	for _, word := range array {
		if len(word) > len(result) {
			result = word
		}
	}
	return result, nil
}
func TestGetMaxLen(t *testing.T) {
	result, err := getWordWithMaxLen("GCW.txt", "TEXTS/")
	if err != nil {
		t.Fatal()
	}
	fmt.Println(result)
	//if length != 7{
	//	t.Fatal("WAIT : 7","EXPECT :",length)
	//}
}

func getWordSoglN(length int, filename string, path string) ([]string, error) {
	data, err := ioutil.ReadFile(path + filename)
	if err != nil {
		return []string{}, err
	}
	lenStr := strconv.Itoa(length)
	sogl := `[qwrtpsdfghjklzxcvbnm]{1}`
	pattern := `\b(*` + sogl + `*){` + lenStr + `}\b`
	re := regexp.MustCompile(pattern)
	array := re.FindAllString(string(data), -1)

	return array, nil
}
func TestGetSogl(t *testing.T) {
	_, err := getWordSoglN(4, "GCW.txt", "TEXTS/")
	if err != nil {
		t.Fatal()
	}
}
