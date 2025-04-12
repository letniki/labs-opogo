package internal

import (
	"context"
	"fmt"
)

// Service визначає інтерфейс бізнес-логіки
type Service struct {
	db Storage
}

func NewService(db Storage) Service {
	return Service{db: db}
}

// Створення депозиту

func (s Service) CreateDeposit(ctx context.Context, deposit DepositType) (int, error) {
	id, err := s.db.CreateDeposit(ctx, deposit)
	if err != nil {
		return 0, fmt.Errorf("не вдалося створити депозит: %w", err)
	}
	return id, nil
}

// Отримання депозиту за ID

func (s Service) GetDeposit(ctx context.Context, id int) (DepositType, error) {
	deposit, err := s.db.GetDeposit(ctx, id)
	if err != nil {
		return DepositType{}, fmt.Errorf("не вдалося отримати депозит: %w", err)
	}
	return deposit, nil
}

// Створення клієнту

func (s Service) CreatePerson(ctx context.Context, person Person) (int, error) {
	id, err := s.db.CreatePerson(ctx, person)
	if err != nil {
		return 0, fmt.Errorf("не вдалося створити клієнта: %w", err)
	}
	return id, nil
}

// Отримання товару за ID

func (s Service) GetPerson(ctx context.Context, id int) (Person, error) {
	person, err := s.db.GetPerson(ctx, id)
	if err != nil {
		return Person{}, fmt.Errorf("не вдалося отримати товар: %w", err)
	}
	return person, nil
}
