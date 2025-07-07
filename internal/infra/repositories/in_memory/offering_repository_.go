package in_memory_repository

import (
	"app/internal/domain/entities"
	"app/internal/infra/repositories"
	"time"
)

type OfferingRepository struct {
	offering []*entities.Offering
}

var _ repositories.IOfferingRepository = (*OfferingRepository)(nil)

func NewOfferingRepository() *OfferingRepository {
	of := &entities.Offering{
		ID:       "offering-1",
		Name:     "Simple Cut",
		Price:    40 * 100,
		Duration: 30 * time.Minute,
	}

	return &OfferingRepository{
		offering: []*entities.Offering{of},
	}
}

func (or *OfferingRepository) Save(offering *entities.Offering) {
	or.offering = append(or.offering, offering)
}

func (or *OfferingRepository) FindById(id string) *entities.Offering {
	for _, offering := range or.offering {
		if offering.ID == id {
			return offering
		}
	}
	return nil
}
