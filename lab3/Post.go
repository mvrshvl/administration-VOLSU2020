package lab3

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type FullInfo struct {
	City
	Street
	House
	Apartment
	Person
}

type Post struct {
	Cities map[string]City
}

func newPost() Post {
	return Post{Cities: make(map[string]City)}
}

func (post *Post) AddPerson(cityName string, streetName string, numberHouse uint64, numberApartment uint64, fio string, numbers string) error {
	if city, ok := post.Cities[cityName]; ok {
		err := city.AddPerson(streetName, numberHouse, numberApartment, fio, numbers)
		if err != nil {
			return err
		}
		post.Cities[cityName] = city
	} else {
		return errors.New("CITY NOT FOUND")
	}
	return nil
}
func (post *Post) GetPerson(cityName string, streetName string, numberHouse uint64, numberApartment uint64, fio string) (Person, error) {
	if city, ok := post.Cities[cityName]; ok {
		person, err := city.GetPerson(streetName, numberHouse, numberApartment, fio)
		if err != nil {
			return Person{}, err
		}
		return person, nil
	} else {
		return Person{}, errors.New("CITY NOT FOUND")
	}
}
func (post *Post) RemovePerson(cityName string, streetName string, numberHouse uint64, numberApartment uint64, fio string) error {
	if city, ok := post.Cities[cityName]; ok {
		err := city.RemovePerson(streetName, numberHouse, numberApartment, fio)
		if err != nil {
			return err
		}
	} else {
		return errors.New("CITY NOT FOUND")
	}
	return nil
}

func (post *Post) AddCity(city City) {
	post.Cities[city.Name] = city
}
func (post *Post) GetCity(name string) (City, error) {
	if _, ok := post.Cities[name]; ok {
		return post.Cities[name], nil
	} else {
		return City{}, errors.New("City not found")
	}
}
func (post *Post) RemoveCity(name string) error {
	if _, ok := post.Cities[name]; ok {
		delete(post.Cities, name)
		return nil
	} else {
		return errors.New("City not found")
	}
}
func (post Post) toString() (string, error) {
	res, err := json.Marshal(post)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
func (post Post) Export(filename string, path string) error {
	res, err := json.Marshal(post)
	if err != nil {
		return err
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0777)
		if err != nil {
			return err
		}
	}
	err = ioutil.WriteFile(path+filename, res, 0644)
	return err
}

func (post *Post) Import(filename, path string) error {
	data, err := ioutil.ReadFile(path + filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &post)
	return err
}

func (post *Post) GetFullInfoByPerson(fio, numbers string) (full FullInfo, err error) {

	for _, city := range post.Cities {
		full, err = city.GetFullInfoByPerson(fio, numbers)
		if err != nil {
			continue
		} else {
			full.City = city
			return full, nil
		}
	}
	return FullInfo{}, errors.New("NOT FOUND")
}

func (post Post) GetSum(name string) (int, error) {
	city, ok := post.Cities[name]
	if !ok {
		return 0, errors.New("CITY NOT FOUND")
	}
	sum := city.GetSum()
	return sum, nil
}
