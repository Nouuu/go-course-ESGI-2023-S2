package main

import (
	"03-fonctions-packages.3.4b/car"
	"03-fonctions-packages.3.4b/dealership"
	"log"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal("ERREUR : ", err)
	}
}

func main() {
	car1, err := car.New("Toyota", 2010, "blue", 1)
	handleErr(err)
	car2, err := car.New("Peugeot", 2015, "red", 2)
	handleErr(err)
	car3, err := car.New("Renault", 2018, "green", 3)
	handleErr(err)

	myDealership := dealership.New()
	myDealership.AddCar(*car1)
	myDealership.AddCar(*car2)
	myDealership.AddCar(*car3)

	log.Printf("%+v\n", myDealership)

	err = myDealership.RemoveCar(1)
	handleErr(err)
	err = myDealership.RemoveCar(5)
	handleErr(err)
}
