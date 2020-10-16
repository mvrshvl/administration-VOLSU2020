package lab3

import "errors"

type Apartment struct {
	Persons []Person
	Number  uint64
	Floor   uint64
}

func (apart *Apartment) AddPerson(FIO string, numbers string) {
	person := Person{
		FIO:     FIO,
		Numbers: numbers,
	}
	apart.Persons = append(apart.Persons, person)
}
func (apart *Apartment) GetPerson(fio string) (Person, error) {
	for _, person := range apart.Persons {
		if person.FIO == fio {
			return person, nil
		}
	}
	return Person{}, errors.New("Person not found")
}
func (apart *Apartment) RemovePerson(FIO string) {
	for i, person := range apart.Persons {
		if person.FIO == FIO {
			apart.Persons = append(apart.Persons[:i], apart.Persons[i+1:]...)
			return
		}
	}
}
