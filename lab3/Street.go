package lab3

import "errors"

type Street struct {
	Name   string
	Houses []House
}

func (street *Street) AddPerson(numberHouse uint64, numberApartment uint64, fio string, numbers string) error {
	for i, house := range street.Houses {
		if numberHouse == house.Number {
			err := house.AddPerson(numberApartment, fio, numbers)
			if err != nil {
				return err
			}
			street.Houses[i] = house
			return nil
		}
	}
	return errors.New("HOUSE NOT FOUND")
}
func (street *Street) RemovePerson(numberHouse uint64, numberApartment uint64, fio string) error {
	for _, house := range street.Houses {
		if numberHouse == house.Number {
			err := house.RemovePerson(numberApartment, fio)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("HOUSE NOT FOUND")
}
func (street *Street) GetPerson(numberHouse uint64, numberApartment uint64, fio string) (Person, error) {
	for _, house := range street.Houses {
		if numberHouse == house.Number {
			person, err := house.GetPerson(numberApartment, fio)
			if err != nil {
				return Person{}, err
			}
			return person, err
		}
	}
	return Person{}, errors.New("HOUSE NOT FOUND")
}
func (street *Street) AddHouse(h House) error {
	for _, house := range street.Houses {
		if h.Number == house.Number {
			return errors.New("HOUSE ALREADY EXIST")
		}
	}
	street.Houses = append(street.Houses, h)
	return nil
}
func (street *Street) GetHouse(number uint64) (House, error) {
	for _, house := range street.Houses {
		if number == house.Number {
			return house, nil
		}
	}
	return House{}, errors.New("House Not Found")
}
func (street *Street) RemoveHouse(number uint64) error {
	for i, house := range street.Houses {
		if number == house.Number {
			street.Houses = append(street.Houses[:i], street.Houses[i+1:]...)
			return nil
		}
	}
	return errors.New("House Not Found")
}

func (street *Street) GetFullInfoByPerson(fio string, numbers string) (full FullInfo, err error) {
	for _, house := range street.Houses {
		fullInfo, err := house.GetFullInfoByPerson(fio, numbers)
		if err != nil {
			continue
		} else {
			fullInfo.House = house
			return fullInfo, nil
		}

	}
	return FullInfo{}, errors.New("NOT FOUND")
}

func (street *Street) GetSum() int {
	var sum int
	for _, house := range street.Houses {
		sum += house.GetSum()
	}
	return sum
}
