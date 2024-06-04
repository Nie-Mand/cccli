package changedfiles

import (
	"os/exec"
	"strings"

	"github.com/charmbracelet/log"
)

type ChangedFilesRepository struct {
	list []string
}

func NewChangedFilesRepository() *ChangedFilesRepository {
	files := &ChangedFilesRepository{}

	err := files.load()

	if err != nil {
		log.Error(err)
		return nil
	}

	return files
}

func (r *ChangedFilesRepository) List() []string {
	return r.list
}

func (r *ChangedFilesRepository) load() error {
	cmd := exec.Command("git", "status", "--porcelain")
	out, err := cmd.Output()

	if err != nil {
		return err
	}

	files := []string{}

	parsed := strings.Split(string(out), "\n")
	for _, line := range parsed {
		if line == "" {
			continue
		}

		files = append(files, strings.TrimSpace(line[3:]))
	}

	r.list = files

	return nil
}
