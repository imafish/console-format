package consoleformat

import "testing"

func TestAnalysisString(t *testing.T) {
	str := "\x1B[34mabc\x1B[0m"

	got, length, pos := analysisString(str)
	if got != "abc" {
		t.Errorf("analysisString(%s) = %s, expected %s", str, got, "abc")
	}
	if length != 3 {
		t.Errorf("analysisString(%s) = %d, expected %d", str, length, 3)
	}
	if len(pos) != 2 {
		t.Errorf("analysisString(%s) = %v, expected %d", str, pos, 2)
	}
}

func TestAdjustString(t *testing.T) {
	str := "123456\x1B[34mabcdefg\x1B[35mABCDEFGH\x1B[0m"

	testAdjustString(str, 4, "1234", t)
	testAdjustString(str, 8, "12345...", t)
	testAdjustString(str, 12, "123456\x1B[34mabc\x1B[0m...", t)
	testAdjustString(str, 16, "123456\x1B[34mabcdefg\x1B[0m...", t)
	testAdjustString(str, 20, "123456\x1B[34mabcdefg\x1B[35mABCD\x1B[0m...", t)
	testAdjustString(str, 24, "123456\x1B[34mabcdefg\x1B[35mABCDEFGH\x1B[0m", t)
	testAdjustString(str, 32, "123456\x1B[34mabcdefg\x1B[35mABCDEFGH\x1B[0m", t)
}

func testAdjustString(str string, expectedWidth int, expectedString string, t *testing.T) {
	trim := adjustString(str, expectedWidth)
	if trim != expectedString {
		t.Errorf(`adjustString("%s", %d) = "%s", expected %s`, str, expectedWidth, trim, expectedString)
	}
}
