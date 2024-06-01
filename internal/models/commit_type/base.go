package committype_model

import (
	"github.com/Nie-Mand/cccli/internal/core/domain"
	committypes "github.com/Nie-Mand/cccli/pkg/commit_types"
)

type CommitTypeModel struct {
	choices  []domain.CommitType
	cursor   int
	selected domain.CommitType
	query    string
	pageSize int
	done     bool
}

func InitCommitTypeModel() CommitTypeModel {
	choices := committypes.GetCommitTypes()

	return CommitTypeModel{
		choices:  choices,
		pageSize: 3,
	}
}
