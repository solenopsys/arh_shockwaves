package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
	"strings"
	"xs/internal/jobs"
)

type model struct {
	jobs     []jobs.PrintableJob
	index    int
	width    int
	height   int
	spinner  spinner.Model
	progress progress.Model
	done     bool
}

var (
	currentPkgNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	doneStyle           = lipgloss.NewStyle().Margin(1, 2)
	checkMark           = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
	errorMark           = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).SetString("x")
)

func newModel(jobs []jobs.PrintableJob) model {
	p := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(20),
		progress.WithoutPercentage(),
	)
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	return model{
		jobs:     jobs,
		spinner:  s,
		progress: p,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(runJob(m.jobs[m.index]), m.spinner.Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		}
	case jobExecMessage:
		return m.funcName(checkMark)
	case jobErrorMessage:
		return m.funcName(errorMark)
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case progress.FrameMsg:
		newModel, cmd := m.progress.Update(msg)
		if newModel, ok := newModel.(progress.Model); ok {
			m.progress = newModel
		}
		return m, cmd
	}
	return m, nil
}

func (m model) funcName(resType lipgloss.Style) (tea.Model, tea.Cmd) {
	if m.index >= len(m.jobs)-1 {
		m.done = true

		progressNext := m.progress.SetPercent(1)
		title := m.jobs[m.index].Title()
		return m, tea.Batch(progressNext, tea.Printf("%s %s - %s", resType, title.Name, title.Description))
	}

	// Update progress bar
	progressCmd := m.progress.SetPercent(float64(m.index) / float64(len(m.jobs)-1))
	title := m.jobs[m.index].Title()

	m.index++

	job := m.jobs[m.index]

	commands := []tea.Cmd{
		progressCmd,
		tea.Printf("%s %s - %s ", resType, title.Name, title.Description),
		runJob(job),
	}

	return m, tea.Batch(
		commands...,
	)
}

func (m model) View() string {
	n := len(m.jobs)
	w := lipgloss.Width(fmt.Sprintf("%d", n))

	if m.done {
		return doneStyle.Render(fmt.Sprintf("Done:  %d jobs! (Press [q] for exit)", n))
	}

	pkgCount := fmt.Sprintf(" %*d/%*d", w, m.index, w, n-1)

	spin := m.spinner.View() + " "
	prog := m.progress.View()
	cellsAvail := max(0, m.width-lipgloss.Width(spin+prog+pkgCount))

	job := m.jobs[m.index]
	pkgName := currentPkgNameStyle.Render(job.Title().Name)
	info := lipgloss.NewStyle().MaxWidth(cellsAvail).Render("Run " + pkgName + " (" + job.Title().Description + ")")

	cellsRemaining := max(0, m.width-lipgloss.Width(spin+info+prog+pkgCount))
	gap := strings.Repeat(" ", cellsRemaining)

	return spin + info + gap + prog + pkgCount
}

type exitMessage string
type jobExecMessage string
type jobErrorMessage string

func runJob(job jobs.PrintableJob) tea.Cmd {

	return func() tea.Msg {
		result := job.Execute()

		if result.Success {
			return jobExecMessage(job.Title().Name)
		} else {
			return jobErrorMessage(result.Error.Error())
		}

	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func ProcessingJobs(jobs []jobs.PrintableJob) {
	if _, err := tea.NewProgram(newModel(jobs)).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
