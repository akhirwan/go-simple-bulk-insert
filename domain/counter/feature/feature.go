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
	tx := f.CounterRepo.BeginTransaction(ctx)

	log.Println("getting number counter...")
	counter, err := f.CounterRepo.GetCounter(tx, req)
	if err != nil {
		tx.Rollback()
		log.Println("invalid type :", err.Error())

		return http.StatusBadRequest, "invalid type"
	}

	startNumber := counter.StartNumber
	lastNumber := counter.LastNumber
	currentNumber := startNumber
	if lastNumber != nil {
		currentNumber = *lastNumber + 1
	}

	log.Printf("inserting %v numbers...", req.Total)
	for i := 0; i < req.Total; i++ {
		transaction := model.TestNumberTransaction{
			Number: currentNumber + i,
			Action: req.Action,
		}

		if err = f.CounterRepo.BulkInsertTransaction(tx, transaction); err != nil {
			tx.Rollback()
			log.Println("failed to insert transactions :", err.Error())

			return http.StatusBadGateway, "failed to insert transactions"
		}
	}
	log.Printf("inserted %v numbers", req.Total)

	newLastNumber := currentNumber + req.Total - 1
	log.Printf("updating last number to %v", newLastNumber)
	if err = f.CounterRepo.UpdateCounter(tx, &newLastNumber, &counter); err != nil {
		tx.Rollback()
		log.Println("failed to update counter :", err.Error())

		return http.StatusBadGateway, "failed to update counter"
	}
	log.Printf("updated last number to %v", newLastNumber)

	tx.Commit()

	return http.StatusOK, "Data inserted"
}
