package confirm

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	okLabel     = "Yes"
	cancelLabel = "No"
)

// Config represents the struct to configure the tui command.
type Config struct {
	Question          string
	OkButtonLabel     string
	CancelButtonLabel string
	Styles            Styles
}

// setDefaults sets default values for Config if not present.
func (cfg *Config) setDefaults() {
	cfg.OkButtonLabel = setOkButtonLabel(cfg)
	cfg.CancelButtonLabel = setCancelButtonLabel(cfg)

	if isEmpty(cfg.Styles) {
		cfg.Styles = DefaultStyles()
	} else {
		cfg.Styles.setDefaults()
	}

}

func (cfg *Config) initialModel() model {
	return model{
		question:          cfg.Question,
		okButtonLabel:     cfg.OkButtonLabel,
		cancelButtonLabel: cfg.CancelButtonLabel,
		confirmation:      true,
		// styles
		styles: cfg.Styles,
	}
}

// Run provides a shell script interface for prompting a user to confirm an
// action with an affirmative or negative answer.
func Run(cfg *Config) (bool, error) {
	cfg.setDefaults()
	m, err := tea.NewProgram(cfg.initialModel(), tea.WithOutput(os.Stderr)).StartReturningModel()

	if err != nil {
		return false, fmt.Errorf("unable to run confirm: %w", err)
	}

	if m.(model).confirmation {
		return true, nil
	}

	return false, nil
}

// =================================================================

func setOkButtonLabel(cfg *Config) string {
	if isEmpty(cfg.OkButtonLabel) {
		return okLabel
	}
	return cfg.OkButtonLabel
}

func setCancelButtonLabel(cfg *Config) string {
	if isEmpty(cfg.CancelButtonLabel) {
		return cancelLabel
	}
	return cfg.CancelButtonLabel
}
