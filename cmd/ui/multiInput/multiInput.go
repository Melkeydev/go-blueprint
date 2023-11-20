// Package multiInput provides functions that
// help define and draw a multi-input step
package multiInput

import (
	"fmt"

	"github.com/melkeydev/go-blueprint/cmd/program"
	"github.com/melkeydev/go-blueprint/cmd/steps"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Change this
var (
	focusedStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
	titleStyle            = lipgloss.NewStyle().Background(lipgloss.Color("#01FAC6")).Foreground(lipgloss.Color("#030303")).Bold(true).Padding(0, 1, 0)
	selectedItemStyle     = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170")).Bold(true)
	selectedItemDescStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170"))
	descriptionStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#40BDA3"))
)

// A Selection represents a choice made in a multiInput step
type Selection struct {
	Choice string
}

// Update changes the value of a Selection's Choice
func (s *Selection) Update(value string) {
	s.Choice = value
}

// A multiInput.model contains the data for the multiInput step.
//
// It has the required methods that make it a bubbletea.Model
type model struct {
	cursor   int
	choices  []steps.Item
	selected map[int]struct{}
	choice   *Selection
	header   string
	exit     *bool
}

func (m model) Init() tea.Cmd {
	return nil
}

// InitialModelMulti initializes a multiInput step with
// the given data
func InitialModelMulti(choices []steps.Item, selection *Selection, header string, program *program.Project) model {
	return model{
		choices:  choices,
		selected: make(map[int]struct{}),
		choice:   selection,
		header:   titleStyle.Render(header),
		exit:     &program.Exit,
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
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			if len(m.selected) == 1 {
				m.selected = make(map[int]struct{})
			}
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "y":
			if len(m.selected) == 1 {
				for selectedKey := range m.selected {
					m.choice.Update(m.choices[selectedKey].Title)
					m.cursor = selectedKey
				}
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

// View is called to draw the multiInput step
func (m model) View() string {
	s := m.header + "\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = focusedStyle.Render(">")
			choice.Title = selectedItemStyle.Render(choice.Title)
			choice.Desc = selectedItemDescStyle.Render(choice.Desc)
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = focusedStyle.Render("x")
		}

		title := focusedStyle.Render(choice.Title)
		description := descriptionStyle.Render(choice.Desc)

		s += fmt.Sprintf("%s [%s] %s\n%s\n\n", cursor, checked, title, description)
	}

	s += fmt.Sprintf("Press %s to confirm choice.\n", focusedStyle.Render("y"))
	return s
}
