package main

import "testing"

// TestСalculateIncome перевіряє, чи правильно обчислюється дохід клієнта.
func TestCalculateIncome(t *testing.T) {
	// Налаштовуємо депозити та клієнта
	deposit1 := DepositType{"Депозит 1", 5.0}
	deposit2 := DepositType{"Депозит 2", 7.0}
	client := Client{
		Name: "Тестовий Клієнт",
		Deposits: map[DepositType]float64{
			deposit1: 20000,
			deposit2: 60000,
		},
	}

	// Очікуваний загальний дохід: (20000 * 5%) + (60000 * 7%) = 1000 + 4200 = 5200
	expectedIncome := 5200.0

	actualIncome := client.CalculateIncome()

	if actualIncome != expectedIncome {
		t.Errorf("Очікуваний дохід %f, але отримано %f", expectedIncome, actualIncome)
	}
}

// TestFindClientsWithLargeDeposits перевіряє,
// чи правильно FindClientsWithLargeDeposits фільтрує клієнтів, чиї депозити перевищують задану суму.
func TestFindClientsWithLargeDeposits(t *testing.T) {
	clients := []Client{
		{
			Name: "Клієнт 1",
			Deposits: map[DepositType]float64{
				{"Депозит 1", 5.0}: 20000,
			},
		},
		{
			Name: "Клієнт 2",
			Deposits: map[DepositType]float64{
				{"Депозит 2", 7.0}: 60000,
			},
		},
		{
			Name: "Клієнт 3",
			Deposits: map[DepositType]float64{
				{"Депозит 3", 6.0}: 100000,
			},
		},
	}

	// Викликаємо функцію FindClientsWithLargeDeposits з порогом 50000
	result := FindClientsWithLargeDeposits(clients, 50000)

	// Очікуємо знайти Клієнта 2 та Клієнта 3
	expectedClientCount := 2

	if len(result) != expectedClientCount {
		t.Errorf("Очікувано знайти %d клієнтів, але знайдено %d", expectedClientCount, len(result))
	}

	// Перевіряємо імена клієнтів
	expectedNames := []string{"Клієнт 2", "Клієнт 3"}
	for i, client := range result {
		if client.Name != expectedNames[i] {
			t.Errorf("Очікуване ім'я клієнта %s, але отримано %s", expectedNames[i], client.Name)
		}
	}
}

// Перевіримо, чи функція правильно не повертає клієнтів, коли всі депозити менші за поріг.
func TestFindClientsWithLargeDepositsBelowThreshold(t *testing.T) {
	clients := []Client{
		{
			Name: "Клієнт 1",
			Deposits: map[DepositType]float64{
				{"Депозит 1", 5.0}: 10000,
			},
		},
		{
			Name: "Клієнт 2",
			Deposits: map[DepositType]float64{
				{"Депозит 2", 7.0}: 20000,
			},
		},
		{
			Name: "Клієнт 3",
			Deposits: map[DepositType]float64{
				{"Депозит 3", 6.0}: 30000,
			},
		},
	}

	// Викликаємо FindClientsWithLargeDeposits з порогом 50000
	result := FindClientsWithLargeDeposits(clients, 50000)

	// Очікуємо знайти 0 клієнтів, оскільки жоден депозит не перевищує 50000
	expectedClientCount := 0

	if len(result) != expectedClientCount {
		t.Errorf("Очікувано знайти %d клієнтів, але знайдено %d", expectedClientCount, len(result))
	}
}

// Перевіримо, чи правильно обчислюється дохід клієнта, якщо ставка на депозит 0.
func TestCalculateIncomeZeroRate(t *testing.T) {
	deposit1 := DepositType{"Депозит 1", 0.0} // Ставка 0%
	client := Client{
		Name: "Тестовий Клієнт",
		Deposits: map[DepositType]float64{
			deposit1: 100000,
		},
	}

	// Оскільки ставка 0%, дохід має бути 0
	expectedIncome := 0.0

	actualIncome := client.CalculateIncome()

	if actualIncome != expectedIncome {
		t.Errorf("Очікуваний дохід %f, але отримано %f", expectedIncome, actualIncome)
	}
}
