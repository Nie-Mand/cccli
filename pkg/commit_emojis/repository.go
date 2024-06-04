package commitemojis

import (
	"encoding/json"
	"os"

	"github.com/Nie-Mand/cccli/internal/core/domain"
)

type CommitEmojiRepository struct {
	list []domain.CommitEmoji
}

func NewCommitEmojiRepository() *CommitEmojiRepository {
	emojis := &CommitEmojiRepository{}

	err := emojis.load()

	if err != nil {
		emojis.list = []domain.CommitEmoji{}
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
	filePath := "pkg/commit_emojis/emojis.json"

	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	loadedEmojis := []EmojiModel{}

	err = json.NewDecoder(file).Decode(&loadedEmojis)

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
