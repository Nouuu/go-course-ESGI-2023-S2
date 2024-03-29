package dealership

import (
	"03-fonctions-packages.3.4a/car"
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
