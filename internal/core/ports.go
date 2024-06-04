package core

import "github.com/Nie-Mand/cccli/internal/core/domain"

type CommitTypeGateway interface {
	List() []domain.CommitType
	FindById(id string) domain.CommitType
}
