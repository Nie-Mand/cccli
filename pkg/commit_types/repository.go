package committypes

import (
	"github.com/Nie-Mand/cccli/internal/core/domain"
	committypes_primitives "github.com/Nie-Mand/cccli/pkg/commit_types/primitives"
)

type CommitTypeRepository struct {
	list []domain.CommitType
}

func NewCommitTypeRepository() *CommitTypeRepository {
	return &CommitTypeRepository{
		list: []domain.CommitType{
			committypes_primitives.NewFeatureCommit(),
			committypes_primitives.NewFixCommit(),
			committypes_primitives.NewDocsCommit(),
			committypes_primitives.NewStyleCommit(),
			committypes_primitives.NewRefactorCommit(),
			committypes_primitives.NewPerfCommit(),
			committypes_primitives.NewTestCommit(),
			committypes_primitives.NewBuildCommit(),
			committypes_primitives.NewCICommit(),
			committypes_primitives.NewChoreCommit(),
			committypes_primitives.NewRevertCommit(),
		},
	}
}

func (r *CommitTypeRepository) FindById(id string) domain.CommitType {
	for _, commitType := range r.list {
		if commitType.Id == id {
			return commitType
		}
	}

	return domain.CommitType{}
}

func (r *CommitTypeRepository) List() []domain.CommitType {
	return r.list
}
