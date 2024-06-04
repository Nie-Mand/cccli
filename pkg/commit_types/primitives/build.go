package committypes_primitives

import "github.com/Nie-Mand/cccli/internal/core/domain"

func NewBuildCommit() domain.CommitType {
	return domain.CommitType{
		Id:          "build",
		Name:        "Builds",
		Description: "Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)",
	}
}
