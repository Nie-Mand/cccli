package committypes_primitives

import "github.com/Nie-Mand/cccli/internal/core/domain"

func NewFeatureCommit() domain.CommitType {
	return domain.CommitType{
		Id:          "feat",
		Name:        "Features",
		Description: "A new feature",
	}
}
