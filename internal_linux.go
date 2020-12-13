package consoleformat

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func updateStatusLine() error {
	return nil
}

// printLine formats a string that took the entire line with prefix and suffix
// and outputs the string in the currentline
func printLine(line Line) {
	lineString := formatString(line, st.width, st.textSuffixAlignMode, st.textSuffixLength, st.overflowMode, st.padding)

	// TODO use io.Writer here.
	fmt.Print("\r")
	fmt.Print(lineString)
}

func getTerminalSize() (int, int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}
	splits := strings.Split(strings.TrimRight(string(out), "\n"), " ")
	if len(splits) < 2 {
		return 0, 0, fmt.Errorf("unexpected output from stty: %s", string(out))
	}

	height, err := strconv.Atoi(splits[0])
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse output for height: %w", err)
	}

	width, err := strconv.Atoi(splits[1])
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse output for width: %w", err)
	}

	return width, height, nil
}

func onResize(resizeChannel chan os.Signal) {
	for range resizeChannel {
		st.width, st.height, _ = getTerminalSize()
	}
}

var resizeChannel chan os.Signal

func registerResizeEvent() error {
	resizeChannel = make(chan os.Signal)
	signal.Notify(resizeChannel, syscall.SIGWINCH)
	go onResize(resizeChannel)
	return nil
}

func unregisterResizeEvent() error {
	// quit system event callback
	if resizeChannel != nil {
		signal.Reset(syscall.SIGWINCH)
		close(resizeChannel)
		resizeChannel = nil
	}

	return nil
}
