package car

import (
	"fmt"
	"time"
)

// Car Structure Car
type Car struct {
	brand  string
	year   int
	color  string
	engine int
}

func New(brand string, year int, color string, engine int) (*Car, error) {
	if brand == "" {
		return nil, fmt.Errorf("la marque ne peut pas être vide")
	}
	if year > time.Now().Year() {
		return nil, fmt.Errorf("l'année %d est supérieure à l'année actuelle", year)
	}
	if color == "" {
		return nil, fmt.Errorf("la couleur ne peut pas être vide")
	}
	return &Car{brand: brand, year: year, color: color, engine: engine}, nil
}

// Equals compare deux voitures
func (c Car) Equals(other Car) bool {
	return c.brand == other.brand && c.year == other.year
}

// Age retourne l'age de la voiture
func (c Car) Age() int {
	return time.Now().Year() - c.year
}

// UpdateYear met à jour l'année de la voiture
func (c *Car) UpdateYear(newYear int) error {
	if newYear > time.Now().Year() {
		return fmt.Errorf("l'année %d est supérieure à l'année actuelle", newYear)
	}
	c.year = newYear
	return nil
}

// SetColor met à jour la couleur de la voiture
func (c *Car) SetColor(newColor string) error {
	if newColor == "" {
		return fmt.Errorf("la couleur ne peut pas être vide")
	}
	c.color = newColor
	return nil
}

// UpgradeEngine met à jour la puissance du moteur
func (c *Car) UpgradeEngine() {
	c.engine++
}
