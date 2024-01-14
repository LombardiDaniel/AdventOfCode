package solv_test

import (
	"adventOfCode-015-1/solver"
	"testing"
)

func TestCountCharInStrReturnZero(t *testing.T) {
	var testStr string
	var testChar rune

	testStr = "aaaaa"
	testChar = 'b'

	var res int
	res = solver.CountCharInStr(testChar, testStr)

	if res != 0 {
		t.Errorf("Test failed res='%d' for: testStr='%s', testChar='%c", res, testStr, testChar)
	}

	testStr = ""
	testChar = 'b'

	res = solver.CountCharInStr(testChar, testStr)

	if res != 0 {
		t.Errorf("Test failed res='%d' for: testStr='%s', testChar='%c", res, testStr, testChar)
	}

	testStr = ""
	testChar = 'a'

	res = solver.CountCharInStr(testChar, testStr)

	if res != 0 {
		t.Errorf("Test failed res='%d' for: testStr='%s', testChar='%c", res, testStr, testChar)
	}
}

func TestCountCharInStrReturnValus(t *testing.T) {
	var testStr string
	var testChar rune

	testStr = "aaaaa"
	testChar = 'a'

	var res int
	res = solver.CountCharInStr(testChar, testStr)

	if res != 5 {
		t.Errorf("Test failed res='%d' for: testStr='%s', testChar='%c", res, testStr, testChar)
	}

	testStr = "abcdefg"
	testChar = 'a'

	res = solver.CountCharInStr(testChar, testStr)

	if res != 1 {
		t.Errorf("Test failed res='%d' for: testStr='%s', testChar='%c", res, testStr, testChar)
	}

	testStr = "abcdabcd"
	testChar = 'a'

	res = solver.CountCharInStr(testChar, testStr)

	if res != 2 {
		t.Errorf("Test failed res='%d' for: testStr='%s', testChar='%c", res, testStr, testChar)
	}
}


func TestCountSantaFloors(t *testing.T) {

	var res int
	var testStr string

	testStr = "(())"
	res = solver.CountSantaFloors(testStr)

	if res != 0 {
		t.Errorf("Error for '%s'=>'%d'", testStr, res)
	}

	testStr = "()()"
	res = solver.CountSantaFloors(testStr)

	if res != 0 {
		t.Errorf("Error for '%s'=>'%d'", testStr, res)
	}

	testStr = "((("
	res = solver.CountSantaFloors(testStr)

	if res != 3 {
		t.Errorf("Error for '%s'=>'%d'", testStr, res)
	}

	testStr = "(()(()("
	res = solver.CountSantaFloors(testStr)

	if res != 3 {
		t.Errorf("Error for '%s'=>'%d'", testStr, res)
	}

	testStr = "))((((("
	res = solver.CountSantaFloors(testStr)

	if res != 3 {
		t.Errorf("Error for '%s'=>'%d'", testStr, res)
	}
}

func TestNavigateSantaFloors(t *testing.T) {

	var res int
	var testStr string
	
	testStr = ")"
	res = solver.NavigateSantaFloors(testStr)

	if res != 1 {
		t.Errorf("Error for '%s'=>'%d'", testStr, res)
	}

	testStr = "()())"
	res = solver.NavigateSantaFloors(testStr)

	if res != 5 {
		t.Errorf("Error for '%s'=>'%d'", testStr, res)
	}
}