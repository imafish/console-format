package consoleformat

type status struct {
	statuslineMode            StatusLineMode
	statuslineSuffixAlignMode SuffixAlignMode
	statuslineSuffixLength    int
	overflowMode              TextOverflowMode
	textSuffixAlignMode       SuffixAlignMode
	textSuffixLength          int
	padding                   rune

	// control objects
	initialized   bool
	doCallback    bool
	inCurrentLine bool

	// current status
	currentline line
	statusline  line
	width       int
	height      int
}

var st status
