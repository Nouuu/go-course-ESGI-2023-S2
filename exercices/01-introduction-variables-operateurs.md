https://github.com/Nouuu/go-course-ESGI-2023-S2


# Introduction - Variables & Opérateurs arithmétiques

## Exercice 1

- Déclarez et affichez une variable de type int.
- Le nom de la variable déclarée devrait être : `height`.

```go
package main

func main() {
    // Déclarez la variable `height` ici

    // Affichez la variable `height` ici
}
```

**Résultat attendu**

```
0
```

## Exercice 2

Déclarez puis affichez quatre variables en utilisant l'instruction de déclaration courte.

```go
package main

func main() {
    // AJOUTEZ VOS DÉCLARATIONS ICI
    //

    // DÉCOMMENTEZ ENSUITE LE CODE CI-DESSOUS

    // fmt.Println(
    //  "i:", i,
    //  "f:", f,
    //  "s:", s,
    //  "b:", b,
    // )
}
```

**Résultat attendu**

```
i: 314 f: 3.14 s: Hello b: true
```

## Exercice 3

- Changez `color` en "orange" et `color2` en "green" sur une même ligne, en utilisant une affectation multiple.
- Affichez les variables.

```go
package main

func main() {
    // DÉCOMMENTEZ LE CODE CI-DESSOUS :

    // color, color2 := "red", "blue"

    // Effectuez l'affectation multiple ici pour changer les valeurs de `color` et `color2`.

    // Affichez les variables ici.
}
```

**Résultat attendu**

```
orange green
```

## Exercice 4

Corrigez le code en utilisant une expression de conversion.

```go
package main

func main() {
    // a, b := 10, 5.5
    // Utilisez une expression de conversion pour corriger l'opération ci-dessous.
    // fmt.Println(a + b)
}
```

**Résultat attendu**

```
15.5
```

## Exercice 5

Écrivez un programme en Go qui calcule le montant final à payer pour un achat avec une remise et une taxe.

**Étapes :**

- Déclarez une constante **`taxe`** avec la valeur **0.08**.
- Déclarez les variables **`montantTotal`** (type **float64**) et **`pourcentageRemise`** (type **float64**).
- Demandez à l'utilisateur d'entrer le montant total de l'achat en utilisant `fmt.Print` et `fmt.Scan`.
- Demandez à l'utilisateur d'entrer le pourcentage de remise en utilisant `fmt.Print` et `fmt.Scan`.
- Calculez le montant de la remise en multipliant **`montantTotal`** par **`pourcentageRemise`** divisé par 100.
- Calculez le montant final à payer en soustrayant la remise et en ajoutant la taxe au montant total.
- Affichez le montant final à payer en utilisant `fmt.Println`.

**Utilisation de `fmt.Scan`**

```go
package main

import "fmt"

func main() {
    var valeurEntiere int

    fmt.Print("Entrez un nombre entier : ")
    fmt.Scan(&valeurEntiere)

    fmt.Println("Vous avez entré :", valeurEntiere)
}
```

Cette application attend deux arguments sur la ligne de commande. Elle les convertit en float64 et les affiche à
l'utilisateur. Si la conversion échoue pour l'un des deux arguments, le programme affiche une erreur et se termine.

**Exemple d'exécution :**

```shell
$ go run main.go 
Entrez le montant total de l'achat : 8
Entrez le pourcentage de remise : 10
Montant final à payer : 7.36
```