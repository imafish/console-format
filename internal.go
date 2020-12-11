package consoleformat

type status struct {
	statuslineMode            StatusLineMode
	statuslineSuffixAlignMode SuffixAlignMode
	statuslineSuffixLength    int
	overflowMode              TextOverflowMode
	textSuffixAlignMode       SuffixAlignMode
	textSuffixLength          int
	padding                   rune

	// settings not settable yet:
	minimumPaddingLength int
	margin               int

	// control objects
	initialized   bool
	doCallback    bool
	inCurrentLine bool

	// current status
	currentline Line
	statusline  Line
	width       int
	height      int
}

var st status
