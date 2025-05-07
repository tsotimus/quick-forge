package ui

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/choose"
	"github.com/cqroot/prompt/input"
	"github.com/cqroot/prompt/multichoose"
)

func CheckErr(err error) {
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			fmt.Fprintln(os.Stderr, "Exited.")
			os.Exit(1)
		} else {
			panic(err)
		}
	}
}

var (
	highlightStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#4fc7c1")).Bold(true)
	normalStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
)

var BlueHighlightTheme choose.Theme = func(choices []choose.Choice, selectedIndex int) string {
	s := ""
	for i, choice := range choices {
		text := choice.Text
		if i == selectedIndex {
			s += highlightStyle.Render(text) // Highlight the selected choice
		} else {
			s += normalStyle.Render(text) // Normal style for other choices
		}

		// If it's not the last choice, add the separator (no styling)
		if i != len(choices)-1 {
			s += " / " // Add the unstyled separator
		}
	}
	return s
}

func AskYesNo(question string) bool {
	answer, err := prompt.New().Ask(question).Choose(
		[]string{"Yes", "No"},
		choose.WithTheme(choose.Theme(BlueHighlightTheme)),
		choose.WithKeyMap(choose.HorizontalKeyMap),
	)
	CheckErr(err)
	return answer == "Yes"
}

func AskForInput(question string, defaultValue string, opts ...input.Option) string {
	answer, err := prompt.New().Ask(question).Input(defaultValue, opts...)
	CheckErr(err)
	return answer
}

func AskSimpleChoice(question string, choices []string, opts ...choose.Option) string {
	answer, err := prompt.New().Ask(question).Choose(choices,
		choose.WithTheme(choose.Theme(BlueHighlightTheme)),
		choose.WithKeyMap(choose.HorizontalKeyMap),
	)
	CheckErr(err)
	return answer
}

func AskMultiChoice(question string, choices []string, opts ...multichoose.Option) []string {
	answer, err := prompt.New().Ask(question).MultiChoose(
		choices,
		append([]multichoose.Option{}, opts...)...,
	)
	CheckErr(err)
	return answer
}
