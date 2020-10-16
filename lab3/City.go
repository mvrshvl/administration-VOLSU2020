package lab3

import "errors"

type City struct {
	Name    string
	Streets []Street
}

func (city *City) AddPerson(streetName string, numberHouse uint64, numberApartment uint64, fio string, numbers string) error {
	for i, c := range city.Streets {
		if c.Name == streetName {
			err := c.AddPerson(numberHouse, numberApartment, fio, numbers)
			if err != nil {
				return err
			}
			city.Streets[i] = c
			return nil
		}
	}
	return errors.New("CITY NOT FOUND")
}
func (city *City) GetPerson(streetName string, numberHouse uint64, numberApartment uint64, fio string) (Person, error) {
	for _, c := range city.Streets {
		if c.Name == streetName {
			person, err := c.GetPerson(numberHouse, numberApartment, fio)
			if err != nil {
				return Person{}, err
			}
			return person, nil
		}
	}
	return Person{}, errors.New("CITY NOT FOUND")
}
func (city *City) RemovePerson(streetName string, numberHouse uint64, numberApartment uint64, fio string) error {
	for _, c := range city.Streets {
		if c.Name == streetName {
			err := c.RemovePerson(numberHouse, numberApartment, fio)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("CITY NOT FOUND")
}

func (city *City) AddStreet(street Street) error {
	for _, street := range city.Streets {
		if street.Name == street.Name {
			return errors.New("STREET ALREAY EXIST")
		}
	}
	city.Streets = append(city.Streets, street)
	return nil
}
func (city *City) GetStreet(name string) (Street, error) {
	for _, street := range city.Streets {
		if street.Name == name {
			return street, nil
		}
	}
	return Street{}, errors.New("STREET NOT FOUND")
}
func (city *City) RemoveStreet(name string) error {
	for i, street := range city.Streets {
		if street.Name == name {
			city.Streets = append(city.Streets[:i], city.Streets[i+1:]...)
		}
	}
	return errors.New("STREET NOT FOUND")
}

func (city *City) GetFullInfoByPerson(fio string, numbers string) (full FullInfo, err error) {
	for _, street := range city.Streets {
		fullInfo, err := street.GetFullInfoByPerson(fio, numbers)
		if err != nil {
			continue
		} else {
			fullInfo.Street = street
			return fullInfo, nil
		}

	}
	return FullInfo{}, errors.New("NOT FOUND")
}

func (city *City) GetSum() int {
	var sum int
	for _, street := range city.Streets {
		sum += street.GetSum()
	}
	return sum
}
