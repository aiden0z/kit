package base

import (
	"testing"
)

func TestIntSlice(t *testing.T) {

	ints := []int{7, 10, 4, 3, 1, 2, 8, 11}
	order := IntComparableSlice(ints)

	for i, v := range ints {
		if order[i].CompareTo(Int(v)) != 0 {
			t.Error("IntSlice work error")

		}

	}
}

func TestInt_CompareTo(t *testing.T) {
	if Int(2).CompareTo(Rune('b')) != 1 {
		t.Error("compare different type not return 1")

	}
	if Int(2).CompareTo(Int(3)) != -1 {
		t.Error("less compare not return -1")
	}

	if Int(2).CompareTo(Int(1)) != 1 {
		t.Error("great compare not return 1")
	}

	if Int(2).CompareTo(Int(2)) != 0 {
		t.Error("equal compare not return 0")

	}

}

func TestRuneSlice(t *testing.T) {
	runes := []int{'a', 'b', 'c'}
	order := IntComparableSlice(runes)

	for i, v := range runes{
		if order[i].CompareTo(Int(v)) != 0 {
			t.Error("IntSlice work error")

		}

	}
}

func TestRune_CompareTo(t *testing.T) {

	if Rune('a').CompareTo(Int(2)) != 1 {
		t.Error("compare different type not return 1")

	}

	if Rune('a').CompareTo(Rune('b')) != -1 {
		t.Error("less compare not return -1")
	}

	if Rune('b').CompareTo(Rune('a')) != 1 {
		t.Error("great compare not return 1")
	}

	if Rune('b').CompareTo(Rune('b')) != 0 {
		t.Error("equal compare not return 0")

	}

}
