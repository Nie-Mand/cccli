package committype_model

import (
	"fmt"

	"github.com/Nie-Mand/cccli/internal/core/domain"
	committypes "github.com/Nie-Mand/cccli/pkg/commit_types"
)

type CommitTypeModel struct {
	choices  []domain.CommitType
	selected domain.CommitType
	cursor   int
	query    string
	pageSize int
}

func InitCommitTypeModel() CommitTypeModel {
	choices := committypes.GetCommitTypes()

	return CommitTypeModel{
		choices:  choices,
		pageSize: 3,
	}
}

func (m CommitTypeModel) Value() domain.CommitType {
	return m.selected
}

func (m CommitTypeModel) Summary() string {
	s := fmt.Sprintf("Select the type of change that you're commiting ~ %s\n", m.selected.Name)
	return s
}
