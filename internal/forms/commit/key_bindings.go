package commit

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
)

func (f *CommitForm) keyBindings() *huh.KeyMap {
	km := huh.NewDefaultKeyMap()

	// Custom key bindings
	km.Input.Prev = key.NewBinding(key.WithKeys("left", "esc"), key.WithDisabled())
	km.Input.Next = key.NewBinding(key.WithKeys("right", "enter"), key.WithDisabled())

	km.Text.Prev = key.NewBinding(key.WithKeys("left", "esc"), key.WithDisabled())
	km.Text.Next = key.NewBinding(key.WithKeys("right", "enter"), key.WithDisabled())

	km.Select.Prev = key.NewBinding(key.WithKeys("left", "esc"), key.WithDisabled())
	km.Select.Next = key.NewBinding(key.WithKeys("right", "enter"), key.WithDisabled())

	km.MultiSelect.Prev = key.NewBinding(key.WithKeys("left", "esc"), key.WithDisabled())
	km.MultiSelect.Next = key.NewBinding(key.WithKeys("right", "enter"), key.WithDisabled())

	return km
}
