package consoleformat

import (
	"fmt"
)

func init() {
	// Init does the following:
	// 1. get terminal window size
	var err error
	st.width, st.height, err = getTerminalSize()
	if err != nil {
		st.err = err
		return
	}

	// 2. register window size change event
	err = registerResizeEvent()
	if err != nil {
		st.err = err
		return

	}

	// 3. set default values for internal objects
	st.initialized = true
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

	err := unregisterResizeEvent()
	if err != nil {
		return fmt.Errorf("failed to unregister resize event: %w", err)
	}

	st.initialized = false
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
func SetStatusLine(prefix, suffix string) error {
	st.statusline = line{prefix, suffix, true}
	return updateStatusLine()
}

// SetStatusLinePrefix sets the prefix
func SetStatusLinePrefix(prefix string) error {
	st.statusline.prefix = prefix
	return updateStatusLine()
}

// SetStatusLineSuffix sets the suffix
func SetStatusLineSuffix(suffix string) error {
	st.statusline.suffix = suffix
	return updateStatusLine()
}

// Print prints a line in the console
// It first move the cursor to beginning of line,
// then print a line that fills the entire line, and leaves the cursor at
// the end of the text
func Print(prefix, suffix string) {
	defer updateStatusLine()
	st.doCallback = false
	defer func() { st.doCallback = true }()

	printLine(line{prefix: prefix, suffix: suffix})
	st.inCurrentLine = true
}

// Println acts similar as PrintInCurrentLine,
// except that it moves the cursor to the next line after printing
func Println(prefix, suffix string) {
	defer updateStatusLine()
	st.doCallback = false
	defer func() { st.doCallback = true }()

	printLine(line{prefix: prefix, suffix: suffix})
	NextLine()
}

// PrintNoSuffix prints a line in the console
// It first move the cursor to beginning of line,
// then print a line that fills the entire line, and leaves the cursor at
// the end of the text
func PrintNoSuffix(prefix string) {
	defer updateStatusLine()
	st.doCallback = false
	defer func() { st.doCallback = true }()

	printLine(line{prefix: prefix, noSuffix: true})
	st.inCurrentLine = true
}

// PrintlnNoSuffix acts similar as PrintInCurrentLine,
// except that it moves the cursor to the next line after printing
func PrintlnNoSuffix(prefix string) {
	defer updateStatusLine()
	st.doCallback = false
	defer func() { st.doCallback = true }()

	printLine(line{prefix: prefix, noSuffix: true})
	NextLine()
}

// NextLine moves cursor to next line
func NextLine() {
	fmt.Println()
	st.inCurrentLine = false
}
