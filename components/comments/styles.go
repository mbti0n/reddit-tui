package comments

import (
	"reddittui/components/colors"

	"github.com/charmbracelet/lipgloss"
)

var viewportStyle = lipgloss.NewStyle().Margin(0, 2, 1, 2)

var (
	commentAuthorStyle  = lipgloss.NewStyle().Foreground(colors.AdaptiveColor(colors.Text)).Bold(true)
	commentDateStyle    = lipgloss.NewStyle().Foreground(colors.AdaptiveColor(colors.Subtext)).Italic(true)
	commentTextStyle    = lipgloss.NewStyle().Foreground(colors.AdaptiveColor(colors.Text))
	popularPointsStyle  = lipgloss.NewStyle().Foreground(colors.AdaptiveColor(colors.Accent))
	defaultPointsStyle  = lipgloss.NewStyle().Foreground(colors.AdaptiveColor(colors.Accent))
	negativePointsStyle = lipgloss.NewStyle().Foreground(colors.AdaptiveColor(colors.Negative))
	collapsedStyle      = lipgloss.NewStyle().Foreground(colors.AdaptiveColor(colors.Accent))
)

var (
	postAuthorStyle    = lipgloss.NewStyle().Foreground(colors.AdaptiveColor(colors.Subtext))
	postPointsStyle    = lipgloss.NewStyle().Foreground(colors.AdaptiveColor(colors.Accent))
	totalCommentsStyle = lipgloss.NewStyle().Foreground(colors.AdaptiveColor(colors.Text))
	postTextStyle      = lipgloss.NewStyle().Foreground(colors.AdaptiveColor(colors.Text))
	postTimestampStyle = lipgloss.NewStyle().Foreground(colors.AdaptiveColor(colors.Text)).Faint(true)
)

func initStyles() {
    commentAuthorStyle = lipgloss.NewStyle().
        Foreground(colors.AdaptiveColor(colors.Text)).
        Bold(true)

    commentDateStyle = lipgloss.NewStyle().
        Foreground(colors.AdaptiveColor(colors.Subtext)).
        Italic(true)

    commentTextStyle = lipgloss.NewStyle().
        Foreground(colors.AdaptiveColor(colors.Text))

    popularPointsStyle = lipgloss.NewStyle().
        Foreground(colors.AdaptiveColor(colors.Accent))

    defaultPointsStyle = lipgloss.NewStyle().
        Foreground(colors.AdaptiveColor(colors.Accent))

    negativePointsStyle = lipgloss.NewStyle().
        Foreground(colors.AdaptiveColor(colors.Negative))

    collapsedStyle = lipgloss.NewStyle().
        Foreground(colors.AdaptiveColor(colors.Accent))

    postAuthorStyle = lipgloss.NewStyle().
        Foreground(colors.AdaptiveColor(colors.Subtext))

    postPointsStyle = lipgloss.NewStyle().
        Foreground(colors.AdaptiveColor(colors.Accent))

    totalCommentsStyle = lipgloss.NewStyle().
        Foreground(colors.AdaptiveColor(colors.Text))

    postTextStyle = lipgloss.NewStyle().
        Foreground(colors.AdaptiveColor(colors.Text))

    postTimestampStyle = lipgloss.NewStyle().
        Foreground(colors.AdaptiveColor(colors.Text)).
        Faint(true)
}
