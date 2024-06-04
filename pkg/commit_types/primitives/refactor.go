package committypes_primitives

import "github.com/Nie-Mand/cccli/internal/core/domain"

func NewRefactorCommit() domain.CommitType {
	return domain.CommitType{
		Id:          "refactor",
		Name:        "Code Refactoring",
		Description: "A code change that neither fixes a bug nor adds a feature",
	}
}
