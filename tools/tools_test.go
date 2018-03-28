package tools

import (
	"fmt"
	"regexp"
	"testing"
)

func TestParseRegex(t *testing.T) {
	var target = "abcdefg"
	//
	var str = "."
	ok, _ := regexp.MatchString(str, target)
	if !ok {
		t.Error("Err")
	}

	fmt.Println(ParseRegex(str))
	ok, _ = regexp.MatchString(ParseRegex(str), target)
	if ok {
		t.Error("Err")
	}

	//
	str = "ab g"
	ok, _ = regexp.MatchString(str, target)
	if ok {
		t.Error("Err")
	}

	fmt.Println(ParseRegex(str))
	ok, _ = regexp.MatchString(ParseRegex(str), target)
	if !ok {
		t.Error("Err")
	}

	//
	str = "^ab"
	ok, _ = regexp.MatchString(str, target)
	if !ok {
		t.Error("Err")
	}

	fmt.Println(ParseRegex(str))
	ok, _ = regexp.MatchString(ParseRegex(str), target)
	if ok {
		t.Error("Err")
	}
}

func TestRound(t *testing.T) {
	f := Round(1.12345, 0)
	if f != 1 {
		t.Error("round err", f)
		return
	}
	f = Round(1.54321, 0)
	if f != 2 {
		t.Error("round err", f)
		return
	}
	f = Round(1.12345, 1)
	if f != 1.1 {
		t.Error("round err", f)
		return
	}
	f = Round(1.55321, 1)
	if f != 1.6 {
		t.Error("round err", f)
		return
	}
}
