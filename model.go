package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const divisor, buffer = 4, 4

type Model struct {
	focused  Status
	lists    []list.Model
	loading  spinner.Model
	err      error
	loaded   bool
	quitting bool
}

func (m *Model) initSpinner() {
	m.loading = spinner.Model{}
	m.loading.Spinner = spinner.Dot
}

func (m *Model) initLists(width, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width/divisor, height-buffer)
	defaultList.SetShowHelp(false)
	m.lists = []list.Model{defaultList, defaultList, defaultList}

	m.lists[todo].Title = "To Do"
	m.lists[todo].SetItems([]list.Item{
		Task{status: todo, title: "buy milk", description: "strawberry milk"},
		Task{status: todo, title: "eat sushi", description: "negitoro roll"},
		Task{status: todo, title: "fold laundry", description: "or wear wrinkly shirts"},
	})

	m.lists[inProgress].Title = "In Progress"
	m.lists[inProgress].SetItems([]list.Item{
		Task{status: inProgress, title: "stay", description: "cool"},
	})

	m.lists[done].Title = "Done"
	m.lists[done].SetItems([]list.Item{
		Task{status: done, title: "dusted", description: "*clap hands*"},
	})
}

func (m Model) Init() tea.Cmd {
	return m.loading.Tick
}

func (m *Model) Next() {
	m.focused++
	m.focused = m.focused.Wrap()
}

func (m *Model) Previous() {
	m.focused--
	m.focused = m.focused.Wrap()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		UpdateStyleWidths(msg.Width/divisor, msg.Height-buffer)
		if !m.loaded {
			m.initLists(msg.Width, msg.Height)
			m.loaded = true
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "left", "h":
			m.Previous()
		case "right", "l":
			m.Next()
		}
	default:
		m.initSpinner()
	}
	var cmd tea.Cmd
	if !m.loaded {
		m.loading, cmd = m.loading.Update(msg)
	} else {
		m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	}
	return m, cmd
}

func (m Model) RenderList(e Status, s string) string {
	switch {
	case m.focused == e:
		return FocusedStyle.Render(s)
	default:
		return ColumnStyle.Render(s)
	}
}

func (m Model) View() string {
	switch {
	case m.quitting:
		return ""
	case !m.loaded:
		return fmt.Sprintf("\n\n   %s Loading...", m.loading.View())
	default:
		todoView := m.lists[todo].View()
		inProgView := m.lists[inProgress].View()
		doneView := m.lists[done].View()
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			m.RenderList(todo, todoView),
			m.RenderList(inProgress, inProgView),
			m.RenderList(done, doneView),
		)
	}
}
