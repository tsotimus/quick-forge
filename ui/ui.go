package ui

import (
	"errors"
	"fmt"
	"os"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/choose"
	"github.com/cqroot/prompt/input"
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

func AskYesNo(question string) bool {
	answer, err := prompt.New().Ask(question).Choose(
		[]string{"Yes", "No"},
		choose.WithTheme(choose.ThemeLine),
		choose.WithKeyMap(choose.HorizontalKeyMap),
		// choose.WithSelectedMark("👉"),
	)
	CheckErr(err)
	return answer == "Yes"
}

func AskForInput(question string, defaultValue string, opts ...input.Option) string {
	answer, err := prompt.New().Ask(question).Input(defaultValue, opts...)
	CheckErr(err)
	return answer
}
