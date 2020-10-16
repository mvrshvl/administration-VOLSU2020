package lab3

import "testing"

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
	if person.Numbers != numbers {

		t.Fatal()
	}

	fullInfo, err := post.GetFullInfoByPerson(fio, numbers)
	if err != nil {
		t.Fatal(err)
	}

	if fullInfo.City.Name != cityName {
		t.Fatal()
	}

	sumPeople, err := post.GetSum(cityName)
	if err != nil {
		t.Fatal(err)
	}
	if sumPeople != 1 {
		t.Fatal()
	}

}
