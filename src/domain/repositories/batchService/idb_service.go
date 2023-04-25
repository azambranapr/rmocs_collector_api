package batchService

import (
	"errors"
	"github.com/azambranapr/rmocs_collector_api/src/data/models/batch"
	"github.com/azambranapr/rmocs_collector_api/src/data/models/total"
)

var (
	errNotFound  = errors.New("data not found")
	errEmpty     = errors.New("data table is empty")
	errNoChanges = errors.New("there are no changes or data not found")
)

type IdbService interface {
	Create(*batch.Model) error
	Update(*batch.Model, string) (int64, error)
	GetAll() (batch.Models, error)
	GetById(int) (batch.Models, error)
	GetByClaimNum(string) (batch.Models, error)
	Delete(string) (int64, error)
	DeleteAll() (int64, error)
	GetTotal() (total.Models, error)
}
