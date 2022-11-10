package input

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Config represents the struct to configure the tui input.
type Config struct {
	Message      string
	Placeholder  string
	Initial      string
	ErrorMsg     string
	Password     bool
	ValidateFunc ValidateFunc

	ShowResult bool
	Styles     Styles
}

// setDefaults sets default values for Config struct.
func (cfg *Config) setDefaults() {
	cfg.ShowResult = true
	cfg.Placeholder = setDefaultPlaceholderMsg(cfg)

	if isEmpty(cfg.Styles) {
		cfg.Styles = DefaultStyles()
	} else {
		cfg.Styles.setDefaults()
	}
}

func (cfg *Config) initialModel() model {
	ti := textinput.New()

	ti.Focus()
	ti.Prompt = promptMessage(cfg.Message, cfg.Styles.PrefixIcon, cfg.Styles.PrefixIconColor)

	ti.Placeholder = cfg.Placeholder
	ti.TextStyle = cfg.Styles.TextStyle
	ti.PromptStyle = prefixIconStyle(cfg.Styles.PrefixIconColor)
	ti.BackgroundStyle = cfg.Styles.BackgroundStyle
	ti.PlaceholderStyle = cfg.Styles.PlaceholderStyle
	ti.CursorStyle = cfg.Styles.CursorStyle
	if cfg.Password {
		ti.EchoMode = textinput.EchoPassword
		ti.EchoCharacter = '*'
	}

	return model{
		textInput:    ti,
		message:      cfg.Message,
		placeholder:  cfg.Placeholder,
		initial:      cfg.Initial,
		err:          nil,
		errMsg:       cfg.ErrorMsg,
		validateFunc: cfg.ValidateFunc,
	}
}

// Run is used to prompt an input the user and retrieve the value.
func Run(cfg *Config) (string, error) {
	cfg.setDefaults()
	p := tea.NewProgram(cfg.initialModel(), tea.WithOutput(os.Stderr))

	tm, err := p.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run input: %w", err)
	}
	m := tm.(model)

	if m.quitting || isEmpty(m.textInput.Value()) {
		if m.err == nil {
			m.err = errors.New(m.errMsg)
		}
		fmt.Println(errorMessage(m.err.Error()))
		return "", m.err
	}

	if cfg.ShowResult && !cfg.Password {
		fmt.Println(resultMessage(m.message, m.textInput.Value()))
	}

	return m.textInput.Value(), m.err
}

// =================================================================

func setDefaultPlaceholderMsg(cfg *Config) string {
	placeholderMsg := ""
	if !isEmpty(cfg.Initial) {
		placeholderMsg = fmt.Sprintf("%s (Default: %s)", cfg.Placeholder, cfg.Initial)
	} else {
		placeholderMsg = cfg.Placeholder
	}
	return placeholderMsg
}
