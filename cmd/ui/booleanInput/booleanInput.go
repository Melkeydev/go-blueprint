package booleanInput

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/melkeydev/go-blueprint/cmd/program"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
	titleStyle   = lipgloss.NewStyle().Background(lipgloss.Color("#01FAC6")).Foreground(lipgloss.Color("#030303")).Bold(true).Padding(0, 1, 0)
)

type Selection struct {
	Choice bool
}

// Update changes the value of a Selection's Choice
func (s *Selection) Update(value bool) {
	s.Choice = value
}

type model struct {
	choice *Selection
	header string
	exit   *bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func InitialBoolInput(selection *Selection, header string, program *program.Project) model {
	return model{
		choice: selection,
		header: titleStyle.Render(header),
		exit:   &program.Exit,
	}
}

// Update is called when "things happen", it checks for
// important keystrokes to signal when to quit, change selection,
// and confirm the selection.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			*m.exit = true
			return m, tea.Quit
		case "y":
			m.choice.Update(true)
			return m, tea.Quit
		case "n":
			m.choice.Update(false)
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := m.header + "\n\n"

	s += fmt.Sprintf("Press %s to accept. \nPress %s to skip.",
		focusedStyle.Render("y"), focusedStyle.Render(("n")))
	return s
}
