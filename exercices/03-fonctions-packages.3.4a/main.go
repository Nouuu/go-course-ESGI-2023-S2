package main

import (
	"03-fonctions-packages.3.4a/car"
	"03-fonctions-packages.3.4a/dealership"
	"log"
)

func main() {
	car1, err := car.New("Toyota", 2010, "blue", 1)
	if err != nil {
		log.Fatal(err)
	}
	car2, err := car.New("Peugeot", 2015, "red", 2)
	if err != nil {
		log.Fatal(err)
	}
	car3, err := car.New("Renault", 2018, "green", 3)
	if err != nil {
		log.Fatal(err)
	}

	myDealership := dealership.New()
	myDealership.AddCar(*car1)
	myDealership.AddCar(*car2)
	myDealership.AddCar(*car3)

	log.Printf("%+v\n", myDealership)
}
