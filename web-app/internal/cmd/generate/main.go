package generate

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
)

var categories = []string{"Їжа", "Оренда", "Канцелярія", "Інше"}
var names = []string{"Кава", "Хліб", "Папір", "Оренда офісу", "Шафа", "Мило", "Ручка"}

func main() {
	file, err := os.Create("import.csv")
	if err != nil {
		log.Fatalf("Не вдалося створити файл: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"amount", "note", "type_name", "type_category"})

	for i := 0; i < 10000; i++ {
		amount := rand.Intn(2000) - 1000
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
