package consoleformat

/*
Code that relates to formatting a string to fit into a line goes here.
*/

import (
	"regexp"
	"strings"
)

// line represents a line to display
type line struct {
	prefix   string
	suffix   string
	noSuffix bool
}

const resetSuffix = "\x1B[0m"
const dot = "..."

// settings not settable yet:
const minimumPaddingLength = 3
const margin = 2

func formatString(l line, totalWidth int, suffixAlignMode SuffixAlignMode, suffixWidth int, overflowMode TextOverflowMode, padding rune) string {
	// TODO what to do if totalWidth is too small, e.g. 20?

	// build suffix
	if suffixAlignMode == SuffixPercentage {
		suffixWidth = int(float64(totalWidth*suffixWidth) / 100)
	}
	// the suffix should not be longer than 1/3 of entire width
	unit := totalWidth / 3
	if suffixWidth > unit {
		suffixWidth = unit
	}
	_, displayWidth, pos := analysisString(l.suffix)
	suffix := adjustStringWithData(l.suffix, suffixWidth, displayWidth, pos)
	if displayWidth < suffixWidth {
		suffix = suffix + strings.Repeat(" ", suffixWidth-displayWidth)
	}

	// build prefix
	maximumPrefixLength := totalWidth - margin*2 - minimumPaddingLength - suffixWidth
	_, displayWidth, pos = analysisString(l.prefix)
	prefix := l.prefix
	if displayWidth >= maximumPrefixLength {
		if overflowMode == TextOverflowTrim {
			displayWidth = maximumPrefixLength
			prefix = adjustStringWithData(prefix, maximumPrefixLength, displayWidth, pos)
		} else {
			// TODO implement
		}
	}

	// calculate padding length
	paddingCount := totalWidth - margin*2 - displayWidth - suffixWidth

	// construct final string
	marginString := strings.Repeat(" ", margin)
	result := prefix + marginString + strings.Repeat(string(padding), paddingCount) + marginString + suffix

	return result
}

func adjustStringWithData(str string, desiredWidth int, displayWidth int, controlCharacterPositions [][]int) string {
	if desiredWidth >= displayWidth {
		if strings.HasSuffix(str, resetSuffix) {
			return str
		}
		return str + resetSuffix
	}

	appendDot := (desiredWidth > 7)

	if controlCharacterPositions == nil {
		if appendDot {
			return str[:desiredWidth-3] + dot
		}
		return str[:desiredWidth]
	}

	if appendDot {
		desiredWidth -= 3
	}

	currentWidth := 0
	prevIndex := 0
	for i, control := range controlCharacterPositions {
		currentWidth += control[0] - prevIndex
		prevIndex = control[1]

		if currentWidth >= desiredWidth {
			offset := currentWidth - desiredWidth
			adjustedString := str[:control[0]-offset]

			if i > 0 {
				adjustedString = adjustedString + resetSuffix
			}

			if appendDot {
				adjustedString = adjustedString + dot
			}
			return adjustedString
		}
	}

	// if it reaches here, means all control characters should be included
	offset := displayWidth - desiredWidth
	adjustedString := str[:len(str)-offset] + resetSuffix
	if appendDot {
		adjustedString = adjustedString + dot
	}

	return adjustedString
}

func adjustString(str string, desiredWidth int) string {
	_, displayWidth, controlCharacterPositions := analysisString(str)
	result := adjustStringWithData(str, desiredWidth, displayWidth, controlCharacterPositions)

	return result
}

func analysisString(str string) (strippedString string, displayWidth int, controlCharacterPositions [][]int) {
	regex := regexp.MustCompile("(\x1B\\[[^m]+m)")
	controlCharacterPositions = regex.FindAllStringSubmatchIndex(str, -1)
	strippedString = regex.ReplaceAllString(str, "")
	displayWidth = len(strippedString)

	return
}
