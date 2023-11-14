// Package textinput provides functions that
// help define and draw a text-input step
package textinput

import (
	"errors"
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/melkeydev/go-blueprint/cmd/program"
)

var (
	titleStyle = lipgloss.NewStyle().Background(lipgloss.Color("#01FAC6")).Foreground(lipgloss.Color("#030303")).Bold(true).Padding(0, 1, 0)
	errorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF8700")).Bold(true).Padding(0, 0, 0)
)
type (
	errMsg error
)

// Output represents the text provided in a textinput step
type Output struct {
	Output string
}

// Output.update updates the value of the Output
func (o *Output) update(val string) {
	o.Output = val
}

// A textnput.model contains the data for the textinput step.
//
// It has the required methods that make it a bubbletea.Model
type model struct {
	textInput textinput.Model
	err       error
	output    *Output
	header    string
	exit      *bool
}

// InitialTextInputModel initializes a textinput step
// with the given data
func InitialTextInputModel(output *Output, header string, program *program.Project) model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,
		err:       nil,
		output:    output,
		header:    titleStyle.Render(header),
		exit:      &program.Exit,
	}
}

// CreateErrorInputModel creates a textinput step
// with the given error
func CreateErrorInputModel(err error) model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	exit := true

	return model{
		textInput: ti,
		err:       errors.New(errorStyle.Render(err.Error())),
		output:    nil,
		header:    "",
		exit:      &exit,
	}
}

// Init is called at the beginning of a textinput step
// and sets the cursor to blink
func (m model) Init() tea.Cmd {
	return textinput.Blink
}

// Update is called when "things happen", it checks for the users text input,
// and for Ctrl+C or Esc to close the program.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if len(m.textInput.Value()) > 1 {
				m.output.update(m.textInput.Value())
				return m, tea.Quit
			}
		case tea.KeyCtrlC, tea.KeyEsc:
			*m.exit = true
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		*m.exit = true
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

// View is called to draw the textinput step
func (m model) View() string {
	return fmt.Sprintf("%s\n\n%s\n\n",
		m.header,
		m.textInput.View(),
	)
}

func (m model) Err() string {
	return m.err.Error()
}
