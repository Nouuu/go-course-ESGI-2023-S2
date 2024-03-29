package main

import (
	"03-fonctions-packages.3.2/car"
	"fmt"
)

func main() {
	myCar, err := car.New("Toyota", 2010, "blue", 1)
	if err != nil {
		return
	}
	fmt.Printf("La voiture a %d ans.", myCar.Age())
	err = myCar.UpdateYear(2009)
	if err != nil {
		return
	}
	fmt.Printf("La voiture a %d ans.", myCar.Age())

	// Comparaison de deux voitures
	myCar2, err := car.New("Toyota", 2010, "blue", 1)
	if err != nil {
		return
	}
	fmt.Println(myCar.Equals(*myCar2))
	err = myCar2.UpdateYear(2009)
	if err != nil {
		return
	}
	fmt.Println(myCar.Equals(*myCar2))

}
