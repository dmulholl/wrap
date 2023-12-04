package wrap

import (
	"regexp"
	"slices"
	"strings"
	"unicode"
)

// Wrapper wraps text to within a specified maximum line-length, using a set of customizable breakpoints where possible.
type Wrapper struct {
	// Wrapping is allowed before these characters.
	BreakBefore []rune

	// Wrapping is allowed after these characters.
	BreakAfter []rune

	// Line-break sequence to use in the output.
	Linebreak string

	// If true, collapses all contiguous sequences of whitespace in the input into a single space.
	CollapseWhitespace bool
}

// NewWrapper returns a Wrapper instance that defaults to collapsing whitespace and wrapping on spaces.
func NewWrapper() *Wrapper {
	return &Wrapper{
		BreakAfter:         []rune{' '},
		BreakBefore:        []rune{' '},
		Linebreak:          "\n",
		CollapseWhitespace: true,
	}
}

// Wrap wraps the input text to within the specified maximum line-length.
func (w *Wrapper) Wrap(text string, limit int) string {
	if text == "" || limit < 1 {
		return ""
	}

	if w.CollapseWhitespace {
		r := regexp.MustCompile(`\s+`)
		text = r.ReplaceAllLiteralString(text, " ")
	}

	output := []string{}
	runes := []rune(text)
	rune_count := len(runes)

	start_index := 0

	for start_index < rune_count {
		for start_index < rune_count && unicode.IsSpace(runes[start_index]) {
			start_index += 1
		}

		if start_index == rune_count {
			break
		}

		end_index := start_index + limit
		if end_index > rune_count {
			end_index = rune_count
		}

		if end_index == rune_count {
			line := string(runes[start_index:end_index])
			output = append(output, strings.TrimSpace(line))
			break
		}

		for end_index > start_index {
			if slices.Contains(w.BreakAfter, runes[end_index-1]) || slices.Contains(w.BreakBefore, runes[end_index]) {
				break
			}
			end_index -= 1
		}

		if end_index == start_index {
			end_index = start_index + limit
		}

		line := string(runes[start_index:end_index])
		output = append(output, strings.TrimSpace(line))
		start_index = end_index
	}

	return strings.Join(output, w.Linebreak)
}

// Wrap is a convenience function for wrapping text using a default Wrapper instance that wraps on spaces.
func Wrap(text string, limit int) string {
	return NewWrapper().Wrap(text, limit)
}
