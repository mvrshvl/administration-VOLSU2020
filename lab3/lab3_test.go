package lab3

import (
	"fmt"
	"sort"
	"testing"
)

// #2 Почтовые адреса
func TestPost(t *testing.T) {
	cityName := "MOSCOW"
	streetName := "KRASNAYA"
	var houseNumber uint64 = 1
	var index uint64 = 100100
	var count uint64 = 45
	var floors uint64 = 5
	var numApartment uint64 = 1
	var floorApartment uint64 = 1
	fio := "PETRENKO PETR PETROVICH"
	numbers := "88005553535"
	htype := B

	house := House{
		Number:      houseNumber,
		Index:       index,
		CountPeople: count,
		Floors:      floors,
		HouseType:   htype,
	}
	apartment := Apartment{
		Number: numApartment,
		Floor:  floorApartment,
	}
	street := Street{
		Name: streetName,
	}
	city := City{
		Name: cityName,
	}

	// Добавление
	err := house.AddApartment(apartment)
	err = street.AddHouse(house)
	err = city.AddStreet(street)
	post := newPost()
	post.AddCity(city)
	err = post.AddPerson(cityName, streetName, houseNumber, numApartment, fio, numbers)

	p1, err := post.toString()
	if err != nil {
		t.Fatal(err)
	}

	//выгрузка/загрузка
	err = post.Export("myPost", "JSON/")
	if err != nil {
		t.Fatal(err)
	}

	post = Post{}
	err = post.Import("myPost", "JSON/")
	if err != nil {
		t.Fatal(err)
	}

	p2, err := post.toString()

	if p1 != p2 {
		t.Fatal("IMPORT != EXPORT")
	}

	person, err := post.GetPerson(cityName, streetName, houseNumber, numApartment, fio)
	if err != nil {
		t.Fatal(err)
	}
	// извлечение объекта человека
	if person.Numbers != numbers {
		t.Fatal()
	}
	//Извлечение полной информации о человеке
	fullInfo, err := post.GetFullInfoByPerson(fio, numbers)
	if err != nil {
		t.Fatal(err)
	}
	if fullInfo.City.Name != cityName {
		t.Fatal()
	}
	//Подсчет числа адресов
	sumPeople, err := post.GetSum(cityName)
	if err != nil {
		t.Fatal(err)
	}
	if sumPeople != 1 {
		t.Fatal()
	}
	// Удаление адресата
	err = post.RemovePerson(cityName, streetName, houseNumber, numApartment, fio)
}
func Test1(t *testing.T) {
	data := []float64{1, 2, 0}
	sortt(data)
	fmt.Print(data)
}
func sortt(data []float64) {
	sort.Float64s(data)
	fmt.Print(data)
}
