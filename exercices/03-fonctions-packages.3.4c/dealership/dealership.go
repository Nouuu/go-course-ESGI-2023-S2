package dealership

import (
	"03-fonctions-packages.3.4c/car"
	"fmt"
	"log"
	"strconv"
)

type Dealership struct {
	cars []car.Car
}

func New() *Dealership {
	cars := make([]car.Car, 0)
	return &Dealership{cars: cars}
}

func (d *Dealership) AddCar(c car.Car) {
	log.Printf("Ajout d'une voiture dans la concession : %+v\n", c)
	d.cars = append(d.cars, c)
}

func (d *Dealership) RemoveCar(index int) error {
	log.Printf("Suppression d'une voiture dans la concession à l'index : %d\n", index)
	if index < 0 || index >= len(d.cars) {
		return fmt.Errorf("l'index %d est invalide", index)
	}
	newCars := make([]car.Car, 0, len(d.cars)-1)
	for i := 0; i < len(d.cars); i++ {
		if i != index {
			newCars = append(newCars, d.cars[i])
		}
	}
	d.cars = newCars
	return nil
}

func (d *Dealership) FindCarsByCriteria(criteria map[string]string) []car.Car {
	log.Printf("Recherche de voitures dans la concession avec les critères : %+v\n", criteria)
	cars := make([]car.Car, 0)
	for _, c := range d.cars {
		match := true
		for criteriaKey, criteriaValue := range criteria {
			switch criteriaKey {
			case "brand":
				if c.GetBrand() != criteriaValue {
					match = false
				}
			case "year":
				year := strconv.Itoa(c.GetYear())
				if year != criteriaValue {
					match = false
				}
			case "color":
				if c.GetColor() != criteriaValue {
					match = false
				}
			case "engine":
				engine := strconv.Itoa(c.GetEngine())
				if engine != criteriaValue {
					match = false
				}
			}
		}
		if match {
			cars = append(cars, c)
		}
	}
	return cars
}
