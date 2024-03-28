# Structures de contrôle

## Exercice 1

```go
package main

import "fmt"

func main() {
	age := 10

	if age > 60 {
		fmt.Println("Getting older")
	} else if age > 30 {
		fmt.Println("Getting wiser")
	} else if age > 20 {
		fmt.Println("Adulthood")
	} else if age > 10 {
		fmt.Println("Young blood")
	} else {
		fmt.Println("Booting up")
	}
}
```

## Exercice 2

```go
package main

import "fmt"

func main() {
	isSphere, radius := true, 200

	if isSphere && radius >= 200 {
		fmt.Println("C'est une grosse sphère.")
	} else {
		fmt.Println("Je ne sais pas.")
	}
}
```

## Exercice 3

```go
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Donnez moi un nom de mois")
		return
	}
	
	annee := time.Now().Year()
	bissextile := annee%4 == 0 && (annee%100 != 0 || annee%400 == 0)
	
	jours, mois := 28, os.Args[1]

	switch strings.ToLower(mois) {
	case "avril", "juin", "septembre", "novembre":
		jours = 30
	case "janvier", "mars", "mai", "juillet",
		"aout", "octobre", "decembre":
		jours = 31
	case "fevrier":
		if bissextile {
			jours = 29
		}
	default:
		fmt.Printf("%q n'est pas un mois.\n", mois)
		return
	}

	fmt.Printf("%q a %d jours.\n", mois, jours)
}
```

## Exercice 4

```go
package main

import (
	"fmt"
)

func main() {
	var sum int
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)
}
```

## Exercice 5

```go
func main() {
	i := 1
	sum := 0

	for ; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Print(i)
		if i != 10 {
			fmt.Print(" + ")
		}
		sum += i
	}
	fmt.Printf(" = %d\n", sum)
}
```