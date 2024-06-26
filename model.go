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

func (m *Model) ProgressTask() tea.Msg {
	return m.moveTask(func(t Task) Task {
		return t.Next()
	})
}

func (m *Model) RegressTask() tea.Msg {
	return m.moveTask(func(t Task) Task {
		return t.Previous()
	})
}

func (m *Model) moveTask(f func(Task) Task) tea.Msg {
	selectedItem := m.lists[m.focused].SelectedItem()
	if selectedItem != nil {
		selectedTask := selectedItem.(Task)
		m.lists[selectedTask.status].RemoveItem(m.lists[m.focused].Index())
		selectedTask = f(selectedTask)
		fmt.Print(selectedTask.status)
		m.lists[selectedTask.status].
			InsertItem(
				len(m.lists[selectedTask.status].Items())-1,
				list.Item(selectedTask))
	}
	return nil
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
	m.lists[inProgress].Title = "In Progress"
	m.lists[done].Title = "Done"
}

func (m Model) Init() tea.Cmd {
	return m.loading.Tick
}

func (m *Model) Next() {
	m.focused = m.focused.Progress()
}

func (m *Model) Previous() {
	m.focused = m.focused.Regress()
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
		case "ctrl+left", "ctrl+h":
			m.RegressTask()
		case "ctrl+right", "ctrl+l":
			m.ProgressTask()
		case "n":
			ms = m // save the curent model state
			return NewForm(m.focused).Update(nil)
		}
	case Task:
		task := msg
		return m, m.lists[task.status].InsertItem(len(m.lists[task.status].Items()), task)
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
