package in_memory_repository

import (
	"app/internal/domain/entities"
	"app/internal/infra/repositories"
)

type OfferingRepository struct {
	offerings []*entities.Offering
}

var _ repositories.IOfferingRepository = (*OfferingRepository)(nil)

func NewOfferingRepository() *OfferingRepository {
	of := &entities.Offering{
		ID:    "offering-1",
		Name:  "Simple Cut",
		Price: 40 * 100,
	}

	return &OfferingRepository{
		offerings: []*entities.Offering{of},
	}
}

func (o *OfferingRepository) Save(offering *entities.Offering) {
	o.offerings = append(o.offerings, offering)
}

func (o *OfferingRepository) FindById(id string) *entities.Offering {
	for _, offering := range o.offerings {
		if offering.ID == id {
			return offering
		}
	}
	return nil
}
