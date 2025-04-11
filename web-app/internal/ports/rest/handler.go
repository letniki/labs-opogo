package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/letniki/labs-opogo/internal"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	service internal.Service
}

func NewHandler(service internal.Service) Handler {
	return Handler{service: service}
}

// Створення депозиту

func (h Handler) CreateDeposit(w http.ResponseWriter, r *http.Request) {
	var deposit internal.DepositType

	if err := json.NewDecoder(r.Body).Decode(&deposit); err != nil {
		log.Printf("Помилка створення депозиту: %v", err)
		http.Error(w, "Невірний формат запиту", http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateDeposit(r.Context(), deposit)
	if err != nil {
		log.Printf("Помилка створення депозиту: %v", err)
		http.Error(w, "Помилка створення депозиту", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// Отримання депозиту за ID

func (h Handler) GetDeposit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Помилка створення депозиту: %v", err)
		http.Error(w, "Невірний формат ID", http.StatusBadRequest)
		return
	}

	deposit, err := h.service.GetDeposit(r.Context(), id)
	if err != nil {
		log.Printf("депозит відсутній: %v", err)
		http.Error(w, "Депозит не знайдено", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Printf("Помилка створення депозиту: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(deposit)
}

// Створення клієнту

func (h Handler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person internal.Person

	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		log.Printf("Помилка створення клієнту: %v", err)
		http.Error(w, "Невірний формат запиту", http.StatusBadRequest)
		return
	}

	id, err := h.service.CreatePerson(r.Context(), person)
	if err != nil {
		log.Printf("Помилка створення клієнту: %v", err)
		http.Error(w, "Помилка створення клієнту", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// Отримання клієнту за ID

func (h Handler) GetPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Помилка створення клієнту: %v", err)
		http.Error(w, "Невірний формат ID", http.StatusBadRequest)
		return
	}

	person, err := h.service.GetPerson(r.Context(), id)
	if err != nil {
		log.Printf("Помилка створення клієнту: %v", err)
		http.Error(w, "Клієнта не знайдено", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(person)
}
