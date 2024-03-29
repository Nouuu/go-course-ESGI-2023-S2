package main

import (
	"03-fonctions-packages.3.1/car"
	"fmt"
)

func main() {
	myCar := car.New("Toyota", 2010, "blue", 1)
	fmt.Printf("La voiture a %d ans.", myCar.Age())
	myCar.UpdateYear(2009)
	fmt.Printf("La voiture a %d ans.", myCar.Age())

	// Comparaison de deux voitures
	myCar2 := car.New("Toyota", 2010, "blue", 1)
	fmt.Println(myCar.Equals(*myCar2))
	myCar2.UpdateYear(2009)
	fmt.Println(myCar.Equals(*myCar2))
}
