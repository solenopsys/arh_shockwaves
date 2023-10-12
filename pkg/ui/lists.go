package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"time"
	"xs/internal/configs"
	"xs/pkg/io"

	"xs/internal/jobs"
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
	List   list.Model
	Filter list.FilterFunc
	Value  string // todo remove and use callback
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
			m.Value = m.List.FilterInput.Value()

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

func JobsToListModel(jobsPlan []jobs.PrintableJob, title string, filter string) (*tea.Program, *Model) {
	items := []list.Item{}

	for _, job := range jobsPlan {
		items = append(items, item{
			title: job.Description().Short,
			desc:  job.Description().Description,
		})
	}

	keyMap := list.DefaultKeyMap()

	m2 := list.New(items, list.NewDefaultDelegate(), 0, 0)
	m2.Filter = ListFilter
	m2.FilterInput.SetValue(filter)
	//	m2.SetShowFilter(true)
	m2.KeyMap = keyMap
	m := Model{List: m2}
	m.List.Title = title

	p := tea.NewProgram(&m)

	if filter != "" {

		go func() {
			p.Send(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("/")})
			time.Sleep(10 * time.Millisecond)
			p.Send(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(" ")})

			time.Sleep(10 * time.Millisecond)
			p.Send(tea.KeyMsg{Type: tea.KeyBackspace})
			time.Sleep(300 * time.Millisecond)
		}()
	}
	return p, &m
}
