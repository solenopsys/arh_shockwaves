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
	"xs/pkg/io"
)

type model struct {
	executor *jobs.Executor

	width    int
	height   int
	spinner  spinner.Model
	progress progress.Model
}

var (
	currentPkgNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	doneStyle           = lipgloss.NewStyle().Margin(1, 2)
	checkMark           = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
	errorMark           = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).SetString("x")
)

func newModel(executor *jobs.Executor) model {
	p := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(20),
		progress.WithoutPercentage(),
	)
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	return model{
		executor: executor,
		spinner:  s,
		progress: p,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.runJob(), m.spinner.Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		case "l":
			return m.printLogs()
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

func (m model) printLogs() (tea.Model, tea.Cmd) {
	keys := m.executor.GetFailKeys()
	store := io.GetLogStore()

	for _, key := range keys {
		fmt.Printf("Log (%s) -------------------->	\n", key)
		log := store.GetLog(key)
		fmt.Printf(log)

	}

	return m, nil
}

func (m model) funcName(resType lipgloss.Style) (tea.Model, tea.Cmd) {
	job := m.executor.GetCurrent()
	if m.executor.IsDone() {

		progressNext := m.progress.SetPercent(1)
		title := job.Title()
		return m, tea.Batch(progressNext, tea.Printf("%s %s - %s", resType, title.Name, title.Description))
	}

	progressCmd := m.progress.SetPercent(m.executor.Progress())
	title := job.Title()

	m.executor.Next()

	commands := []tea.Cmd{
		progressCmd,
		tea.Printf("%s %s - %s ", resType, title.Name, title.Description),
		m.runJob(),
	}

	return m, tea.Batch(
		commands...,
	)
}

func (m model) View() string {
	n := m.executor.JobsCount()
	w := lipgloss.Width(fmt.Sprintf("%d", n))

	if m.executor.IsDone() {
		return doneStyle.Render(fmt.Sprintf("Executed:  %d jobs (success: %d errors: %d) \nPress [q] for exit or [l] for seen logs",
			n, m.executor.SuccessCount(), m.executor.ErrorCount()))
	}

	pkgCount := fmt.Sprintf(" %*d/%*d", w, m.executor.GetCurrentIndex(), w, n-1)

	spin := m.spinner.View() + " "
	prog := m.progress.View()
	cellsAvail := max(0, m.width-lipgloss.Width(spin+prog+pkgCount))

	job := m.executor.GetCurrent()

	pkgName := currentPkgNameStyle.Render(job.Title().Name)
	info := lipgloss.NewStyle().MaxWidth(cellsAvail).Render("Run " + pkgName + " (" + job.Title().Description + ")")

	cellsRemaining := max(0, m.width-lipgloss.Width(spin+info+prog+pkgCount))
	gap := strings.Repeat(" ", cellsRemaining)

	return spin + info + gap + prog + pkgCount
}

type exitMessage string
type jobExecMessage string
type jobErrorMessage string

func (m model) runJob() tea.Cmd {
	job := m.executor.GetCurrent()
	return func() tea.Msg {
		result := m.executor.RunJob()

		if result.Success {
			return jobExecMessage(job.Title().Name)
		} else {
			if result.Error == nil {
				return jobErrorMessage(result.Description)
			}
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

func ProcessingJobs(executor *jobs.Executor) {
	if _, err := tea.NewProgram(newModel(executor)).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
