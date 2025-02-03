package repository

import (
	"context"
	"go-simple-bulk-insert/domain/counter/model"

	"gorm.io/gorm"
)

type CounterRepository interface {
	BeginTransaction(ctx context.Context) *gorm.DB
	BulkInsertTransaction(tx *gorm.DB, transaction model.TestNumberTransaction) (err error)
	GetCounter(tx *gorm.DB, req *model.CreateRequest) (counter model.TestNumberCounter, err error)
	UpdateCounter(tx *gorm.DB, newLastNumber *int, counter *model.TestNumberCounter) (err error)
}

type counterRepository struct {
	DB *gorm.DB
}

func NewCounterRepository(db *gorm.DB) CounterRepository {
	return &counterRepository{
		DB: db,
	}
}

func (r *counterRepository) BeginTransaction(ctx context.Context) *gorm.DB {
	return r.DB.Begin()
}

func (r *counterRepository) BulkInsertTransaction(tx *gorm.DB, transaction model.TestNumberTransaction) (err error) {
	if err = tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return
	}

	return
}

func (r *counterRepository) GetCounter(tx *gorm.DB, req *model.CreateRequest) (counter model.TestNumberCounter, err error) {
	if err = tx.Where("type = ?", req.Type).First(&counter).Error; err != nil {
		return
	}

	return
}

func (r *counterRepository) UpdateCounter(tx *gorm.DB, newLastNumber *int, counter *model.TestNumberCounter) (err error) {
	if err = tx.Model(&counter).Update("last_number", newLastNumber).Error; err != nil {
		return
	}

	return
}
