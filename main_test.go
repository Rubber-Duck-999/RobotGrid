package main

import (
	"strconv"
	"testing"
)

func TestDigitsTwo(t *testing.T) {
	sum := FilterDigits(59)
	if sum != 14 {
		t.Error("Error on filtering digits ")
	}
}

func TestDigitsThree(t *testing.T) {
	sum := FilterDigits(111)
	if sum != 3 {
		t.Error("Error on filtering digits")
	}
}

// Checks whether the absolute sum
// of all points equals zero
func TestAbsSumZero(t *testing.T) {
	x1, y1 := 0, 0
	total := GetSum(x1, y1)
	if total != 0 {
		t.Error("Total doesn't match Zero")
	}
}

// Checks whether the absolute sum is less than
// 23 like expected: 6 + 6 + 6 for positive
func TestAbsSumPos(t *testing.T) {
	x1, y1 := 66, 6
	total := GetSum(x1, y1)
	if total != 18 {
		t.Error("Total doesn't match Positive value required")
		t.Error("Actual: " + strconv.Itoa(total))
	}
}

// Checks whether the absolute sum is less than
// 23 like expected: -6 + 9 + -6 for positive
func TestAbsSumNegative(t *testing.T) {
	x1, y1 := -69, -6
	total := GetSum(x1, y1)
	if total != 21 {
		t.Error("Total doesn't match Positive value required from negative values")
	}
}

// Checks whether the total is less than max should pass
func TestIsSafeBelow(t *testing.T) {
	if IsSafe(6) != true {
		t.Error("IsSafe is below so should pass")
	}
}

// Checks whether the total is equal to safe value and
// should be safe
func TestIsSafeEqual(t *testing.T) {
	if IsSafe(23) != true {
		t.Error("IsSafe is equal to value so should pass")
	}
}

// Checks whether the total is equal to safe value and
// should be safe
func TestIsSafeAbove(t *testing.T) {
	if IsSafe(90) == true {
		t.Error("IsSafe is equal to value so should pass")
	}
}
