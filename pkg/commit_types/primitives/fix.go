package committypes_primitives

import "github.com/Nie-Mand/cccli/internal/core/domain"

func NewFixCommit() domain.CommitType {
	return domain.CommitType{
		Id:          "fix",
		Name:        "Bug Fixes",
		Description: "A bug fix",
	}
}
