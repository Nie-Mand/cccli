package domain

import (
	"strings"
)

type CommitType struct {
	Id string

	Name string

	Description string
}

func (c *CommitType) Matches(q string) bool {
	q = strings.ToLower(q)

	if strings.Contains(strings.ToLower(c.Id), q) {
		return true
	}

	if strings.Contains(strings.ToLower(c.Name), q) {
		return true
	}

	if strings.Contains(strings.ToLower(c.Description), q) {
		return true
	}

	return false
}
