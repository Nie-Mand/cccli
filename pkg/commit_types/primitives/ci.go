package committypes_primitives

import "github.com/Nie-Mand/cccli/internal/core/domain"

func NewCICommit() domain.CommitType {
	return domain.CommitType{
		Id:          "ci",
		Name:        "Continuous Integration",
		Description: "Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs)",
	}
}
