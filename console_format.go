package consoleformat

import (
	"fmt"
)

// Init must be called before all other functions
func Init() error {
	// Init does the following:
	// 1. get terminal window size
	// 2. register window size change event

	// 3. set default values for internal objects
	st.minimumPaddingLength = 3
	st.margin = 2

	return fmt.Errorf("not implemented")
}

// Close terminates the library.
// It moves cursor to correct position;
// It clears status line if necessary;
func Close() error {
	st.doCallback = false
	if st.inCurrentLine {
		fmt.Println()
	}

	// TODO clears status line

	return nil
}

// SetConfig configs this library
func SetConfig(config Config) {
	st.statuslineMode = config.StatuslineMode
	st.statuslineSuffixAlignMode = config.StatuslineSuffixAlignMode
	st.statuslineSuffixLength = config.StatuslineSuffixLength
	st.overflowMode = config.OverflowMode
	st.textSuffixAlignMode = config.TextSuffixAlignMode
	st.textSuffixLength = config.TextSuffixLength
	st.padding = config.Padding

	// validation
	if st.padding == rune(0) {
		st.padding = '.'
	}
}

// SetStatusLine sets the status line
func SetStatusLine(statusline Line) error {
	st.statusline = statusline
	return updateStatusLine()
}

// SetStatusLinePrefix sets the prefix
func SetStatusLinePrefix(prefix string) error {
	st.statusline.Prefix = prefix
	return updateStatusLine()
}

// SetStatusLineSuffix sets the suffix
func SetStatusLineSuffix(suffix string) error {
	st.statusline.Suffix = suffix
	return updateStatusLine()
}

// PrintInCurrentLine prints a line in the console
// It first move the cursor to beginning of line,
// then print a line that fills the entire line, and leaves the cursor at
// the end of the text
func PrintInCurrentLine(line Line) error {
	defer updateStatusLine()
	st.doCallback = false
	defer func() { st.doCallback = true }()

	err := printLine(line)
	if err == nil {
		st.inCurrentLine = true
	}
	return err
}

// Println acts similar as PrintInCurrentLine,
// except that it moves the cursor to the next line after printing
func Println(line Line) error {
	defer updateStatusLine()
	st.doCallback = false
	defer func() { st.doCallback = true }()

	err := printLine(line)
	if err != nil {
		return err
	}
	fmt.Println()

	return err
}
