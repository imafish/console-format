package consoleformat

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func updateStatusLine() error {
	return nil
}

func printLine(line Line) error {
	lineString, err := formatLine(line, st.overflowMode, st.textSuffixLength, st.textSuffixAlignMode)
	if err != nil {
		return err
	}

	fmt.Print("\r")
	fmt.Print(lineString)
	return nil
}

func formatLine(line Line, overflowMode TextOverflowMode, suffixLength int, alignMode SuffixAlignMode) (string, error) {
	if !st.initialized {
		return "", fmt.Errorf("statusline library not initialized")
	}

	if suffixLength < 1 {
		return formatLineNoSuffix(line, overflowMode)
	}
	return formatLineWithSuffix(line, overflowMode, suffixLength, alignMode)
}

func formatLineNoSuffix(line Line, overflowMode TextOverflowMode) (string, error) {
	// TODO implement
	return line.Prefix, nil
}

func formatLineWithSuffix(line Line, overflowMode TextOverflowMode, suffixLength int, alignMode SuffixAlignMode) (string, error) {
	prefix := line.Prefix
	prefixDisplayLength := displayWidth(prefix)

	suffix := line.Suffix
	suffixDisplayLength := displayWidth(suffix)

	width := st.width

	maximumSuffixLength := width / 8
	if suffixLength > maximumSuffixLength {
		suffixLength = maximumSuffixLength
	}
	if suffixDisplayLength > suffixLength {
		if suffixLength > 4 {
			suffix = suffix[:suffixLength-2] + ".."
		} else {
			suffix = suffix[:suffixLength]
		}
	} else {
		suffix = suffix + strings.Repeat(" ", suffixLength-len(suffix))
	}

	maximumPrefixLength := width - suffixLength - st.margin*2 - st.minimumPaddingLength

	if prefixDisplayLength > maximumPrefixLength {
		if maximumPrefixLength > 7 {
			prefix = prefix[:maximumPrefixLength-2] + ".."
		} else {
			prefix = prefix[:maximumPrefixLength]
		}
	}

	paddingCount := width - st.margin*2 - displayWidth(prefix) - displayWidth(suffix)
	fullLineString := prefix + strings.Repeat(" ", st.margin) + strings.Repeat(string(st.padding), paddingCount) + strings.Repeat(" ", st.margin) + suffix
	return fullLineString, nil
}

func displayWidth(str string) int {
	// TODO should do the following:
	// count asian characters as 2
	// trim control characters
	return int(len(str))
}

func getTerminalSize() (int, int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}
	splits := strings.Split(string(out), " ")
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
