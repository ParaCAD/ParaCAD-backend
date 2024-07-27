package utils_test

import (
	"testing"

	"github.com/ParaCAD/ParaCAD-backend/utils"
)

func TestCreateRandomString(t *testing.T) {
	testsCases := []struct {
		input int
	}{
		{input: 1},
		{input: 2},
		{input: 5},
		{input: 10},
		{input: 20},
	}

	for _, tt := range testsCases {
		t.Run("", func(t *testing.T) {
			output := utils.CreateRandomString(tt.input)
			if len(output) != tt.input {
				t.Errorf("len(utils.CreateRandomString(%d)), expected: %d, got: %d (%s)",
					tt.input, tt.input, len(output), output)
			}
		})
	}
}
