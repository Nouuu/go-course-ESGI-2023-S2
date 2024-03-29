package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/mydatabase")
	handleErr(err)

	rows, err := db.Query("SELECT id, name FROM test")
	handleErr(err)

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		handleErr(err)

		fmt.Println("Ligne :", id, name)
	}

	_, err = db.Exec("insert into test (name) values (?)", "Margot")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	/*	//myCar := car.Car{brand: "Toyota", year: 2010}
		myCar := car.New("Toyota", 2010, "red", 6)
		myCar2 := car.New("Toyota2", 2012, "red", 6)

		cars := []car.Car{*myCar, *myCar2}

		myJsonCars, err := json.MarshalIndent(cars, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		os.WriteFile("car.json", myJsonCars, 0644)
		var cars2 = make([]car.Car, 0)
		jsonData, _ := os.ReadFile("car.json")
		err = json.Unmarshal(jsonData, &cars2)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(cars2)*/
}
