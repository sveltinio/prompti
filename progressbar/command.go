package progressbar

import (
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// Config represents the struct to configure the tui command.
type Config struct {
	Items         []string
	OnProgressCmd func(string) tea.Cmd
	//If true, the progressbar runs commands concurrently (tea.Batch) else in order (tea.Sequence).
	OnProgressMsg   string
	OnCompletesMsg  string
	RunConcurrently bool
	Styles          Styles
}

// setDefaults sets default values for Config if not present.
func (cfg *Config) setDefaults() {
	if cfg.OnCompletesMsg == "" {
		cfg.OnCompletesMsg = "Done!"
	}

	if cfg.OnProgressCmd == nil {
		cfg.OnProgressCmd = func(item string) tea.Cmd {
			// This is where you'd do i/o stuff to download and install packages. In
			// our case we're just pausing for a moment to simulate the process.
			d := time.Millisecond * time.Duration(rand.Intn(500))
			return tea.Tick(d, func(t time.Time) tea.Msg {
				return IncrementMsg(item)
			})
		}
	}
}

func (cfg *Config) initialModel() model {
	return model{
		items:                cfg.Items,
		onProgressCmd:        cfg.OnProgressCmd,
		onProgressMsg:        cfg.OnProgressMsg,
		onCompletedMsg:       cfg.OnCompletesMsg,
		runConcurrently:      cfg.RunConcurrently,
		showLabel:            cfg.Styles.ShowLabel,
		currentItemNameStyle: cfg.Styles.CurrentItemStyle,
		progress:             newProgressBarModel(cfg),
	}
}

// Run provides a shell script interface for prompting a user to confirm an
// action with an affirmative or negative answer.
func Run(cfg *Config) (tea.Model, error) {
	cfg.setDefaults()
	return tea.NewProgram(cfg.initialModel()).Run()
}
