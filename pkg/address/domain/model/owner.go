package model

import (
	"time"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type Owner struct {
	domain.ValueObject

	ID        OwnerID
	DeletedAt *time.Time
}

func (aa *Owner) IsAvailable() bool {
	return aa.DeletedAt == nil
}

func (aa *Owner) Equals(other *Owner) bool {
	return aa.SameValueAs(other)
}

func (aa *Owner) SameValueAs(other *Owner) bool {
	return aa.ID.ID() == other.ID.ID() && aa.DeletedAt == other.DeletedAt
}

func NewOwner(id OwnerID, deletedAt *time.Time) *Owner {
	return &Owner{ID: id, DeletedAt: deletedAt}
}
