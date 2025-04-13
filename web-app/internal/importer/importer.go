package importer

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"lab2/models"
	"lab2/repository"
)

func ImportCSV(filename string, repo repository.TransactionRepository) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Не вдалося відкрити CSV файл: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Помилка читання CSV: %v", err)
	}

	successCount := 0
	failCount := 0

	for _, record := range records {
		if len(record) < 4 {
			failCount++
			continue
		}

		amount, err := strconv.Atoi(record[0])
		if err != nil {
			failCount++
			continue
		}

		transaction := models.Transaction{
			Amount: amount,
			Note:   record[1],
			Type: models.Type{
				Name:     record[2],
				Category: record[3],
			},
		}

		if err := repo.Save(transaction); err != nil {
			failCount++
			continue
		}
		successCount++
	}

	log.Printf("Імпорт завершено. Успішно: %d, з помилками: %d", successCount, failCount)
}
