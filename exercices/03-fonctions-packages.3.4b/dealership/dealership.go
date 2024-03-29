package dealership

import (
	"03-fonctions-packages.3.4b/car"
	"fmt"
	"log"
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
