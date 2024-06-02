package root_model

func (m RootModel) View() string {
	s := "\nPress q to quit.\n"

	// Show the previous phases data
	for i, phase := range m.phases {
		if i == m.cursor {
			break
		}

		s += phase.Summary()
	}

	s += m.phases[m.cursor].View()

	return s
}
