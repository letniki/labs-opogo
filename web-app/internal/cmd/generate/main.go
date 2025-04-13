package main

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
)

var categories = []string{"Зростаючий", "Строковий", "Інше"}
var names = []string{"Вася", "Петя", "Коля", "Дима", "Саша"}

func main() {
	file, err := os.Create("import.csv")
	if err != nil {
		log.Fatalf("Не вдалося створити файл: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"rate", "name", "type_deposit"})

	for i := 0; i < 10000; i++ {
		amount := rand.Intn(30)
		note := "Примітка " + strconv.Itoa(i)
		name := names[rand.Intn(len(names))]
		category := categories[rand.Intn(len(categories))]

		writer.Write([]string{
			strconv.Itoa(amount),
			note,
			name,
			category,
		})
	}
	log.Println("CSV файл створено.")
}
