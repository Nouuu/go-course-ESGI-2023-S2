package main

import (
	"03-fonctions-packages.3.3/car"
	"fmt"
	"log"
)

func main() {
	myCar, err := car.New("Toyota", 2010, "blue", 1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("La voiture a %d ans.\n", myCar.Age())
	err = myCar.UpdateYear(2009)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("La voiture a %d ans.\n", myCar.Age())

	// Comparaison de deux voitures
	myCar2, err := car.New("Toyota", 2010, "blue", 1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(myCar.Equals(*myCar2))
	err = myCar2.UpdateYear(2009)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(myCar.Equals(*myCar2))

}
