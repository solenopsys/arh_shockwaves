package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
	"time"
	"xs/internal/configs"
	"xs/internal/jobs"
	"xs/pkg/io"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

//
//type listKeyMap struct {
//	apply key.Binding
//}

type Model struct {
	List        list.Model
	Applied     bool
	Filter      list.FilterFunc
	FilterValue string // todo remove and use callback
}

func (m Model) Init() tea.Cmd {
	return nil

	//	p.Send(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("/")})
	//	p.Send(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(" ")})
	//
	//	time.Sleep(1 * time.Millisecond)
	//	p.Send(tea.KeyMsg{Type: tea.KeyBackspace})
	//	time.Sleep(100 * time.Millisecond)
	//	p.Send(tea.KeyMsg{Type: tea.KeyEnter})
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		if msg.String() == "enter" && (m.List.FilterState() == list.FilterApplied || m.List.FilterState() == list.Unfiltered) {
			m.FilterValue = m.List.FilterInput.Value()
			m.Applied = true

			return m, tea.Quit

		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)

	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	return docStyle.Render(m.List.View())
}

//func newListKeyMap() list.KeyMap {
//	return &listKeyMap{
//		apply: key.NewBinding(
//			key.WithKeys("r"),
//			key.WithHelp("r", "Run selected"),
//		),
//	}
//}

func ListFilter(filter string, strings []string) []list.Rank {
	var ranks []list.Rank
	for i, str := range strings {

		matchedIndexes, err := configs.PatternMatchingRank(str, filter)
		if err != nil {
			io.Println("Error:", err)
			continue
		}

		if matchedIndexes != nil {
			ranks = append(ranks,

				list.Rank{
					MatchedIndexes: matchedIndexes,
					Index:          i,
				})
		}

	}
	return ranks
}

func ListView(items []jobs.ItemTitle, title string) bool {
	apply, _ := filteredListViewAbs(items, title, "", false)
	return apply
}

func filteredListViewAbs(listItems []jobs.ItemTitle, title string, filter string, filtred bool) (bool, string) {
	tea.EnterAltScreen()
	items := []list.Item{}

	for _, it := range listItems {
		items = append(items, item{
			title: it.Name,
			desc:  it.Description,
		})
	}

	keyMap := list.DefaultKeyMap()

	libsModel := list.New(items, list.NewDefaultDelegate(), 0, 0)
	libsModel.Filter = ListFilter
	libsModel.FilterInput.SetValue(filter)
	libsModel.SetShowFilter(filtred)
	libsModel.KeyMap = keyMap
	programModel := Model{List: libsModel, Applied: false, FilterValue: filter}
	programModel.List.Title = title

	program := tea.NewProgram(&programModel, tea.WithAltScreen())

	if filter != "" {

		go func() {
			program.Send(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("/")})
			time.Sleep(10 * time.Millisecond)
			program.Send(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(" ")})

			time.Sleep(10 * time.Millisecond)
			program.Send(tea.KeyMsg{Type: tea.KeyBackspace})
			time.Sleep(300 * time.Millisecond)
		}()
	}

	if _, err := program.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	tea.ExitAltScreen()

	return programModel.Applied, programModel.FilterValue
}

func FilteredListView(listItems []jobs.ItemTitle, title string, filter string) (bool, string) {
	return filteredListViewAbs(listItems, title, filter, true)
}
