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

## Exercice 2

```go
package main

import (
"fmt"
"os"
)

func main() {
animalsHabitat := make(map[string]map[string]int)

// Ajout des données pour la forêt
animalsHabitat["forêt"] = map[string]int{
"Renard":   3,
"Écureuil": 5,
"Cerf":     2,
}

// Ajout des données pour la savane
animalsHabitat["savane"] = map[string]int{
"Lion":   4,
"Zèbre":  7,
"Girafe": 3,
}

// Ajout des données pour l'océan
animalsHabitat["océan"] = map[string]int{
"Dauphin":       6,
"Requin":        2,
"Poisson-clown": 9,
}

// Ajout des données pour la montagne
animalsHabitat["montagne"] = map[string]int{
"Chamois":  5,
"Aigle":    2,
"Marmotte": 4,
}

// Récupération de l'entrée utilisateur
args := os.Args[1:]
if len(args) < 1 {
fmt.Println("Veuillez entrer un nom d'habitat pour obtenir la liste des animaux et leur nombre.")
return
}

// Récupération de l'habitat
habitat := args[0]
animals := animalsHabitat[habitat]
if animals == nil {
fmt.Printf("Désolé, je n'ai aucune information sur %q.\n", habitat)
return
}

// Affichage des animaux
fmt.Printf("~~~ Animaux de la %s ~~~\n\n", habitat)
for animal, nombre := range animals {
fmt.Printf("\t+ %s (%d)\n", animal, nombre)
}
}
```