package tui

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/MachadoMichael/shortcut/terminal"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	appStyle   = lipgloss.NewStyle().Padding(1, 2)
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
)

type item struct {
	title   string
	command string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.command }
func (i item) FilterValue() string { return i.title }

type listKeyMap struct {
	toggleSpinner    key.Binding
	toggleTitleBar   key.Binding
	toggleStatusBar  key.Binding
	togglePagination key.Binding
	toggleHelpMenu   key.Binding
	insertItem       key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		insertItem: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "add item"),
		),
		toggleSpinner: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "toggle spinner"),
		),
		toggleTitleBar: key.NewBinding(
			key.WithKeys("T"),
			key.WithHelp("T", "toggle title"),
		),
		toggleStatusBar: key.NewBinding(
			key.WithKeys("S"),
			key.WithHelp("S", "toggle status"),
		),
		togglePagination: key.NewBinding(
			key.WithKeys("P"),
			key.WithHelp("P", "toggle pagination"),
		),
		toggleHelpMenu: key.NewBinding(
			key.WithKeys("H"),
			key.WithHelp("H", "toggle help"),
		),
	}
}

type model struct {
	list          list.Model
	itemGenerator *ItemGenerator
	keys          *listKeyMap
	delegateKeys  *delegateKeyMap
}

func newModel(dic map[string]string) model {
	var (
		itemGenerator ItemGenerator
		delegateKeys  = newDelegateKeyMap()
		listKeys      = newListKeyMap()
	)

	// Make initial list of items
	items := itemGenerator.read(dic)
	// Setup list
	delegate := newItemDelegate(delegateKeys)
	shortcutsList := list.New(items, delegate, 0, 0)
	shortcutsList.Title = "Shortcuts"
	shortcutsList.Styles.Title = titleStyle
	shortcutsList.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.toggleSpinner,
			listKeys.insertItem,
			listKeys.toggleTitleBar,
			listKeys.toggleStatusBar,
			listKeys.togglePagination,
			listKeys.toggleHelpMenu,
		}
	}

	return model{
		list:          shortcutsList,
		keys:          listKeys,
		delegateKeys:  delegateKeys,
		itemGenerator: &itemGenerator,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		// Handle all key bindings here
		if m.list.FilterState() == list.Filtering {
			break
		}

	case executeCommandMsg:
		// Quit the TUI and execute the command after quitting
		return m, func() tea.Msg {
			// Ensure we quit the TUI before executing the command
			tea.Quit()

			// Execute the command outside the TUI environment
			output, err := terminal.ExecuteInteractive(msg.command)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error executing command:", err)
				return nil
			}

			// Print the output to the terminal
			fmt.Println(output)

			return tea.Quit()
		}

	}

	// Default list update
	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	var cmds []tea.Cmd
//
// 	switch msg := msg.(type) {
// 	case tea.WindowSizeMsg:
// 		h, v := appStyle.GetFrameSize()
// 		m.list.SetSize(msg.Width-h, msg.Height-v)
//
// 	case tea.KeyMsg:
// 		// Handle all key bindings here
// 		if m.list.FilterState() == list.Filtering {
// 			break
// 		}
//
// 		// Handle various toggle commands
// 		// ...
//
// 	case executeCommandMsg:
// 		// Quit the TUI and execute the command
// 		go func() {
// 			output, err := terminal.Execute(msg.command)
// 			if err != nil {
// 				fmt.Fprintln(os.Stderr, "Error:", err)
// 				return
// 			}
// 			fmt.Println(output)
// 		}()
// 		return m, tea.Quit // Close the TUI
//
// 	}
//
// 	// Default list update
// 	newListModel, cmd := m.list.Update(msg)
// 	m.list = newListModel
// 	cmds = append(cmds, cmd)
//
// 	return m, tea.Batch(cmds...)
// }

func (m model) View() string {
	return appStyle.Render(m.list.View())
}

func Init(dic map[string]string) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	if _, err := tea.NewProgram(newModel(dic), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
