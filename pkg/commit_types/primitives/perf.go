package committypes_primitives

import "github.com/Nie-Mand/cccli/internal/core/domain"

func NewPerfCommit() domain.CommitType {
	return domain.CommitType{
		Id:          "perf",
		Name:        "Performance Improvements",
		Description: "A code change that improves performance",
	}
}
