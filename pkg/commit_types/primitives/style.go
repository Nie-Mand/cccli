package committypes_primitives

import "github.com/Nie-Mand/cccli/internal/core/domain"

func NewStyleCommit() domain.CommitType {
	return domain.CommitType{
		Id:          "style",
		Name:        "Styles",
		Description: "Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)",
	}
}
