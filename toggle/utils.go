package toggle

import (
	"fmt"
	"reflect"

	"github.com/charmbracelet/lipgloss"
)

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
