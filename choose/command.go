package choose

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// Config represents the struct to configure the tui command.
type Config struct {
	Title           string
	ListHeight      int
	DefaultWidth    int
	ErrorMsg        string
	ShowHelp        bool
	ShowStatusBar   bool
	EnableFiltering bool // ShowHelp must be true otherwise you will now see the help
	// styles
	Styles Styles
}

// setDefaults sets default values for Config if not present.
func (cfg *Config) setDefaults(numOfItems int) {
	cfg.ListHeight = setListHeight(cfg, numOfItems)

	if isEmpty(cfg.Styles) {
		cfg.Styles = DefaultStyles()
	} else {
		cfg.Styles.setDefaults()
	}
}

func (cfg *Config) initialModel(items []list.Item) model {
	itemDelegate := itemDelegate{
		ItemIcon:          cfg.Styles.ItemIcon,
		ItemStyle:         cfg.Styles.ItemStyle,
		SelectedItemStyle: cfg.Styles.SelectedItemStyle,
	}
	l := list.New(items, itemDelegate, cfg.DefaultWidth, cfg.ListHeight)
	l.Title = titleMessage(cfg.Styles.PrefixIcon, cfg.Styles.PrefixIconColor, cfg.Styles.TitleStyle, cfg.Title)
	l.Styles.Title = cfg.Styles.TitleStyle
	l.Styles.TitleBar = cfg.Styles.TitleBarStyle
	l.SetShowHelp(cfg.ShowHelp)
	l.SetShowStatusBar(cfg.ShowStatusBar)
	l.SetFilteringEnabled(cfg.EnableFiltering)

	return model{
		list: l,
		err:  errors.New(cfg.ErrorMsg),
	}
}

// Run is used to prompt a list of available options to the user and retrieve the selection.
func Run(cfg *Config, items []list.Item) (string, error) {
	cfg.setDefaults(len(items))
	p := tea.NewProgram(cfg.initialModel(items), tea.WithOutput(os.Stderr))

	tm, err := p.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run choose: %w", err)
	}
	m := tm.(model)

	if m.quitting {
		fmt.Println(errorMessage(m.err.Error()))
		return "", fmt.Errorf(m.err.Error())
	}

	fmt.Println(resultMessage(cfg.Title, m.list.SelectedItem().FilterValue()))
	return m.list.SelectedItem().FilterValue(), nil
}

// =================================================================

func setListHeight(cfg *Config, numOfItems int) int {
	if cfg.ListHeight == 0 {
		if cfg.ShowHelp || cfg.ShowStatusBar || cfg.EnableFiltering {
			return numOfItems * 4
		}
		return numOfItems * 3
	}
	return cfg.ListHeight
}
