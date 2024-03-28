# Introduction - Variables & Opérateurs arithmétiques

## Exercice 1

```go
package main

import (
	"fmt"
)

func main() {
	var height int

	fmt.Println(height)
}
```

## Exercice 2

```go
package main

import (
	"fmt"
)

func main() {
	i := 314
	f := 3.14
	s := "Hello"
	b := true

	fmt.Println(
		"i:", i,
		"f:", f,
		"s:", s,
		"b:", b,
	)
}
```

## Exercice 3

```go
package main

import "fmt"

func main() {
	color, color2 := "red", "blue"

	color, color2 = "orange", "green"

	fmt.Println(color, color2)
}
```

## Exercice 4

```go
package main

import "fmt"

func main() {
	a, b := 10, 5.5
	fmt.Println(float64(a) + b)
}
```

## Exercice 5

```go
package main

import "fmt"

func main() {
	const taxe = 0.08

	var montantTotal float64
	var pourcentageRemise float64
	var montantFinal float64

	fmt.Println("Entrez le montant total de l'achat : ")
	fmt.Scan(&montantTotal)

	fmt.Println("Entrez le pourcentage de remise : ")
	fmt.Scan(&pourcentageRemise)

	montantRemise := montantTotal * pourcentageRemise / 100
	montantFinal = montantTotal - montantRemise + (montantTotal * taxe)

	fmt.Println("Montant final à payer :", montantFinal)
}
```

