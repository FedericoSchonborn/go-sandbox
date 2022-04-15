package ascii

import "math"

// ISO C

func IsAlnum(c byte) bool {
	return IsAlpha(c) || IsDigit(c)
}

func IsAlpha(c byte) bool {
	return IsUpper(c) || IsLower(c)
}

func IsBlank(c byte) bool {
	return c == ' ' || c == '\t'
}

func IsCntrl(c byte) bool {
	return c == 0x7F || (c >= 0x00 && c <= 0x1F)
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func IsGraph(c byte) bool {
	return c != ' ' && IsPrint(c)
}

func IsLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func IsPrint(c byte) bool {
	return c != 0x7F && c >= 0x1F
}

func IsPunct(c byte) bool {
	return !IsAlnum(c) && IsGraph(c)
}

func IsSpace(c byte) bool {
	return IsBlank(c) || c == '\n' || c == '\v' || c == '\f' || c == '\r'
}

func IsUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func IsXdigit(c byte) bool {
	return IsDigit(c) || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
}

func ToLower(c byte) byte {
	if !IsAlpha(c) {
		return c
	}

	return c + 0x20
}

func ToUpper(c byte) byte {
	if !IsAlpha(c) {
		return c
	}

	return c - 0x20
}

// POSIX C

func IsAscii(r rune) bool {
	return r >= 0 && r <= math.MaxUint8
}

func ToAscii(r rune) (byte, bool) {
	if !IsAscii(r) {
		return 0, false
	}

	return byte(r), true
}
