package committypes_primitives

import "github.com/Nie-Mand/cccli/internal/core/domain"

func NewRevertCommit() domain.CommitType {
	return domain.CommitType{
		Id:          "revert",
		Name:        "Reverts",
		Description: "Reverts a previous commit",
	}
}
