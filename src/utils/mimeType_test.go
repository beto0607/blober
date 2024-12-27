package utils_test

import (
	"testing"

	"beto0607.com/blober/src/utils"
)

func TestMimeType(t *testing.T) {
	got := utils.GetMimeType([]byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A})
	if got != "image/png" {
		t.Errorf("Got: %s", got)
	}
}
