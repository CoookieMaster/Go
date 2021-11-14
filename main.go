package main

import (
	"database/sql"
	_ "database/sql"
	_ "database/sql/driver"
	"fmt"
	_ "os"

	_ "github.com/go-sql-driver/mysql"
)

//Константы цен на такси
const CarPrice = 50
const LagCarPrice = 100

//Интерфейс
type Taxi interface {
	Go()
}

//Легковое такси
type Car struct {
	driver      string
	number      string
	isavailable bool
}

//Определение метода Go()
func (car Car) Go(pass Passenger) {
	println(" приедет легковое такси")
	fmt.Printf("Имя водителя %v, номер машины %v\n", car.driver, car.number)
	fmt.Printf("Цена поездки %v\n", CarPrice*pass.distance)
	println()
}

//Грузовое такси
type LaggegeCar struct {
	driver      string
	number      string
	isavailable bool
}

//Определение метода Go()
func (lagcar LaggegeCar) Go(pass Passenger) {
	println(" приедет грузовое такси")
	fmt.Printf("Имя водителя %v, номер машины %v\n", lagcar.driver, lagcar.number)
	fmt.Printf("Цена поездки %v\n", LagCarPrice*pass.distance)
	println()
}

//Структура пассажир
type Passenger struct {
	name       string
	passengers int
	laggege    float32
	distance   float32
}

func main() {

	db, err := sql.Open("mysql", "mysql:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	/*
		insert, err := db.Query("INSERT INTO `taxi`(`id`, `driver`, `number`, `isavailable`, `type`) VALUES ('1','Петя','ау584л','0','легковое'), ('2','Виталя','оу749в','1','легковое'), ('3','Виктор','пв225к','1','грузовое'),  ('4','Илья','лу124е','1','грузовое')")
		if err != nil {
			panic(err)
		}
		defer insert.Close()
	*/
	var cars = []*Car{
		{driver: "Петя", number: "ау584л", isavailable: false},
		{driver: "Виталя", number: "оу749в", isavailable: true},
	}

	var lagCars = []*LaggegeCar{
		{driver: "Виктор", number: "пв225к", isavailable: true},
		{driver: "Илья", number: "лу124е", isavailable: true},
	}

	var passengers = []Passenger{
		{name: "Фёдор", passengers: 4, distance: 40, laggege: 45},
		{name: "Мария", passengers: 4, distance: 40, laggege: 45},
		{name: "Миша", passengers: 2, distance: 10, laggege: 80},
		{name: "Борис", passengers: 2, distance: 10, laggege: 80},
		{name: "Катя", passengers: 5, distance: 100, laggege: 20},
	}

	/*fmt.Print("Введите число пассажиров: ")
	    fmt.Fscan(os.Stdin, &pass1.passengers)
	    fmt.Print("Введите расстояние: ")
	    fmt.Fscan(os.Stdin, &pass1.distance)
		fmt.Print("Введите вес багажа: ")
	    fmt.Fscan(os.Stdin, &pass1.laggege)
	    fmt.Println(pass1.passengers, pass1.distance, pass1.laggege)
	*/

	fmt.Printf("Поиск...\n")
	println()
	count := 0
	for _, value := range passengers {
		if value.passengers > 2 && value.passengers <= 4 && value.laggege <= 50 {
			for _, value2 := range cars {
				if value2.isavailable == true {
					println("Заказ", count+1)
					fmt.Printf("Для клиента %v", value.name)
					value2.Go(value)
					value2.isavailable = false
					break
				}
			}
		} else if value.passengers <= 2 && value.passengers > 0 {
			for _, value3 := range lagCars {
				if value3.isavailable == true {
					println("Заказ", count+1)
					fmt.Printf("Для клиента %v", value.name)
					value3.Go(value)
					value3.isavailable = false
					break
				}
			}
		} else {
			println("Заказ", count+1)
			panic("Некорректное число пассажиров!")
		}
		count++
	}

	//Выборка данных
	res, err := db.Query("SELECT `driver`, `number`, `isavailable` FROM 'taxi'")
	if err != nil {
		panic(err)
	}

	for res.Next() {
		var car Car
		err = res.Scan(&car.driver, &car.number, &car.isavailable)
		if err != nil {
			panic(err)
		}
		fmt.Printf(fmt.Sprintf("Водитель: %v, машина № %v свободна: %v. Тип - легковой\n", car.driver, car.number, car.isavailable))
	}

	for res.Next() {
		var lagCar LaggegeCar
		err = res.Scan(&lagCar.driver, &lagCar.number, &lagCar.isavailable)
		if err != nil {
			panic(err)
		}
		fmt.Printf(fmt.Sprintf("Водитель: %v, машина № %v свободна: %v. Тип - грузовой\n", lagCar.driver, lagCar.number, lagCar.isavailable))
	}

	//Попытка добавления данных через цикл
	/*for  _, value := range cars {
		insert, err := db.Query("INSERT INTO 'taxi' ('driver','number', 'isavailable', 'type') VALUES('%v', '%v', '%v', 'легковое')", value.driver,
		value.number, 0)
		if err != nil {
			panic(err)
		}
		defer insert.Close()
	}*/
}
