package consoleformat

// Line represents a line to display
type Line struct {
	Prefix string
	Suffix string
}

// Config holds console format library settings
type Config struct {
	StatuslineMode            StatusLineMode
	StatuslineSuffixAlignMode SuffixAlignMode
	StatuslineSuffixLength    int
	OverflowMode              TextOverflowMode
	TextSuffixAlignMode       SuffixAlignMode
	TextSuffixLength          int
	Padding                   rune
}

// StatusLineMode decides where the status line is placed.
type StatusLineMode int

const (
	// StatusLineFollow mode display a status line next to the current line in console.
	// If console has new output, the status line will adjust itself.
	StatusLineFollow StatusLineMode = iota

	// StatusLineBottom mode displays the status line at the bottom of the console.
	StatusLineBottom StatusLineMode = iota
)

// SuffixAlignMode decides how suffix is aligned.
type SuffixAlignMode int

const (
	// SuffixFixedWidth lets the suffix has a fixed length
	SuffixFixedWidth SuffixAlignMode = iota

	// SuffixPercentage lets the suffix takes a certain percentage of the terminal width
	SuffixPercentage SuffixAlignMode = iota
)

// TextOverflowMode decides what to do if text length is longer than terminal width
type TextOverflowMode int

const (
	// TextOverflowTrim trims the text to fit to terminal and display a trailing '..'
	TextOverflowTrim TextOverflowMode = iota

	// TextOverflowMultiline will display multiple lines if text is too long
	TextOverflowMultiline TextOverflowMode = iota
)
