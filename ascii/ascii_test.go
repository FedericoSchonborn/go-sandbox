package ascii_test

import (
	"fmt"
	"testing"

	"github.com/FedericoSchonborn/go-sandbox/ascii"
)

func TestIs(t *testing.T) {
	tt := []struct {
		Char     byte
		Func     func(byte) bool
		FuncName string
		Result   bool
	}{
		{'a', ascii.IsAlpha, "IsAlpha", true},
		{'z', ascii.IsAlpha, "IsAlpha", true},
		{'0', ascii.IsAlpha, "IsAlpha", false},
		{'9', ascii.IsAlpha, "IsAlpha", false},
		{'\n', ascii.IsAlpha, "IsAlpha", false},
		{'\t', ascii.IsAlpha, "IsAlpha", false},
		{0x1F, ascii.IsAlpha, "IsAlpha", false},
		{0x7F, ascii.IsAlpha, "IsAlpha", false},
		{'a', ascii.IsDigit, "IsDigit", false},
		{'z', ascii.IsDigit, "IsDigit", false},
		{'0', ascii.IsDigit, "IsDigit", true},
		{'9', ascii.IsDigit, "IsDigit", true},
		{'\n', ascii.IsDigit, "IsDigit", false},
		{'\t', ascii.IsDigit, "IsDigit", false},
		{0x1F, ascii.IsDigit, "IsDigit", false},
		{0x7F, ascii.IsDigit, "IsDigit", false},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%s/%c", tc.FuncName, tc.Char), func(t *testing.T) {
			result := tc.Func(tc.Char)
			if result != tc.Result {
				t.Errorf("Expected %t, got %t", tc.Result, result)
			}
		})
	}
}
