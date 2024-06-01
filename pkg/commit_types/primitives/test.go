package committypes_primitives

import "github.com/Nie-Mand/cccli/internal/core/domain"

func NewTestCommit() domain.CommitType {
	return domain.CommitType{
		Id:          "test",
		Name:        "Tests",
		Description: "Adding missing tests or correcting existing tests",
	}
}
