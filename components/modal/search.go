package modal

import (
	"reddittui/components/colors"
	"reddittui/components/messages"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	searchHelpText     = "Choose a subreddit:"
	searchPlaceholder  = "subreddit"
	defaultSearchWidth = 35
)

type SubredditSearchModal struct {
	textinput.Model
	style lipgloss.Style
	searchHelpStyle lipgloss.Style
	searchModelStyle lipgloss.Style
}

func NewSubredditSearchModal() SubredditSearchModal {
	searchHelpStyle := lipgloss.NewStyle().Foreground(colors.AdaptiveColor(colors.Text)).Italic(true)
	searchModelStyle := lipgloss.NewStyle().Foreground(colors.AdaptiveColor(colors.Accent))
	searchTextInput := textinput.New()
	searchTextInput.Placeholder = searchPlaceholder
	searchTextInput.ShowSuggestions = true
	searchTextInput.SetSuggestions(subredditSuggestions)
	searchTextInput.CharLimit = 30

	return SubredditSearchModal{
		Model: searchTextInput,
		style: lipgloss.NewStyle(),
		searchHelpStyle: searchHelpStyle,
		searchModelStyle: searchModelStyle,
	}
}

func (s SubredditSearchModal) Init() tea.Cmd {
	return nil
}

func (s SubredditSearchModal) Update(msg tea.Msg) (SubredditSearchModal, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return s, messages.LoadSubreddit(s.Value())
		case "esc":
			return s, messages.ExitModal
		}
	}

	var cmd tea.Cmd
	s.Model, cmd = s.Model.Update(msg)
	return s, cmd
}

func (s SubredditSearchModal) View() string {
	titleView := s.searchHelpStyle.Render(searchHelpText)
	modelView := s.searchModelStyle.Render(s.Model.View())
	joined := lipgloss.JoinVertical(lipgloss.Left, titleView, modelView)
	return s.style.Render(joined)
}

func (s *SubredditSearchModal) SetSize(w, h int) {
	searchW := min(w-s.style.GetHorizontalFrameSize(), defaultSearchWidth)
	s.style = s.style.Width(searchW)
}

func (s *SubredditSearchModal) Blur() {
	s.Model.Blur()
	s.Reset()
}
