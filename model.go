package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const divisor = 4

type Model struct {
	focused Status
	lists   []list.Model
	loading spinner.Model
	err     error
	loaded  bool
}

func (m *Model) initSpinner() {
	m.loading = spinner.Model{}
	m.loading.Spinner = spinner.Dot
}

func (m *Model) initLists(width, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width/divisor, height-divisor)
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

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.initLists(msg.Width, msg.Height)
			m.loaded = true
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

func (m Model) View() string {
	if !m.loaded {
		return fmt.Sprintf("\n\n   %s Loading...", m.loading.View())
	} else {
		todoView := m.lists[todo].View()
		inProgView := m.lists[inProgress].View()
		doneView := m.lists[done].View()
		switch m.focused {
		default:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				FocusedStyle.Render(todoView),
				ColumnStyle.Render(inProgView),
				ColumnStyle.Render(doneView),
			)
		}
	}
}
