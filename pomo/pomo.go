package pomo

import (
	//"fmt"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/timer"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	//subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	green = lipgloss.Color("#009000")
	yellow = lipgloss.Color("#F8E71C")
	red = lipgloss.Color("#D0021B")
)

type Pomo struct {
	duration time.Duration
	timer timer.Model
	width int
	height int
}

func NewPomo(duration time.Duration) Pomo {
	pomo := Pomo{
		duration: duration,
		timer: timer.New(duration),
	}
	return pomo
}

func (p Pomo) Init() tea.Cmd {
	println("init")
	cmd := p.timer.Init()
	println("init done")
	return cmd
}

func (p Pomo) Update (msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TickMsg:
		var cmd tea.Cmd
		p.timer, cmd = p.timer.Update(msg)
		return p, cmd

	case timer.StartStopMsg:
		var cmd tea.Cmd
		p.timer, cmd = p.timer.Update(msg)
		return p, cmd

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, key.NewBinding(key.WithKeys("ctrl+c", "q"))):
			return p, tea.Quit

		case key.Matches(msg, key.NewBinding(key.WithKeys("r"))):
			var cmd tea.Cmd
			if !p.timer.Timedout() {
				p.timer.Stop()
				p.timer = timer.New(p.duration)
				cmd = p.timer.Start()
			} else {
				p.timer = timer.New(p.duration)
				cmd = p.timer.Init()
			}
			return p, cmd

		case key.Matches(msg, key.NewBinding(key.WithKeys("s"))):
			var cmd tea.Cmd
			if p.timer.Running() {
				cmd = p.timer.Stop()
			} else {
				cmd = p.timer.Start()
			}
			return p, cmd
		}

	case tea.WindowSizeMsg:
		p.height = msg.Height
		p.width = msg.Width
		return p, nil
	}
	return p, nil
}

func (p Pomo) View() string {
	var s string
	if p.timer.Timedout() {
		s = "All done!"
	} else if !p.timer.Timedout() && !p.timer.Running() {
		s = " -- stopped --\n" + p.timer.Timeout.String()
	} else	{
		s = p.timer.Timeout.String()
	}
	s += "\n"

	//return fmt.Sprintf("time remaining: %s", s)
	var color lipgloss.Color
	switch {
	case p.timer.Timedout() || !p.timer.Running():
		color = red
	case p.timer.Timeout.Seconds() > 300:
		color = green
	case p.timer.Timeout.Seconds() > 0 && p.timer.Timeout.Seconds() <= 300:
		color = yellow
	}

	style := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Border(lipgloss.RoundedBorder(), true).
		Width(10).
		Height(10).
		Background(color).
		Padding(10)
	dialog := lipgloss.Place(p.width, p.height, lipgloss.Center, lipgloss.Center, style.Render(s))
	helpText := "Press 'q' to exit, 'r' to restart timer, 's' to start or stop timer"

	finalString := lipgloss.JoinVertical(lipgloss.Center, dialog, helpText)

	return finalString
}
