# Dictionnaires (Maps)

## Exercice 1

Créez un programme qui renvoie les étudiants en fonction du nom de la maison de Poudlard donné (voir les données ci-dessous).

Affichez les étudiants **triés** par nom.

"bobo" n'appartient pas à Poudlard, supprimez-le en utilisant la fonction `delete`, au début du programme.

**RESTRICTIONS**

- Ajoutez les données suivantes à votre dictionnaire (map) telles quelles.
  Ne les triez pas manuellement et ne les modifiez pas.

- Les tranches (slices) dans le dictionnaire (map) ne doivent pas être triées (modifiées).
  ASTUCE : Copiez-les.

**ASTUCE**

- Vous pouvez utiliser la fonction `delete` pour supprimer un élément d'une carte (map).
- Vous pouvez utiliser la fonction `sort.Strings` pour trier une tranche (slice) de chaînes de caractères.

```go
package main

func main() {
	// House        Student Name
	// ---------------------------
	// gryffondor       weasley
	// gryffondor       hagrid
	// gryffondor       larrieu-lacoste
	// gryffondor       dumbledore
	// gryffondor       lupin
	// poufsouffle      wenlock
	// poufsouffle      scamander
	// poufsouffle      helga
	// poufsouffle      diggory
	// serdaigle        flitwick
	// serdaigle        bagnold
	// serdaigle        wildsmith
	// serdaigle        montmorency
	// serpentard       horace
	// serpentard       sananes
	// serpentard       nigellus
	// serpentard       higgs
	// serpentard       scorpius
	// bobo             wizardry
	// bobo             unwanted
}
```

**RÉSULTAT ATTENDU**

```
go run main.go

Veuillez entrer un nom de maison de Poudlard.

go run main.go bobo

Désolé. Je ne sais rien sur "bobo".

go run main.go poufsouffle

~~~ Étudiants de poufsouffle ~~~

    + diggory
    + helga
    + scamander
    + wenlock
```

## Exercice 2

Créez un programme qui répertorie les animaux par leur habitat. Vous devrez également gérer des informations sur le nombre d'animaux de chaque espèce dans chaque habitat.

Ajoutez les données d'animaux et d'habitats suivantes à votre carte (map) comme suit :

- forêt : renard (3), écureuil (5), cerf (2)
- savane : lion (4), zèbre (7), girafe (3)
- océan : dauphin (6), requin (2), poisson-clown (9)
- montagne : chamois (5), aigle (2), marmotte (4)

**RESTRICTIONS**

+ Les habitats et les animaux doivent être stockés dans le dictionnaire (map) sans tri manuel.
+ Les tranches (slices) pour le nombre d'animaux de chaque espèce dans chaque habitat ne doivent pas être triées.

**RÉSULTAT ATTENDU**

```
go run main.go

Veuillez entrer un nom d'habitat pour obtenir la liste des animaux et leur nombre.

go run main.go Désert

Désolé, je n'ai aucune information sur "Désert".

go run main.go Forêt

~~~ Animaux dans la forêt ~~~

Renard (3)
Écureuil (5)
Cerf (2)
```