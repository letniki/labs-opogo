package main

import "fmt"

type DepositType struct {
	Name string
	Rate float64
}

type Client struct {
	Name     string
	Deposits map[DepositType]float64
}

func (c *Client) CalculateIncome() float64 {
	totalIncome := 0.0
	for deposit, amount := range c.Deposits {
		income := amount * deposit.Rate / 100
		totalIncome += income
	}
	return totalIncome
}

func FindClientsWithLargeDeposits(clients []Client, biggerThan float64) []Client {
	var result []Client
	for _, client := range clients {
		for _, amount := range client.Deposits {
			if amount > biggerThan {
				result = append(result, client)
				break
			}
		}
	}
	return result
}

func main() {

	deposit1 := DepositType{"Депозит 1", 5.0}
	deposit2 := DepositType{"Депозит 2", 7.0}

	client1 := Client{
		Name: "Emily Johnson",
		Deposits: map[DepositType]float64{
			deposit1: 20000,
			deposit2: 60000,
		},
	}
	client2 := Client{
		Name: "Michael Williams",
		Deposits: map[DepositType]float64{
			deposit1: 30000,
		},
	}
	client3 := Client{
		Name: "James Davis",
		Deposits: map[DepositType]float64{
			deposit2: 80000,
		},
	}

	clients := []Client{client1, client2, client3}

	clientsWithLargeDeposits := FindClientsWithLargeDeposits(clients, 50000)

	fmt.Println("Клієнти з депозитами понад 50 000:")
	for _, client := range clientsWithLargeDeposits {
		fmt.Println(client.Name)
	}

	for _, client := range clients {
		income := client.CalculateIncome()
		fmt.Printf("Загальний дохід клієнта %s: %.2f\n", client.Name, income)
	}
}
