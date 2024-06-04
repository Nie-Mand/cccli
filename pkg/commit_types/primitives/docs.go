package committypes_primitives

import "github.com/Nie-Mand/cccli/internal/core/domain"

func NewDocsCommit() domain.CommitType {
	return domain.CommitType{
		Id:          "docs",
		Name:        "Documentation",
		Description: "Documentation only changes",
	}
}
