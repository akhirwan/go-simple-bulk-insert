package repository

import (
	"fmt"
	"go-simple-bulk-insert/domain/counter/model"
	"sync"

	"gorm.io/gorm"
)

type CounterRepository interface {
	CommitTX()
	BulkInsertTransaction(transaction []*model.TestNumberTransaction) (err error)
	GetCounter(req *model.CreateRequest) (counter model.TestNumberCounter, err error)
	UpdateCounter(newLastNumber *int, counter *model.TestNumberCounter) (err error)
}

type counterRepository struct {
	DB, TX *gorm.DB
}

func NewCounterRepository(db *gorm.DB) CounterRepository {
	return &counterRepository{
		DB: db,
		TX: db.Begin(),
	}
}

func (r *counterRepository) CommitTX() {
	r.TX.Commit()
}

func (r *counterRepository) BulkInsertTransaction(transactions []*model.TestNumberTransaction) (err error) {

	batchSize := 100 // Smaller batch is faster
	batchID := 0
	total := len(transactions) // align with maxconns of database
	ch := make(chan error, total)
	var wg sync.WaitGroup

	for i := 0; i < total; i += batchSize {
		j := i + batchSize
		if j > total {
			j = total
		}

		batchID++

		wg.Add(1)
		// log.Printf("Batch insert %v to %v of %v\n", i+1, j, total)
		go r.batchInsertTransaction(batchID, batchSize, transactions[i:j], ch, &wg)
	}

	wg.Wait()
	close(ch)

	// read Values from the channel
	for err = range ch {
		if err != nil {
			r.TX.Rollback()
			err = fmt.Errorf("[FATAL] S/4 Hana Sales insert query error: %s", err.Error()) // return from func not for loop
			return
		}
	}

	return
}

func (r *counterRepository) batchInsertTransaction(id, size int, transactions []*model.TestNumberTransaction, ch chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	var err error

	if err = r.TX.CreateInBatches(&transactions, size).Error; err != nil {
		err = fmt.Errorf("[%d] %v", id, err)
	}

	ch <- err
}

func (r *counterRepository) GetCounter(req *model.CreateRequest) (counter model.TestNumberCounter, err error) {
	if err = r.TX.Where("type = ?", req.Type).First(&counter).Error; err != nil {
		r.TX.Rollback()
		return
	}

	return
}

func (r *counterRepository) UpdateCounter(newLastNumber *int, counter *model.TestNumberCounter) (err error) {
	if err = r.TX.Model(&counter).Update("last_number", newLastNumber).Error; err != nil {
		r.TX.Rollback()
		return
	}

	return
}
