# Dictionnaires (Maps)

## Exercice 1

```go
package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	houses := map[string][]string{
		"gryffondor":  {"weasley", "hagrid", "larrieu-lacoste", "dumbledore", "lupin"},
		"poufsouffle": {"wenlock", "scamander", "helga", "diggory", "bobo"},
		"serdaigle":   {"flitwick", "bagnold", "wildsmith", "montmorency"},
		"serpentard":  {"horace", "nigellus", "higgs", "bobo", "scorpius"},
		"bobo":        {"wizardry", "unwanted"},
	}

	// Suppression de la maison "bobo"
	delete(houses, "bobo")

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Veuillez entrer un nom de maison de Poudlard.")
		return
	}
	house := args[0]
	students := houses[house]
	if students == nil {
		fmt.Printf("Désolé. Je n'ai rien sur %q.\n", house)
		return
	}

	// Tri du clone
	clone := append([]string(nil), students...)
	sort.Strings(clone)

	fmt.Printf("~~~ Etudiants de %s ~~~\n\n", house)
	for _, student := range clone {
		fmt.Printf("\t+ %s\n", student)
	}
}
```