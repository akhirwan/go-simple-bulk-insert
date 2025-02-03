package feature

import (
	"context"
	"go-simple-bulk-insert/domain/counter/model"
	"go-simple-bulk-insert/domain/counter/repository"
	"log"
	"net/http"
)

type CounterFeature interface {
	CreateFeature(ctx context.Context, req *model.CreateRequest) (code int, message string)
}

type counterFeature struct {
	CounterRepo repository.CounterRepository
}

func NewCounterFeature(counterRepo repository.CounterRepository) CounterFeature {
	return &counterFeature{
		CounterRepo: counterRepo,
	}
}

func (f counterFeature) CreateFeature(ctx context.Context, req *model.CreateRequest) (code int, message string) {
	log.Println("getting number counter...")
	counter, err := f.CounterRepo.GetCounter(req)
	if err != nil {
		log.Println("invalid type :", err.Error())

		return http.StatusBadRequest, "invalid type"
	}

	startNumber := counter.StartNumber
	lastNumber := counter.LastNumber
	currentNumber := startNumber
	if lastNumber != nil {
		currentNumber = *lastNumber + 1
	}

	var transactions []*model.TestNumberTransaction

	log.Printf("inserting %v numbers...", req.Total)
	for i := 0; i < req.Total; i++ {
		transaction := model.TestNumberTransaction{
			Number: currentNumber + i,
			Action: req.Action,
		}

		transactions = append(transactions, &transaction)
	}

	if err = f.CounterRepo.BulkInsertTransaction(transactions); err != nil {
		log.Println("failed to insert transactions :", err.Error())

		return http.StatusBadGateway, "failed to insert transactions"
	}
	log.Printf("inserted %v numbers", req.Total)

	newLastNumber := currentNumber + req.Total - 1
	log.Printf("updating last number to %v", newLastNumber)
	if err = f.CounterRepo.UpdateCounter(&newLastNumber, &counter); err != nil {
		log.Println("failed to update counter :", err.Error())

		return http.StatusBadGateway, "failed to update counter"
	}
	log.Printf("updated last number to %v", newLastNumber)

	f.CounterRepo.CommitTX()

	return http.StatusOK, "Data inserted"
}
