package main

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
)

var depositTypes = []string{"Зростаючий", "Строковий", "Інше"}
var names = []string{"Вася", "Петя", "Коля", "Дима", "Саша"}

func main() {
	file, err := os.Create("import.csv")
	if err != nil {
		log.Fatalf("Не вдалося створити файл: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"rate", "note", "name", "type_deposit"})

	for i := 0; i < 10000; i++ {
		rate := rand.Intn(30)
		note := "Примітка " + strconv.Itoa(i)
		name := names[rand.Intn(len(names))]
		depositTypes := depositTypes[rand.Intn(len(depositTypes))]

		writer.Write([]string{
			strconv.Itoa(rate),
			note,
			name,
			depositTypes,
		})
	}
	log.Println("CSV файл створено.")
}
