package core

import "github.com/Nie-Mand/cccli/internal/core/domain"

type CommitTypeGateway interface {
	List() []domain.CommitType
	FindById(id string) domain.CommitType
}

type CommitEmojiGateway interface {
	List() []domain.CommitEmoji
	FindById(id string) domain.CommitEmoji
}

type ChangedFilesGateway interface {
	List() []string
}
