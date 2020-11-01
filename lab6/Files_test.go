package lab6

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// 1. Найти файл максимального размера.
func GetMostLargeFile(path string) {
	type relustStruct struct {
		name     string
		size     int64
		filepath string
	}
	var result relustStruct
	filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {

		if wPath == path {
			return nil
		}

		if !info.IsDir() {
			f, err := os.Open(wPath)
			if err != nil {
				panic(err)
			}
			fi, err := f.Stat()
			if err != nil {
				panic(err)
			}
			if fi.Size() > result.size {
				result.size = fi.Size()
				result.name = fi.Name()
				result.filepath = wPath
			}
		}
		return nil
	})
	fmt.Println("FILENAME", result.name)
	fmt.Println("FILEPATH", result.filepath)
	fmt.Println("SIZE", result.size)

}
func GetFiles(path string, f func(f *os.File)) {
	filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {

		if wPath == path {
			return nil
		}

		if !info.IsDir() {
			file, err := os.Open(wPath)
			if err != nil {
				panic(err)
			}
			f(file)
		}
		return nil
	})
}
func TestGetMostLargeFile(t *testing.T) {
	GetMostLargeFile(".")
}

//2. Построить список имен файлов, размер которых находится в заданном диапазоне.
func GetFilesWithNSize(path string, from, to int64) {
	f := func(file *os.File) {
		fi, err := file.Stat()
		if err != nil {
			panic(err)
		}
		if fi.Size() > from && fi.Size() < to {
			fmt.Println("FILENAME", fi.Name())
			fmt.Println("SIZE", fi.Size())
		}
	}
	GetFiles(path, f)
}
func TestGetFilesWithNSize(t *testing.T) {
	//GetFilesWithNSize(".",0,10000)
	GetFilesWithNSize(".", 10000, 999999999)
}
