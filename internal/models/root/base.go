package root_model

import (
	commitmessage_model "github.com/Nie-Mand/cccli/internal/models/commit_message"
	committype_model "github.com/Nie-Mand/cccli/internal/models/commit_type"
	"github.com/Nie-Mand/cccli/internal/utils"
)

type RootModel struct {
	cursor int
	phases []utils.Model
}

func InitRootModel() RootModel {
	return RootModel{
		cursor: 0,
		phases: []utils.Model{
			committype_model.InitCommitTypeModel(),
			commitmessage_model.InitCommitMessageModel(),
		},
	}
}
