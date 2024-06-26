package main

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	ColumnStyle = lipgloss.NewStyle().
			Padding(1, 2)
	FocusedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))
	HelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
)

func UpdateStyleWidths(width, height int) {
	ColumnStyle.Height(height)
	ColumnStyle.Width(width)
	FocusedStyle.Height(height)
	FocusedStyle.Width(width)
}
