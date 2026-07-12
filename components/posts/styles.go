package posts

import (
	"reddittui/components/colors"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var postsListStyle = lipgloss.NewStyle().MarginRight(4)

func NewPostsDelegate() list.DefaultDelegate {
	delegate := list.NewDefaultDelegate()

	listStyle := delegate.Styles

	selectedColor := colors.AdaptiveColor(colors.Accent)

	listStyle.NormalTitle = listStyle.NormalTitle.Bold(false)

	listStyle.SelectedTitle = listStyle.SelectedTitle.
		Bold(true).
		Foreground(selectedColor).
		BorderForeground(selectedColor)

	listStyle.SelectedDesc = listStyle.SelectedDesc.
		Foreground(colors.AdaptiveColor(colors.Accent)).
		BorderForeground(selectedColor)

	delegate.Styles = listStyle

	return delegate
}