package importer

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func GenerateCSV(filePath string, rows int) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Не вдалося створити файл: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"amount", "note", "type_name", "type_category"})

	rand.Seed(time.Now().UnixNano())

	types := []struct {
		Name     string
		Category string
	}{
		{"Salary", "Income"},
		{"Office Rent", "Expenses"},
		{"Stationery", "Expenses"},
		{"Bonus", "Income"},
		{"Software", "Expenses"},
	}

	for i := 0; i < rows; i++ {
		t := types[rand.Intn(len(types))]
		amount := rand.Float64()*10000 - 5000
		note := "Auto-generated record " + strconv.Itoa(i+1)

		writer.Write([]string{
			strconv.FormatFloat(amount, 'f', 2, 64),
			note,
			t.Name,
			t.Category,
		})
	}

	log.Printf("Згенеровано CSV файл: %s (%d рядків)\n", filePath, rows)
}
