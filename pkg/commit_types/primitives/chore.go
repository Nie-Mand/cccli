package committypes_primitives

import "github.com/Nie-Mand/cccli/internal/core/domain"

func NewChoreCommit() domain.CommitType {
	return domain.CommitType{
		Id:          "chore",
		Name:        "Chores",
		Description: "Other changes that don't modify src or test files",
	}
}
