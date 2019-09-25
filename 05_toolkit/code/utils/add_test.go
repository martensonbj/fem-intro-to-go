package utils

import (
	"fem-apps/utils"
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 4
	actual := utils.Add(2, 2)

	if actual != expected {
		t.Errorf("Add function does not add up: Expected: %d, Actual: %d", expected, actual)
	}
}
