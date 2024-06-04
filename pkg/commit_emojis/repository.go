package commitemojis

import (
	"encoding/json"

	"github.com/charmbracelet/log"

	"github.com/Nie-Mand/cccli/internal/core/domain"
)

type CommitEmojiRepository struct {
	list []domain.CommitEmoji
}

func NewCommitEmojiRepository() *CommitEmojiRepository {
	emojis := &CommitEmojiRepository{}

	err := emojis.load()

	if err != nil {
		log.Error(err)
		return nil
	}

	return emojis
}

func (r *CommitEmojiRepository) FindById(id string) domain.CommitEmoji {
	for _, commitType := range r.list {
		if commitType.Code == id {
			return commitType
		}
	}

	return domain.CommitEmoji{}
}

func (r *CommitEmojiRepository) List() []domain.CommitEmoji {
	return r.list
}

func (r *CommitEmojiRepository) load() error {
	file := emojisJson

	loadedEmojis := []EmojiModel{}

	err := json.Unmarshal([]byte(file), &loadedEmojis)

	if err != nil {
		panic(err)
	}

	for _, emoji := range loadedEmojis {
		r.list = append(r.list, domain.CommitEmoji{
			Emoji:       emoji.Emoji,
			Code:        emoji.Code,
			Description: emoji.Description,
			Name:        emoji.Name,
		})
	}

	return nil
}
