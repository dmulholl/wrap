package wrap

import (
	"strings"
)

// UnconditionalWrapper wraps the input text at exactly the specified line-length.
type UnconditionalWrapper struct {
	// Line-break sequence to use in the output.
	Linebreak string
}

// NewUnconditionalWrapper returns an UnconditionalWrapper instance.
func NewUnconditionalWrapper() *UnconditionalWrapper {
	return &UnconditionalWrapper{
		Linebreak: "\n",
	}
}

// Wrap wraps the input text at exactly the specified line-length.
func (uw *UnconditionalWrapper) Wrap(text string, limit int) string {
	if text == "" || limit < 1 {
		return ""
	}

	output := []string{}
	runes := []rune(text)
	rune_count := len(runes)

	start_index := 0

	for start_index < rune_count {
		end_index := start_index + limit
		if end_index > rune_count {
			end_index = rune_count
		}

		line := string(runes[start_index:end_index])
		output = append(output, line)
		start_index = end_index
	}

	return strings.Join(output, uw.Linebreak)
}

// UnconditionallyWrap is a convenience function for wrapping text using a default UnconditionalWrapper instance.
func UnconditionallyWrap(text string, limit int) string {
	return NewUnconditionalWrapper().Wrap(text, limit)
}
