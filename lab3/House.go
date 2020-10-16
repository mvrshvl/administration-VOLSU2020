package lab3

import "errors"

type houseType uint8

const (
	A houseType = iota
	B
	C
)

type House struct {
	Number      uint64
	Index       uint64
	Apartments  []Apartment
	CountPeople uint64
	Floors      uint64
	HouseType   houseType
}

func (house *House) AddPerson(numberApartment uint64, fio string, numbers string) error {
	for i, apart := range house.Apartments {
		if apart.Number == numberApartment {
			apart.AddPerson(fio, numbers)
			house.Apartments[i] = apart
			return nil
		}
	}
	return errors.New("APARTMENT NOT FOUND")
}
func (house *House) GetPerson(numberApartment uint64, fio string) (Person, error) {
	for _, apart := range house.Apartments {
		if apart.Number == numberApartment {
			person, err := apart.GetPerson(fio)
			if err != nil {
				return Person{}, err
			}
			return person, nil
		}
	}
	return Person{}, errors.New("APARTMENT NOT FOUND")
}
func (house *House) RemovePerson(number uint64, fio string) error {
	for _, apart := range house.Apartments {
		if apart.Number == number {
			apart.RemovePerson(fio)
		}
	}
	return errors.New("APARTMENT NOT FOUND")
}

func (house *House) AddApartment(apartment Apartment) error {
	for _, apart := range house.Apartments {
		if apart.Number == apartment.Number {
			return errors.New("APARTMENT ALREADY EXIST")
		}
	}
	house.Apartments = append(house.Apartments, apartment)
	return nil
}
func (house *House) RemoveApartment(numberApartment uint64) error {
	for i, apart := range house.Apartments {
		if apart.Number == numberApartment {
			house.Apartments = append(house.Apartments[:i], house.Apartments[i+1:]...)
			return nil
		}
	}
	return errors.New("APARTMENT NOT EXIST")
}

func (house *House) GetFullInfoByPerson(fio string, numbers string) (full FullInfo, err error) {
	for _, apart := range house.Apartments {
		person, err := apart.GetPerson(fio)
		if err != nil {
			continue
		}
		if person.Numbers != numbers {
			continue
		} else {
			return FullInfo{
				Apartment: apart,
				Person:    person,
			}, nil
		}
	}
	return FullInfo{}, errors.New("NOT FOUND")
}

func (house *House) GetSum() int {
	var sum int
	for _, apart := range house.Apartments {
		sum += len(apart.Persons)
	}
	return sum
}
