package progressbar

import (
	"fmt"
	"reflect"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"
)

func newProgressBarModel(cfg *Config) progress.Model {
	if cfg.Styles.GradientFrom != "" && cfg.Styles.GradientTo != "" {
		return progress.New(
			progress.WithScaledGradient(cfg.Styles.GradientFrom, cfg.Styles.GradientTo),
			progress.WithWidth(40),
			progress.WithoutPercentage(),
		)
	}
	return progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(40),
		progress.WithoutPercentage(),
	)
}

func isEmpty(s interface{}) bool {
	switch s.(type) {
	case string:
		return len(fmt.Sprint(s)) == 0
	case Styles:
		return reflect.DeepEqual(s, Styles{})
	case lipgloss.Style:
		return reflect.DeepEqual(s, lipgloss.Style{})
	case lipgloss.AdaptiveColor:
		return reflect.DeepEqual(s, lipgloss.AdaptiveColor{})
	default:
		return false
	}
}
